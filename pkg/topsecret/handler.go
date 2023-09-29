package topsecret

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestTopSecret struct {
	Satellites []Satellite `json:"satellites"`
}

// Handler for topsecret handler
type Handler struct {
	service Service
}

// NewHandler create a new handler instance
func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// TopSecret for handle topsecret request
// @Tags TopSecret
// @Summary To find the location of the ship and the message
// @Description To find the location of the ship and the message
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param info body RequestTopSecret true "satellites"
// @Success 200 {object} Ship
// @Failure 404 {object} string
// @Router /topsecret [post]
func (h *Handler) TopSecret(ctx *gin.Context) {
	var body RequestTopSecret
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusNotFound, "no se pueda determinar la posición o el mensaje")
		return
	}

	response, err := h.service.TopSecret(body.Satellites)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "no se pueda determinar la posición o el mensaje")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// TopSecretSplitInfo for handle topsecret_split/info request
// @Tags TopSecretSplit
// @Summary To get all information saved in the application database about the satellites
// @Description To get all information saved in the application database about the satellites
// @Accept json
// @Produce json
// @Success 200 {object} []Satellite
// @Failure 404 {object} string
// @Router /topsecret_split/info [get]
func (h *Handler) TopSecretSplitInfo(ctx *gin.Context) {
	info, err := h.service.TopSecretSplitInfo()
	if err != nil {
		ctx.JSON(http.StatusNotFound, "no se pudo encontrar información relacionada a los satelites")
		return
	}
	ctx.JSON(http.StatusOK, info)
}
