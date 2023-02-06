package delivery

import (
	"net/http"
	"ohlc-data-api/api/dto"
	"ohlc-data-api/api/service"

	"github.com/gin-gonic/gin"
)

type dataController struct {
	svc service.DataService
}

func NewDataController(svc service.DataService) *dataController {
	return &dataController{
		svc: svc,
	}
}

// Add a new academic record
// @Summary Update user's academic details
// @Description Update user's academic details with access_token
// @Description Add Bearer prefix before Authorization value.
// @Tags profile
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer + Token"
// @Param body body dto.Academic true "Request Body"
// @Success 200 {object} dto.APIResponse
// @Failure 400,404 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Failure default {object} dto.APIResponse
// @Router /data [post]
func (e *dataController) Create(ctx *gin.Context) {
	resp := &dto.APIResponse{
		StatusCode: 2000,
		Message:    "started",
	}
	// if err := ctx.SaveUploadedFile(); err != nil {
	// 	resp := &dto.APIResponse{
	// 		StatusCode: 4000,
	// 		Message:    err.Error(),
	// 	}
	// 	ctx.SecureJSON(http.StatusBadRequest, resp)
	// 	return
	// }
	ctx.SecureJSON(http.StatusOK, resp)
}
