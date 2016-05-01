package handlers

import (
	"log"
	"net/http"

	"github.com/eric-orenge/tweet_test/model"
	"github.com/eric-orenge/tweet_test/utils"
)

var consumerKey = ""
var consumerSecret = ""
var accessToken = ""
var accessSecret = ""

var t model.TwitterAccess

func init() {
	// accessToken = os.Getenv("TwitterAccessToken")
	// accessSecret = os.Getenv("TwitterAccessSecret")
	// consumerKey = os.Getenv("TwitterConsumerKey")
	// consumerSecret = os.Getenv("TwitterConsumerSecret")

	t.ConsumerKey = consumerKey
	t.ConsumerSecret = consumerSecret

	t.AccessToken = accessToken
	t.AccessSecret = accessSecret
	t.Debug = false
}

func unpackTweet(w http.ResponseWriter, r *http.Request) (string, error) {

	r.ParseForm()
	message := r.FormValue("message")

	// if len(message) > 140 || len(message) == 0 {
	// 	return "", errors.New("Invalid message length")
	// }
	log.Println("message is ", message)
	return message, nil
}
func PostTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	status, er := unpackTweet(w, r)
	if er != nil {
		utils.RespondErr(w, r, "Invalid message length")
		return
	}
	resp, err := t.Tweet(status)
	if err != nil {
		log.Printf("%s", err)

		utils.RespondErr(w, r, "Unable to post tweet")
		return
	}
	utils.RespondSuccess(w, r, resp)
}
