package domain_test

import (
	"fmt"
	"gitlab.grupoesfera.com.ar/gesferaGoTut/CAP-00082-GrupoEsfera-GO/src/domain"
	"testing"
)

func TestCanGetAPrintableTweet(t *testing.T) {

	// Initialization
	textTweet := domain.NewTweet("grupoesfera", "This is my textTweet")

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
	textTweet = domain.NewTweet("grupoesfera", "This is my textTweet")

	// Operation
	text := textTweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my textTweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
