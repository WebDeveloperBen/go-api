package presence

import (
	"github.com/labstack/echo/v4"
)

type Router struct {
	Group   *echo.Group
	Handler *PresenceHandler
}

func NewPresenceRouter(group *echo.Group, handler *PresenceHandler) {
	r := &Router{
		Group:   group,
		Handler: handler,
	}
	r.AttachRoutes()
}

// NOTE: Don't use trailing slashes on routes with Echo
func (r *Router) AttachRoutes() {
	r.Group.GET("/presence", r.Handler.HandleGetPresences)
	r.Group.GET("/presence/:id", r.Handler.HandleGetPresence)
	r.Group.POST("/presence", r.Handler.HandleCreatePresence)
	r.Group.PATCH("/presence/:id", r.Handler.HandleUpdatePresence)
	r.Group.DELETE("/presence/:id", r.Handler.HandleDeletePresence)
}
