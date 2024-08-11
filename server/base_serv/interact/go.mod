module github.com/haomiao000/DY/server/base_serv/interact

go 1.22.5

toolchain go1.22.6

replace github.com/haomiao000/DY/server/redis_svr => ../../redis_svr

replace github.com/haomiao000/DY/server/common => ../../common

replace github.com/haomiao000/DY/comm/redis => ../../../comm/redis

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/haomiao000/DY/comm/redis v0.0.0-00010101000000-000000000000
	github.com/haomiao000/DY/internal/grpc_gen v0.0.0-20240807131301-3036cdff1630
	github.com/haomiao000/DY/server/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.11
	gorm.io/plugin/opentelemetry v0.1.4
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/haomiao000/DY/comm/discovery v0.0.0-20240807131301-3036cdff1630 // indirect
	github.com/haomiao000/DY/server/redis_svr v0.0.0-00010101000000-000000000000 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	go.etcd.io/etcd/api/v3 v3.5.15 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.15 // indirect
	go.etcd.io/etcd/client/v3 v3.5.15 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240730163845-b1a4ccb954bf // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240730163845-b1a4ccb954bf // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
