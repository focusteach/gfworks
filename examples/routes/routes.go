package routes

import (
	"net/http"

	"github.com/focusteach/gfworks/bus"
	v1 "github.com/focusteach/gfworks/examples/routes/api/v1"
	"github.com/focusteach/gfworks/server/web"
	"github.com/gin-gonic/gin"
)

func InitRouter(s web.Engine) {
	s.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	s.Router.GET("/quit", func(c *gin.Context) {
		app := bus.GetInstance()
		c.String(http.StatusOK, "Bye from Gin Server")

		app.Quit()
	})

	g := s.Router.Group("/api/v1")
	{
		g.GET("/faces", v1.GetTopFaces)
		g.POST("/upload", v1.Upload)
	}
}
