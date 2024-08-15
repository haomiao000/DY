package trace

import (
	"context"
	"time"

	otel "go.opentelemetry.io/otel" 
	otlptracehttp "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	resource "go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func SetUpTracer(ctx context.Context , endpoint string , serviceName string) (func(context.Context) error , error) {
	tracerProvider, err := newTraceProvider(ctx , endpoint , serviceName)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(tracerProvider)
	return tracerProvider.Shutdown, nil
}

func newTraceProvider(ctx context.Context , endpoint string , serviceName string) (*traceSDK.TracerProvider, error) {
	exp , err := otlptracehttp.New(ctx,
		// 指定 Jaeger 端点
		otlptracehttp.WithEndpoint(endpoint), 		
		otlptracehttp.WithInsecure()) 
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
	if err != nil {
		return nil, err
	}
	// 创建 Tracer Provider 并配置采样器、批处理器和资源
	traceProvider := traceSDK.NewTracerProvider(
		traceSDK.WithResource(res), 
		 // 策略为始终采样
		traceSDK.WithSampler(traceSDK.AlwaysSample()),          
		// 设置批处理器，定时发送数据到 Jaeger
		traceSDK.WithBatcher(exp, traceSDK.WithBatchTimeout(time.Second)), 
	)
	return traceProvider, nil
}