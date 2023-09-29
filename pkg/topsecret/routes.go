package topsecret

import "github.com/gin-gonic/gin"

type Routes struct {
	handler *Handler
}

func NewRoutes(handler *Handler) Routes {
	return Routes{handler}
}

func (r Routes) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/topsecret", r.handler.TopSecret)
	router.POST("/topsecret_split/:satellite_name", r.handler.TopSecretSplit)
	router.GET("/topsecret_split", r.handler.TopSecretSplitGet)
	router.GET("/topsecret_split/info", r.handler.TopSecretSplitInfo)
	router.DELETE("/topsecret_split", r.handler.TopSecretSplitDelete)
}
