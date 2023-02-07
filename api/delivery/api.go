package delivery

import (
	"fmt"
	"log"
	"net/http"
	"ohlc-data-api/api/dto"
	"ohlc-data-api/api/internal/utils"
	"ohlc-data-api/api/service"
	"os"

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
// @Tags v1
// @Accept x-www-form-urlencoded
// @Produce json
// @Param formData body dto.Academic true "Request Body"
// @Success 200 {object} dto.APIResponse
// @Failure 400,404,422 {object} dto.APIResponse
// @Failure default {object} dto.APIResponse
// @Router /data [post]
func (d *dataController) Create(ctx *gin.Context) {
	response := &dto.APIResponse{}
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	if !utils.ValidateFile(file) {
		response.StatusCode = 4000
		response.Message = "file not supported, upload a csv file"
		response.Data = make([]interface{}, 0)
		ctx.SecureJSON(http.StatusBadRequest, response)
		return
	}

	path, _ := os.Getwd()
	targetDestination := fmt.Sprintf("%s%s/%s", path, os.Getenv("MEDIA_PATH"), file.Filename)

	err = ctx.SaveUploadedFile(file, targetDestination)
	if err != nil {
		log.Fatal(err)
	}

	response, err = d.svc.Create(targetDestination)
	if err != nil {
		response.StatusCode = 4000
		response.Message = "failed to process file"
		response.Data = make([]interface{}, 0)
		ctx.SecureJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response = &dto.APIResponse{
		StatusCode: 2000,
		Message:    "file uploaded and it's processing",
	}
	ctx.SecureJSON(http.StatusOK, response)
}
