package service

import (
	"fmt"
	"time"

	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
)

//var tweet *domain.Tweet

var tweetsMap map[string][]*domain.Tweet
var tweets []*domain.Tweet

func Initialize() {
	tweets = make([]*domain.Tweet, 0)
	tweetsMap = make(map[string][]*domain.Tweet)
}

func PublishTweet(tweetToPublish *domain.Tweet) (int, error) {
	now := time.Now()
	tweetToPublish.Date = &now
	tweetToPublish.ID = len(tweets) + 1
	tweets = append(tweets, tweetToPublish)

	tweetsMap[tweetToPublish.User] = GetTweetsByUser(tweetToPublish.User)

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

func GetTweets() []*domain.Tweet {
	return tweets
}

func GetTweetById(id int) *domain.Tweet {

	return tweets[id-1]

}

func CountTweetsByUser(user string) int {

	var sum int = 0
	for _, valor := range tweets {

		if valor.User == user {
			sum++
		}
	}
	return sum
}

func GetTweetsByUser(user string) []*domain.Tweet {

	var aux []*domain.Tweet

	for _, valor := range tweets {

		if valor.User == user {
			aux = append(aux, valor)
		}
	}
	return aux
}
