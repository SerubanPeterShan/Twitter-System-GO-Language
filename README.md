# Twitter Post/Delete API with Go

## Introduction

This project demonstrates how to create a simple API using Go to post and delete tweets on Twitter. It teaches how to interact with the Twitter API using the `gotwi` library and how to handle HTTP requests in a Go web server.

## Setup Instructions

### Step 1: Set Up a Twitter Developer Account

1. Go to the [Twitter Developer](https://developer.twitter.com/en/apps) website.
2. Sign in with your Twitter account.
3. Create a new application to get your API keys and access tokens.

### Step 2: Generate API Keys

1. After creating the application, navigate to the "Keys and tokens" tab.
2. Generate and copy the following keys:
    - API Key
    - API Key Secret
    - Access Token
    - Access Token Secret

### Step 3: Set Up the Environment

1. Create a `.env` file in the root directory of your project and add the following content:

   ```properties
   GOTWI_API_KEY=your_api_key
   GOTWI_API_KEY_SECRET=your_api_key_secret
   ACCESS_TOKEN=your_access_token
   ACCESS_TOKEN_SECRET=your_access_token_secret
   ```

2. Ensure that `.env` is listed in your `.gitignore` to keep your keys secure.

### Step 4: Run the Program

#### Using Docker Compose

1. Build and run the Docker container using Docker Compose:

   ```sh
   docker-compose up --build
   ```

2. The server will start on `http://localhost:8099`.

#### Without Docker Compose

##### On Linux

1. Export the environment variables:

   ```sh
   export GOTWI_API_KEY=your_api_key
   export GOTWI_API_KEY_SECRET=your_api_key_secret
   export ACCESS_TOKEN=your_access_token
   export ACCESS_TOKEN_SECRET=your_access_token_secret
   ```

2. Run the Go application:

   ```sh
   go run .
   ```

##### On Windows

1. Create a `setenv.bat` file with the following content:

   ```bat
   @echo off
   set GOTWI_API_KEY=your_api_key
   set GOTWI_API_KEY_SECRET=your_api_key_secret
   set ACCESS_TOKEN=your_access_token
   set ACCESS_TOKEN_SECRET=your_access_token_secret
   ```

2. Run the `setenv.bat` file to set the environment variables:

   ```bat
   setenv.bat
   ```

3. Run the Go application:

   ```bat
   go run .
   ```

## Program Details

### Posting a New Tweet

The `postTweetHandler` function handles POST requests to post a new tweet. It decodes the request payload, creates a new tweet using the `gotwi` client, and returns the tweet ID.

Example API request:

```sh
curl -X POST http://localhost:8099/ -d '{"tweet": "Hello, Twitter!"}' -H "Content-Type: application/json"
```

Example response:

```json
{
  "tweet_id": "1234567890123456789"
}
```

### Deleting an Existing Tweet

The `deleteTweetHandler` function handles DELETE requests to delete an existing tweet. It decodes the request payload, deletes the tweet using the `gotwi` client, and returns a success message.

Example API request:

```sh
curl -X DELETE http://localhost:8099/ -d '{"tweet_id": "1234567890123456789"}' -H "Content-Type: application/json"
```

Example response:

```json
{
  "status": "Tweet deleted successfully",
  "deleted": true
}
```

## Error Handling

The program handles API errors and invalid inputs by returning appropriate HTTP status codes and error messages. For example:

- If the request payload is invalid, it returns a `400 Bad Request` status.
- If there is an error with the Twitter API, it returns a `500 Internal Server Error` status with the error message.
