package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eric-orenge/tweet_test/handlers"
)

func main() {

	// now := time.Now()
	// secs := now.Unix()

	//log file>output all into logs.text
	f, er := os.OpenFile("logs/logs.text", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if er != nil {
		panic(er)
	}
	defer f.Close()

	log.SetOutput(f)

	//handle web requests
	http.HandleFunc("/post", handlers.PostTweet)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Listen And Serve: ", err)
	}

}
