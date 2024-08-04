package configs
import(
	rpc_user "main/server/grpc_gen/rpc_user"
	rpc_interact "main/server/grpc_gen/rpc_interact"
	rpc_relation "main/server/grpc_gen/rpc_relation"
)

var (
	GlobalUserClient   rpc_user.UserServiceImplClient
	GlobalInteractClient rpc_interact.InteractServiceImplClient
	GlobalRelationClient rpc_relation.RelationServiceImplClient
)

var (
	UserServerAddress = "127.0.0.1:8081"
	InteractServerAddress = "127.0.0.1:8082"
	RelationServerAddress = "127.0.0.1:8083"
)