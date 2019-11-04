package domain_test

import (
	"fmt"
	"gitlab.grupoesfera.com.ar/gesferaGoTut/CAP-00082-GrupoEsfera-GO/src/domain"
	"testing"
)

func TestCanGetAPrintableTweet(t *testing.T) {

	// Initialization
	textTweet := domain.NewTextTweet("grupoesfera", "This is my textTweet")

	// Operation
	text := textTweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my textTweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	var textTweet fmt.Stringer
	textTweet = domain.NewTextTweet("grupoesfera", "This is my textTweet")

	// Operation
	text := textTweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my textTweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", *quotedTweet)
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}
