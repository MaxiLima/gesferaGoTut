package service

import (
	"fmt"
	"gitlab.grupoesfera.com.ar/gesferaGoTut/CAP-00082-GrupoEsfera-GO/src/domain"
	"time"
)

type TweetManager struct {
	tweetsMap map[string][]*domain.Tweet
	tweets    []*domain.Tweet
}

func NewTweetManager() *TweetManager {
	tm := &TweetManager{tweets: make([]*domain.Tweet, 0), tweetsMap: make(map[string][]*domain.Tweet)}
	return tm
}

func (tm *TweetManager) PublishTweet(tweetToPublish *domain.Tweet) (int, error) {

	now := time.Now()
	tweetToPublish.Date = &now
	tweetToPublish.ID = len(tm.tweets) + 1
	tm.tweets = append(tm.tweets, tweetToPublish)

	tm.tweetsMap[tweetToPublish.User] = tm.GetTweetsByUser(tweetToPublish.User)

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
	return tweetToPublish.ID, fmt.Errorf("nil")
}

func (tm *TweetManager) GetTweets() []*domain.Tweet {
	return tm.tweets
}

func (tm *TweetManager) GetTweetById(id int) *domain.Tweet {

	return tm.tweets[id-1]

}

func (tm *TweetManager) CountTweetsByUser(user string) int {

	var sum int = 0
	for _, valor := range tm.tweets {

		if valor.User == user {
			sum++
		}
	}
	return sum
}

func (tm *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {

	var aux []*domain.Tweet

	for _, valor := range tm.tweets {

		if valor.User == user {
			aux = append(aux, valor)
		}
	}
	return aux
}
