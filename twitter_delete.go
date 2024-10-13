package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

// DeleteTweetRequest represents the structure of the request payload for deleting a tweet.
type DeleteTweetRequest struct {
	TweetID string `json:"tweet_id"`
}

// deleteTweetHandler handles HTTP requests for deleting a tweet.
// It decodes the request payload, initializes the Twitter client, and calls the deleteTweet function.
// If successful, it responds with a success message; otherwise, it returns an error.
func deleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	var deleteTweetRequest DeleteTweetRequest
	if err := json.NewDecoder(r.Body).Decode(&deleteTweetRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	client, err := getClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := deleteTweet(client, deleteTweetRequest.TweetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "Tweet deleted successfully", "deleted": b})
}

// deleteTweet deletes a tweet with the given tweetID using the provided Twitter client.
// It returns true if the tweet was successfully deleted, otherwise it returns an error.
func deleteTweet(client *gotwi.Client, tweetID string) (bool, error) {
	input := &types.DeleteInput{
		ID: tweetID,
	}
	res, err := managetweet.Delete(context.Background(), client, input)
	if err != nil {
		return false, err
	}
	return gotwi.BoolValue(res.Data.Deleted), nil
}
