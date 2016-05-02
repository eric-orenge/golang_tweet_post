package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/eric-orenge/tweet_test/twitter"
	"github.com/eric-orenge/tweet_test/utils"
)

var (
	t                                                      twitter.TwitterAccess
	accessToken, accessSecret, consumerKey, consumerSecret string
)

func init() {
	accessToken = os.Getenv("TwitterAccessToken")
	accessSecret = os.Getenv("TwitterAccessSecret")
	consumerKey = os.Getenv("TwitterConsumerKey")
	consumerSecret = os.Getenv("TwitterConsumerSecret")

	t.ConsumerKey = consumerKey
	t.ConsumerSecret = consumerSecret

	t.AccessToken = accessToken
	t.AccessSecret = accessSecret
	t.Debug = false
}

func unpackTweet(w http.ResponseWriter, r *http.Request) (string, error) {

	r.ParseForm()
	message := r.FormValue("message")

	if len(message) > 140 || len(message) == 0 {
		return "", errors.New("Invalid message length")
	}
	return message, nil
}
func PostTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	status, er := unpackTweet(w, r)
	if er != nil {
		utils.RespondErr(w, r, "Invalid message length")
		return
	}
	_, err := t.Tweet(status) //resp for the whole t
	if err != nil {
		log.Printf("%s", err)

		utils.RespondErr(w, r, "Unable to post tweet")
		return
	}
	utils.RespondSuccess(w, r, "Tweet successfully posted")
}
