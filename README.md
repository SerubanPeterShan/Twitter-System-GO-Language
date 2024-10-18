# Twitter Post/Delete API with Go

## Introduction

This project demonstrates how to create a simple API using Go to post and delete tweets on Twitter. It covers how to interact with the Twitter API using the `gotwi` library and handle HTTP requests in a Go web server.

## Setup Instructions

### Step 1: Set Up a Twitter Developer Account

1. Visit the [Twitter Developer](https://developer.twitter.com/en/apps) website.
2. Sign in with your Twitter account.
3. Create a new application to obtain your API keys and access tokens.

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

### Install `gotwi`

To interact with the Twitter API, you need to install the `gotwi` library. Run the following command to add it to your project:

```sh
go get github.com/michimani/gotwi
```

### Step 5: Run the Program

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

3. Install the required libraries:

   ```bat
   go mod downlaod
   ```

4. Run the Go application:

   ```bat
   go run .
   ```

## Program Details

### Posting a New Tweet -POST tweet

The `postTweetHandler` function handles POST requests to post a new tweet. It decodes the request payload, creates a new tweet using the `gotwi` client, and returns the tweet ID.

#### Function Overview -POST tweet

- **Function Name**: `postTweetHandler`
- **HTTP Method**: POST
- **Endpoint**: `/`
- **Request Payload**: JSON object containing the tweet text.
- **Response**: JSON object containing the tweet ID.

#### Example API Request -POST tweet

To post a new tweet, use the following `curl` command:

```sh
curl -X POST http://localhost:8099/ -d '{"tweet": "Hello, Twitter!"}' -H "Content-Type: application/json"
```

#### Example API Response -POST tweet

If the tweet is successfully posted, the response will be:

```json
{
   "tweet_id": "1234567890123456789"
}
```

#### Detailed Steps -POST tweet

1. **Decode Request Payload**: The function decodes the incoming JSON payload to extract the tweet text.
2. **Create Tweet**: It uses the `gotwi` client to create a new tweet with the extracted text.
3. **Return Response**: It returns a JSON response containing the ID of the newly created tweet.

#### Error Handling -POST tweet

The `postTweetHandler` function includes error handling to manage various scenarios:

- **Invalid Payload**: If the request payload is invalid or missing the tweet text, it returns a `400 Bad Request` status.
- **Twitter API Errors**: If there is an error while interacting with the Twitter API, it returns a `500 Internal Server Error` status along with the error message.

### Posting a Tweet-POST tweet

The `postTweet` function is responsible for posting a tweet using the provided Twitter client and tweet text. It returns the ID of the posted tweet or an error if the operation fails.

#### Function Overview-POST tweet

- **Function Name**: `postTweet`
- **Parameters**:
  - `client`: The Twitter client used to interact with the Twitter API.
  - `tweet`: The text of the tweet to be posted.
- **Returns**: The ID of the posted tweet or an error.

#### Detailed Steps-POST tweet

1. **Create Input**: The function creates an input object containing the tweet text.
2. **Post Tweet**: It calls the `Create` method from the `gotwi` library to post the tweet.
3. **Return Tweet ID**: If successful, it returns the ID of the posted tweet; otherwise, it returns an error.

#### Error Handling-POST tweet

The `postTweet` function includes error handling to manage various scenarios:

- **Twitter API Errors**: If there is an error while interacting with the Twitter API, it returns the error.

By following these steps, you can successfully post new tweets using the provided Go application.

### Deleting an Existing Tweet

The `deleteTweetHandler` function handles DELETE requests to delete an existing tweet. It decodes the request payload, deletes the tweet using the `gotwi` client, and returns a success message.

#### Function Overview -DELETE tweet

- **Function Name**: `deleteTweetHandler`
- **HTTP Method**: DELETE
- **Endpoint**: `/`
- **Request Payload**: JSON object containing the tweet ID.
- **Response**: JSON object indicating the success status and deletion confirmation.

#### Example API Request -DELETE tweet

To delete an existing tweet, use the following `curl` command:

```sh
curl -X DELETE http://localhost:8099/ -d '{"tweet_id": "1234567890123456789"}' -H "Content-Type: application/json"
```

#### Example API Response -DELETE tweet

If the tweet is successfully deleted, the response will be:

```json
{
   "status": "Tweet deleted successfully",
   "deleted": true
}
```

### Detailed Steps -DELETE tweet

1. **Decode Request Payload**: The function decodes the incoming JSON payload to extract the tweet ID.
2. **Delete Tweet**: It uses the `gotwi` client to delete the tweet with the extracted ID.
3. **Return Response**: It returns a JSON response indicating the success status and deletion confirmation.

### Error Handling -DELETE tweet

The `deleteTweetHandler` function includes error handling to manage various scenarios:

- **Invalid Payload**: If the request payload is invalid or missing the tweet ID, it returns a `400 Bad Request` status.
- **Twitter API Errors**: If there is an error while interacting with the Twitter API, it returns a `500 Internal Server Error` status along with the error message.

### Deleting a Tweet

The `deleteTweet` function is responsible for deleting a tweet using the provided Twitter client and tweet ID. It returns a boolean indicating whether the tweet was successfully deleted or an error if the operation fails.

#### Function Overview-DELETE tweet

- **Function Name**: `deleteTweet`
- **Parameters**:
  - `client`: The Twitter client used to interact with the Twitter API.
  - `tweetID`: The ID of the tweet to be deleted.
- **Returns**: A boolean indicating whether the tweet was successfully deleted or an error.

#### Detailed Steps-DELETE tweet

1. **Create Input**: The function creates an input object containing the tweet ID.
2. **Delete Tweet**: It calls the `Delete` method from the `gotwi` library to delete the tweet.
3. **Return Deletion Status**: If successful, it returns `true`; otherwise, it returns an error.

#### Error Handling-Delete tweet

The `deleteTweet` function includes error handling to manage various scenarios:

- **Twitter API Errors**: If there is an error while interacting with the Twitter API, it returns the error.

By following these steps, you can successfully delete existing tweets using the provided Go application.

## Error Handling

The program handles API errors and invalid inputs by returning appropriate HTTP status codes and error messages. For example:

- **400 Bad Request**: Returned if the request payload is invalid.
- **500 Internal Server Error**: Returned if there is an error with the Twitter API, along with the error message.

By following these guidelines, you can effectively interact with the Twitter API to post and delete tweets using the provided Go application.

