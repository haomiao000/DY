package router

import (
	"main/middleware"
	CommentService "main/server/service/comment/pkg"
	FavoriteService "main/server/service/favorite/pkg"
	FeedService "main/server/service/feed/pkg"
	MessageService "main/server/service/message/pkg"
	PublishService "main/server/service/publish/pkg"
	RelationService "main/server/service/relation/pkg"
	UserService "main/server/service/user/pkg"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/assets/public", "./assets/public/")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", FeedService.Feed)
	apiRouter.GET("/user/", middleware.VerifyToken(), UserService.UserInfo)
	apiRouter.POST("/user/register/", UserService.Register)
	apiRouter.POST("/user/login/", UserService.Login)

	// apply VerifyToken middleware to /publish routes
	publishRouter := apiRouter.Group("/publish")
	publishRouter.Use(middleware.VerifyToken())
	{
		publishRouter.POST("/action/", PublishService.Publish)
		publishRouter.GET("/list/", PublishService.PublishList)
	}
	// apply VerifyToken middleware to /favorite routes
	favoriteRouter := apiRouter.Group("/favorite")

	favoriteRouter.Use(middleware.VerifyToken())
	{
		favoriteRouter.GET("/list/", FavoriteService.FavoriteList)
		favoriteRouter.POST("/action/", FavoriteService.FavoriteAction)
	}
	// apply VerifyToken middleware to /comment routes
	commentRouter := apiRouter.Group("/comment")
	commentRouter.GET("/list/", CommentService.CommentList)
	commentRouter.Use(middleware.VerifyToken())
	{
		commentRouter.POST("/action/", CommentService.CommentAction)
	}
	// apply VerifyToken middleware to /relation routes
	relationRouter := apiRouter.Group("/relation")
	relationRouter.Use(middleware.VerifyToken())
	{
		relationRouter.POST("/action/", RelationService.RelationAction)
		relationRouter.GET("/follow/list/", RelationService.FollowList)
		relationRouter.GET("/follower/list/", RelationService.FollowerList)
		relationRouter.GET("/friend/list/", RelationService.FriendList)
	}
	// apply VerifyToken middleware to /message routes
	messageRouter := apiRouter.Group("/message")
	messageRouter.Use(middleware.VerifyToken())
	{
		messageRouter.GET("/chat/", MessageService.MessageChat)
		messageRouter.POST("/action/", MessageService.MessageAction)
	}
}
