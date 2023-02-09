package delivery

import (
	"fmt"
	"net/http"
	"ohlc-data-api/api/dto"
	"ohlc-data-api/api/internal/utils"
	"ohlc-data-api/api/service"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type dataController struct {
	svc    service.DataService
	logger *zap.Logger
}

func NewDataController(svc service.DataService, log *zap.Logger) *dataController {
	return &dataController{
		svc:    svc,
		logger: log,
	}
}

// Upload Data
// @Summary Uploads csv data into database
// @Description Uploads csv data into database
// @Tags v1
// @Accept x-www-form-urlencoded
// @Produce json
// @Param file formData file true "Request Body"
// @Success 200 {object} dto.APIResponse
// @Failure 400,404,422 {object} dto.APIResponse
// @Failure default {object} dto.APIResponse
// @Router /data [post]
func (d *dataController) Create(ctx *gin.Context) {
	response := &dto.APIResponse{}
	file, err := ctx.FormFile("file")
	if err != nil {
		d.logger.Error(utils.CREATION_FAILED, zap.Error(err))

		response.StatusCode = utils.CREATION_FAILED_CODE
		response.Message = utils.NOT_SUPPORTED_FILE
		response.Data = make([]interface{}, 0)

		ctx.SecureJSON(http.StatusBadRequest, response)
		return
	}

	if !utils.ValidateFile(file) {
		d.logger.Error(utils.CREATION_FAILED, zap.Error(err))

		response.StatusCode = utils.CREATION_FAILED_CODE
		response.Message = utils.NOT_SUPPORTED_FILE
		response.Data = make([]interface{}, 0)

		ctx.SecureJSON(http.StatusBadRequest, response)
		return
	}

	path, _ := os.Getwd()
	targetDestination := fmt.Sprintf("%s%s/%s", path, os.Getenv("MEDIA_PATH"), file.Filename)

	err = ctx.SaveUploadedFile(file, targetDestination)
	if err != nil {
		d.logger.Error(utils.CREATION_FAILED, zap.Error(err))

		response.StatusCode = utils.CREATION_FAILED_CODE
		response.Message = utils.UPLOAD_FAILED
		response.Data = make([]interface{}, 0)

		ctx.SecureJSON(http.StatusBadRequest, response)
		return
	}

	response, err = d.svc.Create(targetDestination)
	if err != nil {
		ctx.SecureJSON(http.StatusUnprocessableEntity, response)
		return
	}

	d.logger.Info(utils.SUCCESSFUL)
	ctx.SecureJSON(http.StatusOK, response)
}

// GET
// @Summary Search all data
// @Description Search data by using the query parameters
// @Tags v1
// @Accept json
// @Produce json
// @Param page_num query string false "page_num"
// @Param page_size query string false "page_size"
// @Param sort query string false "sort_field ASC | DESC"
// @Success 200 {object} dto.APIResponse
// @Failure 400,404 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Failure default {object} dto.APIResponse
// @Router /v1/data [get]
func (d *dataController) Retrieve(ctx *gin.Context) {
	args := &dto.SearchParams{}
	args.Sort = ctx.DefaultQuery("sort", "created_at DESC")
	args.PageNum, _ = strconv.Atoi(ctx.DefaultQuery("page_num", "0"))
	args.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("page_size", "50"))

	args.High = ctx.DefaultQuery("high", "")
	args.Low = ctx.DefaultQuery("low", "")
	args.Unix, _ = strconv.ParseInt(ctx.DefaultQuery("unix", ""), 10, 64)
	args.Open = ctx.DefaultQuery("open", "")
	args.Close = ctx.DefaultQuery("close", "")
	args.Symbol = ctx.DefaultQuery("symbol", "")

	if err := utils.SearchValidation(args); err != nil {
		d.logger.Error(utils.RETRIEVE_FAILED, zap.Error(err))

		response := &dto.APIResponse{
			StatusCode: utils.RETRIEVE_FAILED_CODE,
			Message:    utils.RETRIEVE_FAILED,
			Data:       make([]interface{}, 0),
		}

		ctx.SecureJSON(http.StatusBadRequest, response)
		return
	}

	res, err := d.svc.Retrieve(args)
	if err != nil {
		ctx.SecureJSON(http.StatusNoContent, res)
		return
	}

	d.logger.Info(utils.SUCCESSFUL)
	ctx.SecureJSON(http.StatusOK, res)
}
