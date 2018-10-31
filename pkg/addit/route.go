package addit

import (
	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	e.GET("/", HelloRoot)
	e.POST("/api/v1/url", PostUrlIndex)
	e.GET("/api/v1/url/:hash", GetUrlIndex)
	e.POST("/api/v1/index", PostIndex)
	e.GET("/api/v1/index/:urlhash", GetIndex)
}