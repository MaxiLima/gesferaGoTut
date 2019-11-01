package service_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/service"
)

/*
func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}

	//assert.Equal(t, publishedTweet.User)
}
func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}

	//Assert
	assert.EqualError(t, err, "user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	var text string
	user := "Maxo"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}

	//Assert
	assert.EqualError(t, err, "text is required")
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var text string
	user := "Maxo"
	text = "asdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkd"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text cannot exced 140 chars" {
		t.Error("Expected error is text is required")
	}

	//Assert
	assert.EqualError(t, err, "text cannot exced 140 chars")
}
*/
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	service.Initialize()
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data

	tweet = domain.NewTweet("Maxi", "tweet1")
	secondTweet = domain.NewTweet("Maxi", "tweet2")

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	validTweet(t, firstPublishedTweet, "Maxi", "tweet1")
	validTweet(t, secondPublishedTweet, "Maxi", "tweet2")

	// Same for secondPublishedTweet
}

func validTweet(t *testing.T, tweet *domain.Tweet, user, text string) {

	assert.Equal(t, user, tweet.User)
	assert.Equal(t, text, tweet.Text)

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.Initialize()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	validTweetId(t, publishedTweet, id, user, text)
}

func validTweetId(t *testing.T, tweet *domain.Tweet, id int, user, text string) {

	assert.Equal(t, id, tweet.ID)
	assert.Equal(t, user, tweet.User)
	assert.Equal(t, text, tweet.Text)

}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.Initialize()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.Initialize()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	// publish the 3 tweets

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	tweets := service.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */
		fmt.Errorf("error")

	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet

	validTweet(t, firstPublishedTweet, user, text)
	validTweet(t, secondPublishedTweet, user, secondText)

}