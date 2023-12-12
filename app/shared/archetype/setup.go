package archetype

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"context"
	"fmt"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// ARCHETYPE CONFIGURATION
func Setup() error {

	if !config.Installations.EnableCobraCli {
		if err := config.Setup(); err != nil {
			return err
		}
	}

	if err := config.Setup(); err != nil {
		return err
	}
	ctx := context.Background()
	tp, err := tracerProvider(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			fmt.Println(err.Error())

		}
	}(ctx)

	if err := InjectInstallations(); err != nil {
		return err
	}

	if err := injectOutboundAdapters(); err != nil {
		return err
	}

	if err := injectInboundAdapters(); err != nil {
		return err
	}

	if !config.Installations.EnableHTTPServer {
		return nil
	}
	if err := container.HTTPServerContainer.LoadDependency(); err != nil {
		return err
	}
	return nil
}

func InjectInstallations() error {
	for _, v := range container.InstallationsContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}

// CUSTOM INITIALIZATION OF YOUR DOMAIN COMPONENTS
func injectOutboundAdapters() error {
	for _, v := range container.OutboundAdapterContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}

func injectInboundAdapters() error {
	for _, v := range container.InboundAdapterContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}

func tracerProvider(ctx context.Context) (*tracesdk.TracerProvider, error) {
	client := otlptracegrpc.NewClient()
	exporter, err := otlptrace.New(ctx, client)

	if err != nil {
		return nil, fmt.Errorf("creating OTLP trace exporter: %w", err)
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(os.Getenv("DD_SERVICE")),
		)),
	)
	return tp, nil
}
