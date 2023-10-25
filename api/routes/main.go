package routes

import (
	controllers "api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()

	v1 := router.Group("/v1") 
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweet", tweetController.CreateTweet)
		v1.DELETE("/tweet/:id", tweetController.Delete)
	}

	return v1
}