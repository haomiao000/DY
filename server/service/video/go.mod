module github.com/haomiao000/DY/server/service/video

go 1.22.3

replace github.com/haomiao000/DY/server/grpc_gen => ../../grpc_gen

replace github.com/haomiao000/DY/server/common => ../../common

replace github.com/haomiao000/DY => ../../../..

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/haomiao000/DY/server/grpc_gen v0.0.0-00010101000000-000000000000 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	gorm.io/gorm v1.25.11 // indirect
	gorm.io/plugin/opentelemetry v0.1.4 // indirect
)
