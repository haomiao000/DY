package api_server

import (
	"context"
	"fmt"
	t "go.opentelemetry.io/otel/trace"
	otel "go.opentelemetry.io/otel" 
	attribute "go.opentelemetry.io/otel/attribute"
)

func TransmitSpan(ctx context.Context) (context.Context , t.Span) {
	tracer := otel.Tracer("Register")
	ctx, iSpan := tracer.Start(ctx, fmt.Sprintf("span-%d", 1))
	iSpan.AddEvent("Child span processing", t.WithAttributes(
		attribute.String("child-span", fmt.Sprintf("span-loggggggggggggg%d", 1)),
	))
	return ctx , iSpan
}