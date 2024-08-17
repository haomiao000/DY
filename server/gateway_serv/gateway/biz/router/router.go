package router

import (
	gin "github.com/gin-gonic/gin"
	handler "github.com/haomiao000/DY/server/gateway_serv/gateway/biz/handler"
	middleware "github.com/haomiao000/DY/server/common/middleware"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/assets/public", "./assets/public/")

	apiRouter := r.Group("/douyin")
	apiRouter.Use(middleware.Jaeger())
	// basic apis
		apiRouter.GET("/user/", middleware.VerifyToken(), handler.UserInfo)
		apiRouter.POST("/user/register/", handler.Register)
		apiRouter.POST("/user/login/", handler.Login)
	// apiRouter.Use(middleware.VerifyToken())
	// {
	// 	apiRouter.GET("/feed/", FeedService.Feed)
	// }
	// // apply VerifyToken middleware to /publish routes
	// publishRouter := apiRouter.Group("/publish")
	// publishRouter.Use(middleware.VerifyToken())
	// {
	// 	publishRouter.POST("/action/", PublishService.Publish)
	// 	publishRouter.GET("/list/", PublishService.PublishList)
	// }
	// apply VerifyToken middleware to /favorite routes
	favoriteRouter := apiRouter.Group("/favorite")

	favoriteRouter.Use(middleware.VerifyToken())
	{
		favoriteRouter.GET("/list/", handler.FavoriteList)
		favoriteRouter.POST("/action/", handler.FavoriteAction)
	}
	// apply VerifyToken middleware to /comment routes
	commentRouter := apiRouter.Group("/comment")
	commentRouter.GET("/list/", handler.CommentList)
	commentRouter.Use(middleware.VerifyToken())
	{
		commentRouter.POST("/action/", handler.CommentAction)
	}
	// apply VerifyToken middleware to /relation routes
	relationRouter := apiRouter.Group("/relation")
	relationRouter.Use(middleware.VerifyToken())
	{
		relationRouter.POST("/action/", handler.RelationAction)
		relationRouter.GET("/follow/list/", handler.FollowList)
		relationRouter.GET("/follower/list/", handler.FollowerList)
		relationRouter.GET("/friend/list/", handler.FriendList)
	}
	// apply VerifyToken middleware to /message routes
	// messageRouter := apiRouter.Group("/message")
	// messageRouter.Use(middleware.VerifyToken())
	// {
	// 	messageRouter.GET("/chat/", MessageService.MessageChat)
	// 	messageRouter.POST("/action/", MessageService.MessageAction)
	// }
}
