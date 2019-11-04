package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

func (tweet TextTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
	//return "@" + tweet.User + ": " + tweet.Text
}

func NewTweet(user, text string) *TextTweet {
	return &TextTweet{User: user, Text: text}
}
