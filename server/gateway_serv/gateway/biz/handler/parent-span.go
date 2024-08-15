package api_server

import (
	"context"
	"fmt"

	otel "go.opentelemetry.io/otel"
	attribute "go.opentelemetry.io/otel/attribute"
	t "go.opentelemetry.io/otel/trace"
)

func createRegisterParentSpan(ctx context.Context) (context.Context , t.Span) {
	// 获取一个 Tracer 实例，用于创建 span
	tracer := otel.Tracer("Register")
	baseAttrs := []attribute.KeyValue{
		attribute.String("linhaohaoaho", "zuishuai"), // 自定义属性：domain
	}
	// 开启一个父 span，传入父上下文和自定义属性
	cx, span := tracer.Start(ctx, "parent-span", t.WithAttributes(baseAttrs...),t.WithSpanKind(t.SpanKindClient))
	span.AddEvent("Parent span started", t.WithAttributes(
		attribute.String("hahhahahahhah", "????????????"),
	))
	fmt.Println("-----------")
	fmt.Printf("Parent SpanID: %s\n", span.SpanContext().SpanID().String())
	fmt.Printf("Parent TraceID: %s\n", span.SpanContext().TraceID().String())
	fmt.Println("-----------")
	return cx , span
}