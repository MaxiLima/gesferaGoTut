package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	PrintableTweet() string
	SetDate(i *time.Time)
	SetID(i int)
	GetUser() string
	GetID() int
	GetText() string
}

type TextTweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

func (tweet *TextTweet) SetDate(newDate *time.Time) {
	tweet.Date = newDate
}

func (tweet *TextTweet) SetID(ID int) {
	tweet.ID = ID
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetID() int {
	return tweet.ID
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
	//return "@" + tweet.User + ": " + tweet.Text
}

func NewTweet(user, text string) *TextTweet {
	return &TextTweet{User: user, Text: text}
}
