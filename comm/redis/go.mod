module github.com/haomiao000/DY/comm/redis

go 1.22

replace github.com/haomiao000/DY/server/redis_svr => ../../server/redis_svr

replace github.com/haomiao000/DY/comm/discovery => ../discovery

require (
	github.com/haomiao000/DY/comm/discovery v0.0.0-00010101000000-000000000000
	github.com/haomiao000/DY/server/redis_svr v0.0.0-20240728141946-64cee6b12469
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	go.etcd.io/etcd/api/v3 v3.5.15 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.15 // indirect
	go.etcd.io/etcd/client/v3 v3.5.15 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240730163845-b1a4ccb954bf // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240730163845-b1a4ccb954bf // indirect
)
