package main

import (
	"encoding/json"
	"net/http"

	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

type TweetRequest struct {
	Tweet string `json:"tweet"`
}

// postTweetHandler handles HTTP requests to post a tweet.
// It decodes the request payload, retrieves a Twitter client,
// posts the tweet, and returns the tweet ID in the response.
func postTweetHandler(w http.ResponseWriter, r *http.Request) {
	var tweetRequest TweetRequest
	if err := json.NewDecoder(r.Body).Decode(&tweetRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	client, err := getClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tweetID, err := postTweet(client, tweetRequest.Tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"tweet_id": tweetID})
}

// postTweet posts a tweet using the provided Twitter client and tweet text.
// It returns the ID of the posted tweet or an error if the operation fails.
func postTweet(client *gotwi.Client, tweet string) (string, error) {
	input := &types.CreateInput{
		Text: &tweet,
	}
	tweetData, err := managetweet.Create(context.Background(), client, input)
	if err != nil {
		return "", err
	}
	return *tweetData.Data.ID, nil
}
