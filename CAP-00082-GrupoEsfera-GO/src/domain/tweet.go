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

type ImageTweet struct {
	TextTweet
	Url string
}

type QuoteTweet struct {
	TextTweet
	citado Tweet
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
}

func (tweet *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", tweet.User, tweet.Text, tweet.Url)
}

func (tweet *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"%s\"", tweet.User, tweet.Text, tweet.citado.PrintableTweet())
}

func NewTextTweet(user, text string) *TextTweet {
	return &TextTweet{User: user, Text: text}
}

func NewImageTweet(user, text, image string) *ImageTweet {
	return &ImageTweet{TextTweet: TextTweet{User: user, Text: text}, Url: image}
}

func NewQuoteTweet(user string, text string, citado Tweet) *QuoteTweet {
	return &QuoteTweet{TextTweet: TextTweet{User: user, Text: text}, citado: citado}
}
