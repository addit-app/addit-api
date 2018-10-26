package addit

import (
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.GET("/", HelloRoot)
	e.POST("/api/v1/url", PostUrlIndex)
}