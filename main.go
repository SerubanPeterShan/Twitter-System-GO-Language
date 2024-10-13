package main

import (
	"log"
	"net/http"
	"os"

	"github.com/michimani/gotwi"
)

// main is the entry point of the application. It starts the HTTP server.
func main() {
	log.Println("Starting server...")
	log.Println(os.Getenv("GOTWI_API_KEY"))
	http.HandleFunc("/", UserHandler)

	log.Println("Server started on: http://localhost:8099")
	log.Fatal(http.ListenAndServe(":8099", nil))
}

// UserHandler handles HTTP requests for user-related actions.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		postTweetHandler(w, r)
	case "DELETE":
		deleteTweetHandler(w, r)
	}
}

// getClient initializes and returns a new gotwi.Client using OAuth1 user context.
// gotwi will automatically get the API key and API secret key from the environment variables
// GOTWI_API_KEY and GOTWI_API_KEY_SECRET.
func getClient() (*gotwi.Client, error) {
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	client, err := gotwi.NewClient(&gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           accessToken,
		OAuthTokenSecret:     accessTokenSecret,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
