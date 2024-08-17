package configs

import (
	rpc_interact "github.com/haomiao000/DY/internal/grpc_gen/rpc_interact"
	rpc_relation "github.com/haomiao000/DY/internal/grpc_gen/rpc_relation"
	rpc_user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
)

var (
	GlobalUserClient     rpc_user.UserServiceImplClient
	GlobalInteractClient rpc_interact.InteractServiceImplClient
	GlobalRelationClient rpc_relation.RelationServiceImplClient
)
var (
	JaegerEndpoint 				= "127.0.0.1:4318" 
	RegisterTraceServiceName    = "Register"
	LoginTraceServiceName 		= "Login"
	UserInfoTraceServiceName 	= "UserInfo"
	FavoriteActionTraceServiceName = "FavoriteAction"
	FavoriteListTraceServiceName   = "FavoriteList"
	CommentActionTraceServiceName  = "CommentAction"
	CommentListTraceServiceName	   = "CommentList"
	RelationActionTraceServiceName = "RelationAction"
	FollowListTraceServiceName	   = "FollowList"
	FollowerListTraceServiceName   = "FollowerList"
	FriendListTraceServiceName	   = "FriendList"
)
