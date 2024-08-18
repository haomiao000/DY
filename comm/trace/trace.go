package trace

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

const (
	LocalAgentHostPort = "localhost:6831"
)

// NewTracer 使用 opentracing 统一标准
func NewTracer(service string) (opentracing.Tracer, io.Closer) {
	// config := parseConfig()
	return newTracer(service, "")
}
// newTracer
func newTracer(service, collectorEndpoint string) (opentracing.Tracer, io.Closer) {
	// 参数详解 https://www.jaegertracing.io/docs/1.20/sampling/
	cfg := jaegerConfig.Configuration{
		ServiceName: service,
		// 采样配置
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans: true,
			// CollectorEndpoint:  CollectorEndpoint2, // 将span发往jaeger-collector的服务地址
			LocalAgentHostPort: LocalAgentHostPort,
		},
	}
	// 不传递 logger 就不会打印日志
	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	// tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}