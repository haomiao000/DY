module github.com/haomiao000/DY/server/service/interact

go 1.22.3

replace github.com/haomiao000/DY/server/grpc_gen => ../../grpc_gen

replace github.com/haomiao000/DY/server/common => ../../common

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/haomiao000/DY/server/common v0.0.0-00010101000000-000000000000
	github.com/haomiao000/DY/server/grpc_gen v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.11
	gorm.io/plugin/opentelemetry v0.1.4
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240730163845-b1a4ccb954bf // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
