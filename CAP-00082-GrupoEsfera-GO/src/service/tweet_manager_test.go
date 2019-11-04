package service_test

import (
	"fmt"
	"gitlab.grupoesfera.com.ar/gesferaGoTut/CAP-00082-GrupoEsfera-GO/src/domain"
	"gitlab.grupoesfera.com.ar/gesferaGoTut/CAP-00082-GrupoEsfera-GO/src/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.TextTweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

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
	var tweet *domain.TextTweet

	var user string
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

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
	var tweet *domain.TextTweet

	var text string
	user := "Maxo"
	tweet = domain.NewTextTweet(user, text)

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
	var tweet *domain.TextTweet

	var text string
	user := "Maxo"
	text = "asdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkdasdasljknsdkfjnasdjkfnsdkjfnsdkjnsdkjfnsdkjfnsjkd"
	tweet = domain.NewTextTweet(user, text)

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
}*/

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization

	tm := service.NewTweetManager()

	var tweet, secondTweet *domain.TextTweet // Fill the Tweets with data

	tweet = domain.NewTextTweet("Maxi", "tweet1")
	secondTweet = domain.NewTextTweet("Maxi", "tweet2")

	// Operation
	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tm.GetTweets()
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

func validTweet(t *testing.T, tweet domain.Tweet, user, text string) {

	assert.Equal(t, user, tweet.GetUser())
	assert.Equal(t, text, tweet.GetText())

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization

	tm := service.NewTweetManager()

	var tweet *domain.TextTweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tm.PublishTweet(tweet)

	// Validation
	publishedTweet := tm.GetTweetById(id)

	validTweetId(t, publishedTweet, id, user, text)
}

func validTweetId(t *testing.T, tweet domain.Tweet, id int, user, text string) {

	assert.Equal(t, id, tweet.GetID())
	assert.Equal(t, user, tweet.GetUser())
	assert.Equal(t, text, tweet.GetText())

}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization

	tm := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)
	tm.PublishTweet(thirdTweet)
	// Operation
	count := tm.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization

	tm := service.NewTweetManager()

	var textTweet, secondTextTweet, thirdTextTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first textTweet"
	secondText := "This is my second textTweet"
	textTweet = domain.NewTextTweet(user, text)
	secondTextTweet = domain.NewTextTweet(user, secondText)
	thirdTextTweet = domain.NewTextTweet(anotherUser, text)
	// publish the 3 Tweets

	tm.PublishTweet(textTweet)
	tm.PublishTweet(secondTextTweet)
	tm.PublishTweet(thirdTextTweet)

	// Operation
	tweets := tm.GetTweetsByUser(user)

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
