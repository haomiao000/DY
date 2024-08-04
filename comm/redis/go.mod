module github.com/haomiao000/DY/comm/redis

go 1.22.5

replace github.com/haomiao000/DY/server/redis_svr => ../../server/redis_svr

require (
	github.com/haomiao000/DY/server/redis_svr v0.0.0-20240728141946-64cee6b12469
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)
