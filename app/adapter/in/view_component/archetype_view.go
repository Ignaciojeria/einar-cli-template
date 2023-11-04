package ArchetypeView

import (
	"archetype/app/adapter/in/view/component"
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/echo_server"
	"embed"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:embed *.html
var html embed.FS

//go:embed *.css
var css embed.FS

func init() {
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: html,
		Pattern: archetype_constant + component.DOT_HTML,
	})
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: css,
		Pattern: archetype_constant + component.DOT_CSS,
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/"+archetype_constant, render)
		einar.Echo.GET("/"+archetype_constant+component.DOT_CSS, echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func render(c echo.Context) error {
	routerState := einar.NewRoutingState(c, map[string]string{
		component.IndexComponentDefault: archetype_constant,
	})
	return c.Render(http.StatusOK, archetype_constant+component.DOT_HTML, routerState)
}
