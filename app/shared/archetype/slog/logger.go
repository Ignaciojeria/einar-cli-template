package slog

import (
	"log/slog"
	"os"
	"strconv"

	"go.opentelemetry.io/otel/trace"
)

// Datadog trace and log correlation :
// https://docs.datadoghq.com/tracing/other_telemetry/connect_logs_and_traces/opentelemetry/?tab=go
const (
	ddTraceIDKey = "dd.trace_id"
	ddSpanIDKey  = "dd.span_id"
	ddServiceKey = "dd.service"
	ddEnvKey     = "dd.env"
	ddVersionKey = "dd.version"
)

// Default opentelemetry trace and log correlation :
const (
	traceIDKey = "trace_id"
	spanIDKey  = "span_id"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func SpanLogger(span trace.Span) *slog.Logger {
	traceID := span.SpanContext().TraceID().String()
	ddTraceIDValue := convertTraceID(traceID)
	spanID := span.SpanContext().SpanID().String()
	ddSpanIDValue := convertTraceID(spanID)
	return logger.
		With(
			slog.String(traceIDKey, traceID),
			slog.String(spanID, spanID),
			slog.String(ddTraceIDKey, ddTraceIDValue),
			slog.String(ddSpanIDKey, ddSpanIDValue),
			slog.String(ddServiceKey, os.Getenv("DD_SERVICE")),
			slog.String(ddEnvKey, os.Getenv("DD_ENV")),
			slog.String(ddVersionKey, os.Getenv("DD_VERSION")),
		)
}

func Logger() *slog.Logger {
	return logger
}

func convertTraceID(id string) string {
	if len(id) < 16 {
		return ""
	}
	if len(id) > 16 {
		id = id[16:]
	}
	intValue, err := strconv.ParseUint(id, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatUint(intValue, 10)
}
