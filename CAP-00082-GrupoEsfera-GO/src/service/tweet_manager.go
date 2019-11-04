package service

import (
	"fmt"
	"gitlab.grupoesfera.com.ar/gesferaGoTut/CAP-00082-GrupoEsfera-GO/src/domain"
	"time"
)

type TweetManager struct {
	textTweetsMap map[string][]*domain.TextTweet
	textTweets    []*domain.TextTweet
}

func NewTweetManager() *TweetManager {
	tm := &TweetManager{textTweets: make([]*domain.TextTweet, 0), textTweetsMap: make(map[string][]*domain.TextTweet)}
	return tm
}

func (tm *TweetManager) PublishTweet(tweetToPublish *domain.TextTweet) (int, error) {

	now := time.Now()
	tweetToPublish.Date = &now
	tweetToPublish.ID = len(tm.textTweets) + 1
	tm.textTweets = append(tm.textTweets, tweetToPublish)

	tm.textTweetsMap[tweetToPublish.User] = tm.GetTweetsByUser(tweetToPublish.User)

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

func (tm *TweetManager) GetTweets() []*domain.TextTweet {
	return tm.textTweets
}

func (tm *TweetManager) GetTweetById(id int) *domain.TextTweet {

	return tm.textTweets[id-1]

}

func (tm *TweetManager) CountTweetsByUser(user string) int {

	var sum int = 0
	for _, valor := range tm.textTweets {

		if valor.User == user {
			sum++
		}
	}
	return sum
}

func (tm *TweetManager) GetTweetsByUser(user string) []*domain.TextTweet {

	var aux []*domain.TextTweet

	for _, valor := range tm.textTweets {

		if valor.User == user {
			aux = append(aux, valor)
		}
	}
	return aux
}
