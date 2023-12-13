package client

import (
	einar "archetype/app/shared/archetype/resty"
	"context"
	"archetype/app/exception"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/constants"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var ArchetypeRestyClient = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {
	
	req := einar.Client.NewRequest().SetContext(ctx)

	//Replace Get by your http method
	res, err := req.Get("http://localhost:8080/api/ping")

	ctx, span := tracer.Start(ctx, "ArchetypeRestyClient")
	defer span.End()

	if err != nil {
		networkErr := exception.HTTP_NETWORK_ERROR
		slog.
			SpanLogger(span).
			Error(networkErr.Error(),
				constants.Error, err.Error())
		span.SetStatus(codes.Error, networkErr.Error())
		span.RecordError(networkErr, trace.WithAttributes(
			attribute.String(constants.ResponseBody, constants.NoResponse),
		))
		return networkErr
	}

	if res.StatusCode() >= 500 && res.StatusCode() <= 599 {
		externalServerErr := exception.EXTERNAL_SERVER_ERROR
		slog.
			SpanLogger(span).
			Error(externalServerErr.Error(),
				constants.ResponseBody, res.Body())
		span.SetStatus(codes.Error, externalServerErr.Error())
		span.RecordError(externalServerErr, trace.WithAttributes(
			attribute.String(constants.ResponseBody, string(res.Body())),
		))
		return externalServerErr
	}

	if res.StatusCode() > 299 {
		externalUnknownErr := exception.EXTERNAL_UNKNOWN_ERROR
		slog.
			SpanLogger(span).
			Error(externalUnknownErr.Error(),
				constants.ResponseBody, res.Body())
		span.SetStatus(codes.Error, externalUnknownErr.Error())
		span.RecordError(externalUnknownErr, trace.WithAttributes(
			attribute.String(constants.ResponseBody, string(res.Body())),
		))
		return externalUnknownErr
	}

	//Handle successfull response Here

	return nil
}
