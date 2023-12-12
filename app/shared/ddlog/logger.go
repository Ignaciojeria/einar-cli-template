package ddlog

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

// https://confluence.falabella.tech/display/LC/ARCH+-+Logging+Guideline
var ddLogger zerolog.Logger

func init() {
	container.InjectOutBoundAdapter(func() error {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.LevelFieldName = "severity"
		zerolog.TimestampFieldName = "timeMills"
		ddLogger = zerolog.New(os.Stdout).
			With().
			Str("timeMills", strconv.FormatInt(time.Now().Unix(), 10)).
			Interface("contextMap", map[string]string{
				"industry":   "rtl",
				"commerce":   "corp",
				"country":    config.COUNTRY.Get(),
				"domain":     "schn",
				"capability": "trmg",
				"product":    "dispatch",
			}).
			Interface("resource", map[string]string{
				"service.name":    config.PROJECT_NAME.Get(),
				"service.version": "1.0.0",
				"k8s.pod.uid":     os.Getenv("HOSTNAME"),
			}).
			Timestamp().
			Logger()
		return nil
	}, container.InjectionProps{DependencyID: uuid.NewString()})

}

func Info(ctx ddtrace.SpanContext) *zerolog.Event {
	return retrieveSpanAndTraceFromSpanContext(ctx, ddLogger.Info())
}

func Warn(ctx ddtrace.SpanContext) *zerolog.Event {
	return retrieveSpanAndTraceFromSpanContext(ctx, ddLogger.Warn())
}

func Err(ctx ddtrace.SpanContext) *zerolog.Event {
	return retrieveSpanAndTraceFromSpanContext(ctx, ddLogger.Error())
}

func retrieveSpanAndTraceFromSpanContext(ctx ddtrace.SpanContext, e *zerolog.Event) *zerolog.Event {
	return e.
		Str("dd.trace_id", strconv.FormatUint(ctx.TraceID(), 10)).
		Str("dd.span_id", strconv.FormatUint(ctx.SpanID(), 10))
}
