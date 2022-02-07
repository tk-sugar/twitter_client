package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"twitter_client/database"
	"twitter_client/lib"
	"twitter_client/models"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()

	db := database.GormConnect()
	defer db.Close()

	wp_posts := []models.WpPost{}
	status := "publish"
	db.Where("post_title LIKE ? AND post_status = ?", os.Getenv("KEYWORD"), status).Order("post_date desc").Limit(10).Find(&wp_posts)

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(wp_posts))

	text := fmt.Sprintf("%s\n%s%s/\n\n%s",
		wp_posts[num].PostTitle, os.Getenv("URL"), wp_posts[num].PostName, os.Getenv("TAG"))

	api := lib.GetTwitterApi()

	tweet_ret, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet_ret.Text)
}
