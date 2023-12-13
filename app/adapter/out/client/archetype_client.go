package client

import (
	"archetype/app/exception"
	einar "archetype/app/shared/archetype/resty"
	"archetype/app/shared/constants"
	"context"
	"errors"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var ArchetypeRestyClient = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	req := einar.Client.NewRequest().SetContext(ctx)
	//Replace Get by your http method
	res, err := req.Get("http://localhost:8080/api/ping")

	ctx, span := tracer.Start(ctx, "ArchetypeRestyClient")
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			if errors.Is(err, exception.HTTP_NETWORK_ERROR) {
				span.RecordError(err, trace.WithAttributes(
					attribute.String(constants.ResponseBody, constants.NoResponse),
				))
			}
			span.RecordError(err, trace.WithAttributes(
				attribute.String(constants.ResponseBody, string(res.Body())),
			))
		}
		span.End()
	}()

	if err != nil {
		err = exception.HTTP_NETWORK_ERROR
		return err
	}

	if res.StatusCode() >= 500 && res.StatusCode() <= 599 {
		err = exception.EXTERNAL_SERVER_ERROR
		return err
	}

	if res.StatusCode() > 299 {
		err = exception.EXTERNAL_UNKNOWN_ERROR
		return err
	}

	//Handle successfull response Here
	return nil
}
