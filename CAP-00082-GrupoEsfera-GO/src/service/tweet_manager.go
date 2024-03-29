package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
)

type TweetWriter interface {
	Write(t domain.Tweet)
}

func NewMemoryTweetWriter() TweetWriter {
	return &MemoryTweetWriter{}
}

type MemoryTweetWriter struct {
	tweet domain.Tweet
}

type FileTweetWriter struct {
	file *os.File
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0666)
	ftw := new(FileTweetWriter)
	ftw.file = file
	return ftw
}

func (Ftw *FileTweetWriter) Write(t domain.Tweet) {
	go func() {
		Ftw.file.WriteString(t.PrintableTweet() + "\n")
		Ftw.file.WriteString("\n")
	}()
}

func (Mtw *MemoryTweetWriter) Write(t domain.Tweet) {
	Mtw.tweet = t
}

func (Mtw *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return Mtw.tweet
}

type TweetManager struct {
	TweetsMap map[string][]domain.Tweet
	Tweets    []domain.Tweet
	Twriter   TweetWriter
}

func NewTweetManager(tw TweetWriter) *TweetManager {
	tm := &TweetManager{Tweets: make([]domain.Tweet, 0), TweetsMap: make(map[string][]domain.Tweet), Twriter: tw}
	return tm
}

func (tm *TweetManager) PublishTweet(tweetToPublish domain.Tweet) (int, error) {

	now := time.Now()
	tweetToPublish.SetDate(&now)
	tweetToPublish.SetID(len(tm.Tweets) + 1)

	tm.Tweets = append(tm.Tweets, tweetToPublish)
	tm.TweetsMap[tweetToPublish.GetUser()] = tm.GetTweetsByUser(tweetToPublish.GetUser())

	if tweetToPublish.GetUser() == "" {
		return -1, fmt.Errorf("user is required")
	}
	if tweetToPublish.GetText() == "" {
		return -1, fmt.Errorf("text is required")
	}
	if len(tweetToPublish.GetText()) > 140 {
		return -1, fmt.Errorf("text cannot exced 140 chars")
	}

	return tweetToPublish.GetID(), nil
}

func (tm *TweetManager) GetTweets() []domain.Tweet {
	return tm.Tweets
}

func (tm *TweetManager) GetLastTweet() domain.Tweet {
	return tm.Tweets[len(tm.Tweets)-1]
}

func (tm *TweetManager) GetTweetById(id int) domain.Tweet {

	return tm.Tweets[id-1]

}

func (tm *TweetManager) CountTweetsByUser(user string) int {

	var sum int = 0
	for _, valor := range tm.Tweets {

		if valor.GetUser() == user {
			sum++
		}
	}
	return sum
}

func (tm *TweetManager) GetTweetsByUser(user string) []domain.Tweet {

	var aux []domain.Tweet

	for _, valor := range tm.Tweets {

		if valor.GetUser() == user {
			aux = append(aux, valor)
		}
	}
	return aux
}

func (tm *TweetManager) SearchTweetsContaining(query string, tweets chan domain.Tweet) {
	go func() {
		for _, valor := range tm.Tweets {

			if strings.Contains(valor.GetText(), query) {
				tweets <- valor
			}
		}
		close(tweets)
	}()
}
