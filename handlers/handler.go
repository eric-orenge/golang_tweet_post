package handlers

import (
	"net/http"

	"github.com/eric-orenge/tweet_test/utils"
)

var consumerKey string
var consumerSecret string
var accessSecret string
var accessToken string
var twitterURL string

func init() {

	accessToken = "312661274-2UIFgzw89B9r3B91nTu4cM52hNqQlm8U4f3iniPn"
	accessSecret = "yotx0h9yWAwjqMOaOu6xvDXELgQdfl11Ump4ICZlq6mjT"

	consumerKey = "1n5knNIgqWdAYNRt88jLmf08f"
	consumerSecret = "5uTZB7yThyDE5GR4rHzn6M3re3xDLrC6S9JbX1IR7Xg5WpDUnR"
}

func unpackTweet(w http.ResponseWriter, r *http.Request) {
	tweet := r.FormValue("message")
	if len(tweet) > 140 {
		utils.RespondErr(w, r, "Invalid message length")
	}
}
func PostTweet(w http.ResponseWriter, r *http.Request) {

}
