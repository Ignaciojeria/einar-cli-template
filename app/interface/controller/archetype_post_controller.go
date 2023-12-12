package controller

import (
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/echo_server"
	"archetype/app/shared/ddlog"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func init() {
	container.InjectInBoundAdapter(func() error {
		einar.Echo.POST("/INSERT_YOUR_PATTERN_HERE", archetypePostController)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func archetypePostController(c echo.Context) (err error) {
	span, _ := tracer.StartSpanFromContext(c.Request().Context(), "archetypePostController", tracer.AnalyticsRate(1))
	//REPLACE REQUEST AND RESPONSE BY YOUR CUSTOM MODEL
	var request interface{}
	var response interface{}
	//REPLACE STATUS CODE BY YOUR RETURNED STATUS CODE DEPENDING ON RESULT OF YOUR USECASE
	var statusCode int

	defer func() {
		if err != nil {
			span.SetTag("error.msg", err.Error())
			ddlog.Err(span.Context()).
				Int("statusCode", statusCode).
				Interface("request", request).
				Interface("response", response).
				Msg("INSERT_YOUR_CUSTOM_LOG_METRIC_ERROR")
			span.Finish(tracer.WithError(err))
			return
		}
		ddlog.Info(span.Context()).
			Int("statusCode", statusCode).
			Interface("request", request).
			Interface("response", response).
			Msg("INSERT_YOUR_CUSTOM_LOG_METRIC_SUCCEDDED")
		span.Finish()
	}()

	statusCode = http.StatusOK
	return c.JSON(http.StatusOK, "INSERT_YOUR_CUSTOM_RESPONSE")
}
