package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func main() {
	loadEnv()

	api := getTwitterApi()

	sendText := "Test"
	tweet_ret, err := api.PostTweet(sendText, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet_ret.Text)
}
