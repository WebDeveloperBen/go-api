package assets

import (
	"github.com/labstack/echo/v4"
)

type Router struct {
	Group   *echo.Group
	Handler *AssetsHandler
}

func NewAssetRouter(group *echo.Group, handler *AssetsHandler) {
	r := &Router{
		Group:   group,
		Handler: handler,
	}
	r.AttachRoutes()
}

// NOTE: Don't use trailing slashes on routes with Echo
func (r *Router) AttachRoutes() {
	r.Group.GET("/assets", r.Handler.HandleGetAllAssets)
	r.Group.GET("/assets/public", r.Handler.HandleGetPublicAssets)
	r.Group.GET("/assets/:id", r.Handler.HandleGetAssetByID)
	r.Group.GET("/assets/file/:fileName", r.Handler.HandleGetAssetByFileName)
	r.Group.POST("/assets", r.Handler.HandleCreateAsset)
	r.Group.PATCH("/assets/:id", r.Handler.HandleUpdateAsset)
	r.Group.DELETE("/assets/:id", r.Handler.HandleDeleteAsset)
	r.Group.GET("/assets/count", r.Handler.HandleGetAssetsCount)
}
