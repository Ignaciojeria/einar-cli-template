package client

import (
	"archetype/app/domain/ports/out"
	einar "archetype/app/shared/archetype/resty"
	"archetype/app/shared/ddlog"
	"context"

	"github.com/go-resty/resty/v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var ArchetypeRestyClient out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "ArchetypeRestyClient", tracer.AnalyticsRate(1))
	defer func() {
		if err != nil {
			span.SetTag("error.msg", err.Error())
			ddlog.
				Err(span.Context()).
				Err(err).
				Msg("ArchetypeRestyClient" + "Error")
			span.Finish(tracer.WithError(err))
		} else {
			span.Finish()
		}
	}()

	var _ *resty.Client = einar.Client
	//PUT YOUR HTTP OPERATION USING EINAR HERE :
	//....einar.Client
	return nil
}
