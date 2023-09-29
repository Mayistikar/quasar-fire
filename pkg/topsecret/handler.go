package topsecret

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// TopSecretSplit for handle topsecret_split request
// @Tags TopSecretSplit
// @Summary To save information about the satellites
// @Description To save information about the satellites
// @Accept json
// @Produce json
// @Param satellite_name path string true "satellite_name" Enums(kenobi,skywalker,sato)
// @Param info body RequestTopSecretSplit true "satellite"
// @Success 200 {object} string
// @Failure 404 {object} string
// @Router /topsecret_split/{satellite_name} [post]
func (h *Handler) TopSecretSplit(ctx *gin.Context) {
	var body RequestTopSecretSplit
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusNotFound, "no se pueda determinar la posición o el mensaje")
		return
	}

	satelliteName := ctx.Param("satellite_name")

	// Validate satellite name
	if satelliteName == "" || satelliteName != "kenobi" && satelliteName != "skywalker" && satelliteName != "sato" {
		ctx.JSON(http.StatusBadRequest, "el nombre del satelite no es valido")
		return
	}

	satellite := &Satellite{
		Name:     satelliteName,
		Distance: body.Distance,
		Message:  body.Message,
	}

	if err := h.service.TopSecretSplit(satellite); err != nil {
		ctx.JSON(http.StatusNotFound, "no se pueda determinar la posición o el mensaje")
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}

// TopSecretSplitGet for handle topsecret_split request
// @Tags TopSecretSplit
// @Summary To find the location of the ship and the message
// @Description To find the location of the ship and the message
// @Accept json
// @Produce json
// @Success 200 {object} Ship
// @Failure 404 {object} string
// @Router /topsecret_split [get]
func (h *Handler) TopSecretSplitGet(ctx *gin.Context) {
	response, err := h.service.TopSecretSplitGet()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// TopSecretSplitInfo for handle topsecret_split/info request
// @Tags Additional Utils
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

// TopSecretSplitDelete for handle topsecret_split request
// @Tags Additional Utils
// @Summary To delete all information saved in the application database about the satellites
// @Description To delete all information saved in the application database about the satellites
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 404 {object} string
// @Router /topsecret_split [delete]
func (h *Handler) TopSecretSplitDelete(ctx *gin.Context) {
	if err := h.service.TopSecretSplitDelete(); err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}
