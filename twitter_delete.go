package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

type DeleteTweetRequest struct {
	TweetID string `json:"tweet_id"`
}

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

func deleteTweet(client *gotwi.Client, tweetID string) (bool, error) {
	input := &types.DeleteInput{
		ID: tweetID,
	}
	res, err := managetweet.Delete(context.Background(), client, input)
	log.Println(managetweet.Delete(context.Background(), client, input))
	if err != nil {
		return false, err
	}
	return gotwi.BoolValue(res.Data.Deleted), nil
}
