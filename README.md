# Grand Rising API

A simple Go Fiber API that calls the Gemini AI model with a "Grand Rising" greeting template.

## Features

- GET `/grandRising` endpoint that returns AI-generated responses
- Customizable weather parameters via query strings
- Gemini AI integration for generating creative responses

## Setup

1. Clone this repository
2. Copy `.env.example` to `.env` and add your Gemini API key:
   ```
   cp .env.example .env
   ```
3. Edit the `.env` file and replace `your_gemini_api_key_here` with your actual Gemini API key

## Running the API

```bash
go run main.go
```

The API will start on port 3000 by default (configurable in `.env`).

## API Usage

### GET /grandRising

Returns a Gemini-generated response based on a template with your provided weather parameters.

**Query Parameters:**

- `weatherConditions`: Weather conditions (e.g., "sunny and 75°F")
- `precipitationChance`: Chance of precipitation (e.g., "10%")
- `precipitationAmount`: Amount of precipitation (e.g., "0.2 inches")

**Example Request:**

```
GET /grandRising?weatherConditions=cloudy%20and%2068°F&precipitationChance=30%&precipitationAmount=0.5%20inches
```

**Example Response:**

```
Grand rising! 

Today's high will be cloudy and 68°F with a 30% chance of precipitation and 0.5 inches of rain. Perfect weather for ducks and meteorologists with a rain fetish.

I have 01100110 01110101 01101110 01101110 01111001 jokes in my database, but they're all in binary so humans rarely laugh.
