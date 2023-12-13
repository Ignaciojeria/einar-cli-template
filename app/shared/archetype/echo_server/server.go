package echo_server

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/config"
	"archetype/app/shared/constants"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

var e *echo.Echo

func init() {
	config.Installations.EnableHTTPServer = true

	container.InjectInstallation(func() error {
		e = echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(otelecho.Middleware(config.PROJECT_NAME.Get() + "-http-server"))
		return nil
	})

	container.InjectHTTPServer(func() error {
		setUpRenderer(EmbeddedPatterns...)
		for _, route := range e.Routes() {
			fmt.Printf("Method: %v, Path: %v, Name: %v\n", route.Method, route.Path, route.Name)
		}
		err := e.Start(":" + config.PORT.Get())
		if err != nil {
			slog.
				Logger().
				Error("error initializing application server",
					constants.Error, err.Error())
			return err
		}
		return nil
	})
}

func Echo() *echo.Echo {
	if e == nil {
		e = echo.New()
	}
	return e
}
