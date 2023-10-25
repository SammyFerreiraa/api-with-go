package controllers

import (
	entities "api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweet []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweet)
}

func (t *tweetController) CreateTweet(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	t.tweet = append(t.tweet, *tweet)

	ctx.JSON(http.StatusOK, tweet)
}

func (t *tweetController) Delete(ctx *gin.Context) {
	id:= ctx.Param("id")

	for idx, tweet := range t.tweet {
		if id == tweet.ID {
			t.tweet = append(t.tweet[0:idx], t.tweet[idx+1:]...)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})
}