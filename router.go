package main

import (
	"github.com/gin-gonic/gin"
	"main/pkg/controller"
	"main/middleware"
)

func initRouter(r *gin.Engine) {
    // public directory is used to serve static resources
    r.Static("/static", "./public")

    apiRouter := r.Group("/douyin")
    
    // basic apis
    apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.VerifyToken(), controller.UserInfo)
    apiRouter.POST("/user/register/", controller.Register)
    apiRouter.POST("/user/login/", controller.Login)
	
    // apply VerifyToken middleware to /publish routes
    publishRouter := apiRouter.Group("/publish")
    publishRouter.Use(middleware.VerifyToken())
	{
		publishRouter.POST("/action/", controller.Publish)
		publishRouter.GET("/list/", controller.PublishList)
	}
    // apply VerifyToken middleware to /favorite routes
    favoriteRouter := apiRouter.Group("/favorite")
    favoriteRouter.Use(middleware.VerifyToken())
	{
		favoriteRouter.POST("/action/", controller.FavoriteAction)
		favoriteRouter.GET("/list/", controller.FavoriteList)
	}
    // apply VerifyToken middleware to /comment routes
    commentRouter := apiRouter.Group("/comment")
    commentRouter.Use(middleware.VerifyToken())
	{
		commentRouter.POST("/action/", controller.CommentAction)
		commentRouter.GET("/list/", controller.CommentList)
	}
    // apply VerifyToken middleware to /relation routes
    relationRouter := apiRouter.Group("/relation")
    relationRouter.Use(middleware.VerifyToken())
	{
		relationRouter.POST("/action/", controller.RelationAction)
		relationRouter.GET("/follow/list/", controller.FollowList)
		relationRouter.GET("/follower/list/", controller.FollowerList)
		relationRouter.GET("/friend/list/", controller.FriendList)
	}
    // apply VerifyToken middleware to /message routes
    messageRouter := apiRouter.Group("/message")
    messageRouter.Use(middleware.VerifyToken())
	{
		messageRouter.GET("/chat/", controller.MessageChat)
		messageRouter.POST("/action/", controller.MessageAction)
	}
}