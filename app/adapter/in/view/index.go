package view

import (
	"archetype/app/adapter/in/view/component"
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/echo_server"
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed *.html
var html embed.FS

//go:embed *.css
var css embed.FS

func init() {
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: html,
		Pattern: component.Index + component.DOT_HTML,
	})
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: css,
		Pattern: component.Index + component.DOT_CSS,
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/", render)
		einar.Echo.GET("/"+component.Index+component.DOT_CSS, echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	})
}

func render(c echo.Context) error {
	routerState := einar.NewRoutingState(c, map[string]string{
		component.IndexComponentDefault: component.App,
	})
	return c.Render(http.StatusOK, component.Index+component.DOT_HTML, routerState)
}
