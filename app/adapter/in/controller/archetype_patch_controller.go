package controller

import (
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/echo_server"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func init() {
	container.InjectInboundAdapter(func() error {
		einar.Echo.PATCH("/api/insert_your_pattern_here", archetypePatchController)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func archetypePatchController(c echo.Context) error {
	return c.JSON(http.StatusOK, "insert_your_custom_response")
}
