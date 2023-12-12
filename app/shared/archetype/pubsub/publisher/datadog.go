// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016 Datadog, Inc.

// Package pubsub provides functions to trace the cloud.google.com/pubsub/go package.
package publisher

import (
	"archetype/app/shared/constants"
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"cloud.google.com/go/pubsub"
)

// Publish publishes a message on the specified topic and returns a PublishResult.
// This function is functionally equivalent to t.Publish(ctx, msg), but it also starts a publish
// span and it ensures that the tracing metadata is propagated as attributes attached to
// the published message.
// It is required to call (*PublishResult).Get(ctx) on the value returned by Publish to complete
// the span.
func Publish(ctx context.Context, t *pubsub.Topic, msg *pubsub.Message) *PublishResult {
	span, ctx := tracer.StartSpanFromContext(
		ctx,
		"pubsub.publish",
		tracer.ResourceName(t.String()),
		tracer.SpanType(ext.SpanTypeMessageProducer),
		//tracer.Tag("message_size", len(msg.Data)),
		//tracer.Tag("ordering_key", msg.OrderingKey),
	)
	if msg.Attributes == nil {
		msg.Attributes = make(map[string]string)
	}
	if err := tracer.Inject(span.Context(), tracer.TextMapCarrier(msg.Attributes)); err != nil {

		log.Error().
			Str(constants.ERROR, err.Error()).
			Str("topic", t.String()).
			Uint64(constants.DD_TRACE_ID, span.Context().TraceID()).
			Uint64(constants.DD_SPAN_ID, span.Context().SpanID()).
			Msg("contrib/cloud.google.com/go/pubsub.v1/: failed injecting tracing attributes")

	}

	span.SetTag("num_attributes", len(msg.Attributes))
	return &PublishResult{
		PublishResult: t.Publish(ctx, msg),
		span:          span,
	}
}

// PublishResult wraps *pubsub.PublishResult
type PublishResult struct {
	*pubsub.PublishResult
	once sync.Once
	span tracer.Span
}

// Get wraps (pubsub.PublishResult).Get(ctx). When this function returns the publish
// span created in Publish is completed.
func (r *PublishResult) Get(ctx context.Context) (string, error) {
	serverID, err := r.PublishResult.Get(ctx)
	r.once.Do(func() {
		r.span.SetTag("server_id", serverID)
		r.span.Finish(tracer.WithError(err))
	})
	return serverID, err
}
