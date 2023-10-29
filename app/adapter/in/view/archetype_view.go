package controller

import (
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
		Pattern: "html_archetype_view.html",
	})
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: css,
		Pattern: "css_archetype_view.css",
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/endpoint_archetype_view", render)
		einar.Echo.GET("/styles/css_archetype_view.css", echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func render(c echo.Context) error {
	data := map[string]interface{}{
		"componentName": "endpoint_archetype_view",
	}
	return c.Render(http.StatusOK, "html_archetype_view.html", data)
}
