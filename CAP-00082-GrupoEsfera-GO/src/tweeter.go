package main

import (
	"github.com/abiosoft/ishell"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			var tweet *domain.Tweet

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			tweetText := c.ReadLine()

			c.Print("Write your name: ")

			tweetUser := c.ReadLine()

			tweet = domain.NewTweet(tweetUser, tweetText)

			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.Run()

}
