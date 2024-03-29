package main

import (
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/rest"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/service"
)

func main() {

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	gs := rest.NewGinServer(tweetManager)

	gs.StartServer()

	/*shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	var ftw service.TweetWriter
	ftw = service.NewFileTweetWriter() // Mock implementation
	tm := service.NewTweetManager(ftw)

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			var textTweet *domain.TextTweet

			defer c.ShowPrompt(true)

			c.Print("Write your textTweet: ")

			tweetText := c.ReadLine()

			c.Print("Write your name: ")

			tweetUser := c.ReadLine()

			textTweet = domain.NewTextTweet(tweetUser, tweetText)

			tm.PublishTweet(textTweet)

			c.Print("TextTweet sent\n")

			tw := service.NewFileTweetWriter()

			tw.Write(textTweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			lastTextTweet := tm.GetLastTweet()

			c.Println(lastTextTweet)

			return
		},
	})

	shell.Run()
	*/
}
