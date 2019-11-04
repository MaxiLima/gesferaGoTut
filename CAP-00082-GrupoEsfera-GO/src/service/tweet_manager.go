package service

import (
	"fmt"
	"gitlab.grupoesfera.com.ar/gesferaGoTut/CAP-00082-GrupoEsfera-GO/src/domain"
	"time"
)

type TweetManager struct {
	TweetsMap map[string][]domain.Tweet
	Tweets    []domain.Tweet
}

func NewTweetManager() *TweetManager {
	tm := &TweetManager{Tweets: make([]domain.Tweet, 0), TweetsMap: make(map[string][]domain.Tweet)}
	return tm
}

func (tm *TweetManager) PublishTweet(tweetToPublish domain.Tweet) (int, error) {

	now := time.Now()
	tweetToPublish.SetDate(&now)
	tweetToPublish.SetID(len(tm.Tweets) + 1)

	tm.Tweets = append(tm.Tweets, tweetToPublish)
	tm.TweetsMap[tweetToPublish.GetUser()] = tm.GetTweetsByUser(tweetToPublish.GetUser())

	/*if tweet.User == "" {
		return fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return fmt.Errorf("text is required")
	}
	if len(tweet.Text) > 140 {
		return fmt.Errorf("text cannot exced 140 chars")
	}

	return nil*/
	return tweetToPublish.GetID(), fmt.Errorf("nil")
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
