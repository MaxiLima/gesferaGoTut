package domain

import "time"

type Tweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

func NewTweet(user, text string) *Tweet {
	return &Tweet{User: user, Text: text}
}
