package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/domain"
	"gitlab.grupoesfera.com.ar/CAP-00082-GrupoEsfera-GO/src/service"
)

type GinTextTweet struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type GinImageTweet struct {
	GinTextTweet
	Url string `json:"url"`
}

type GinQuoteTweet struct {
	GinTextTweet
	QuotedTweetId int `json:"quotedtweet"`
}

type GinServer struct {
	tweetManager *service.TweetManager
}

func NewGinServer(tm *service.TweetManager) *GinServer {
	return &GinServer{tweetManager: tm}
}

func (gs *GinServer) StartServer() {
	router := gin.Default()

	router.GET("/tweets", gs.getTweets)
	router.GET("/tweets/:idTweet", gs.getTweetsById)
	router.POST("/publishTextTweet", gs.publishTextTweet)
	router.POST("/publishImageTweet", gs.publishTextTweet)
	router.POST("/publishQuotedTweet", gs.publishQuotedTweet)
	router.Run(":8081")
}

func (gs *GinServer) getTweets(c *gin.Context) {

	c.JSON(http.StatusOK, gs.tweetManager.GetTweets())
}

func (gs *GinServer) getTweetsById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Params.ByName("idTweet"))

	c.JSON(http.StatusOK, gs.tweetManager.GetTweetById(id))
}

func (gs *GinServer) publishTextTweet(c *gin.Context) {

	var tweetdata GinTextTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewTextTweet(tweetdata.User, tweetdata.Text)

	id, err := gs.tweetManager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet : "+err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func (gs *GinServer) publishImageTweet(c *gin.Context) {

	var tweetdata GinImageTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewImageTweet(tweetdata.User, tweetdata.Text, tweetdata.Url)

	id, err := gs.tweetManager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet : "+err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func (gs *GinServer) publishQuotedTweet(c *gin.Context) {

	var tweetdata GinQuoteTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewQuoteTweet(tweetdata.User, tweetdata.Text, gs.tweetManager.GetTweetById(tweetdata.QuotedTweetId))

	id, err := gs.tweetManager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet : "+err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}
