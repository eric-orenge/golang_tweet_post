package model

import (
	"io/ioutil"
	"log"

	"github.com/mrjones/oauth"
)

type Tweet struct {
	Status string
}

//TwitterAccess is a representation of all the details needed by twitter to use the REST API
type TwitterAccess struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
	Debug          bool //to set OAUTH log level, if you want to see the response headers
}

//Tweet tweets a given status, returns the raw twitter api response
func (ta *TwitterAccess) Tweet(status string) (string, error) {
	var endpoint = "https://api.twitter.com/1.1/statuses/update.json"

	consumer := oauth.NewConsumer(ta.ConsumerKey, ta.ConsumerSecret, oauth.ServiceProvider{})
	consumer.Debug(true)
	accessToken := &oauth.AccessToken{Token: ta.AccessToken, Secret: ta.AccessSecret}

	params := map[string]string{
		"status": status,
	}

	response, err1 := consumer.Post(endpoint, params, accessToken)
	if err1 != nil {
		log.Println("LOG_FATAL", err1)
		return "", err1
	}
	defer response.Body.Close()

	_, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println("LOG_FATAL", err2)
		return "", err2
	}
	return "Tweet successfully posted", nil
}
