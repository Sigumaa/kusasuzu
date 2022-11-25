package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

var metaBytes []byte

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("API_KEY")
	API_KEY_SECRET := os.Getenv("API_KEY_SECRET")
	ACCESS_TOKEN := os.Getenv("ACCESS_TOKEN")
	ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")
	BEARER_TOKEN := os.Getenv("BEARER_TOKEN")
	DISCORD_TOKEN := os.Getenv("DISCORD_TOKEN")
	DISCORD_WEBHOOK := os.Getenv("DISCORD_WEBHOOK")
	CHANNEL_ID := os.Getenv("CHANNEL_ID")
	hoge := fmt.Sprintf("%s\n %s\n %s\n %s\n %s\n %s\n %s\n %s\n", API_KEY, API_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET, BEARER_TOKEN, DISCORD_TOKEN, DISCORD_WEBHOOK, CHANNEL_ID)
	fmt.Println(hoge)

	client := &twitter.Client{
		Authorizer: authorize{
			Token: BEARER_TOKEN,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	opts := twitter.TweetRecentSearchOpts{
		Expansions:  []twitter.Expansion{twitter.ExpansionEntitiesMentionsUserName, twitter.ExpansionAuthorID},
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldConversationID, twitter.TweetFieldAttachments},
	}

	fmt.Println("Callout to tweet recent search callout")

	query := "草鈴 has:links -is:retweet"
	tweetResponse, err := client.TweetRecentSearch(context.Background(), query, opts)
	if err != nil {
		log.Panicf("tweet lookup error: %v", err)
	}

	dictionaries := tweetResponse.Raw.TweetDictionaries()

	enc, err := json.MarshalIndent(dictionaries, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))

	metaBytes, err = json.MarshalIndent(tweetResponse.Meta, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(metaBytes))
}