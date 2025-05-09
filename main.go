package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// Get API key from environment variables
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is required")
	}

	// Create Gemini client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	// Get the Gemini Pro model
	model := client.GenerativeModel("gemini-2.0-flash")

	// Create a new Fiber app
	app := fiber.New()

	// Use middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Grand Rising API is running! Try the /grandRising endpoint with weather parameters.")
	})

	app.Get("/grandRising", func(c *fiber.Ctx) error {
		// Get query parameters
		weatherConditions := c.Query("weatherConditions", "sunny and 75°F")
		precipitationChance := c.Query("precipitationChance", "0%")
		precipitationAmount := c.Query("precipitationAmount", "0 inches")

		// Create the prompt
		prompt := fmt.Sprintf(
			"Your friend just said \"Grand Rising\" which is another way to say \"Good morning\" write a reply that says \"Grand rising\". Then give a joke deadpan response about today's weather but do not reference any other days weather. Today the high will be %s and the precipitation chance is %s and the precipitation amount is %s. Then, after 1 line break make a random joke a robot would make that doesn't end in a question.",
			weatherConditions, precipitationChance, precipitationAmount,
		)

		// Generate response from Gemini
		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			log.Printf("Error generating content: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error generating content")
		}

		// Extract the response text
		var responseText string
		if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
			responseText = fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
		} else {
			responseText = "No response generated"
		}

		// Set content type to plain text
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return c.SendString(responseText)
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
