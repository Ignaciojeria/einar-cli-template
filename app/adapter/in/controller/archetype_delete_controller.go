package controller

import (
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/echo_server"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func init() {
	container.InjectInBoundAdapter(func() error {
		einar.Echo.DELETE("/INSERT_YOUR_PATTERN_HERE", archetypeDeleteController)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func archetypeDeleteController(c echo.Context) error {
	return c.JSON(http.StatusOK, "INSERT_YOUR_CUSTOM_RESPONSE")
}
