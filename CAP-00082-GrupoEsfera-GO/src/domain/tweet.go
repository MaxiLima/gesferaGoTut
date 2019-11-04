package domain

import (
	"fmt"
	"time"
)

type Tweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

func (tweet Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
	//return "@" + tweet.User + ": " + tweet.Text
}

func NewTweet(user, text string) *Tweet {
	return &Tweet{User: user, Text: text}
}
