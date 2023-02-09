package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"ohlc-data-api/api/dto"
	"ohlc-data-api/api/internal/utils"
	"ohlc-data-api/api/models"
	"ohlc-data-api/api/repository"
	"os"
	"strconv"

	"go.uber.org/zap"
)

type DataService interface {
	Create(targetDestination string) (res *dto.APIResponse, err error)
	Retrieve(queryParams *dto.SearchParams) (res *dto.APIResponse, err error)
}

type dataSvc struct {
	repo   repository.DataRepo
	logger *zap.Logger
}

func NewDataService(r repository.DataRepo, l *zap.Logger) DataService {
	return &dataSvc{repo: r, logger: l}
}

func (d *dataSvc) Create(file string) (res *dto.APIResponse, err error) {

	csvDataList, err := extractData(file)
	if err != nil {
		d.logger.Error(utils.CREATION_FAILED, zap.Error(err))
		response := &dto.APIResponse{
			StatusCode: utils.CREATION_FAILED_CODE,
			Message:    utils.CREATION_FAILED,
			Data:       make([]interface{}, 0),
		}
		return response, err
	}

	batchData := make([]*models.Data, 0)

	for _, csvData := range csvDataList {

		data := &models.Data{}

		unix, _ := strconv.ParseInt(csvData[0], 10, 64)
		data.Unix = unix

		data.Symbol = csvData[1]
		data.Open = csvData[2]
		data.High = csvData[3]
		data.Low = csvData[4]
		data.Close = csvData[5]

		batchData = append(batchData, data)
		batchSize, _ := strconv.Atoi(os.Getenv("BATCH_SIZE"))

		if len(batchData) == batchSize {
			err = d.repo.Create(batchData)
			if err != nil {
				d.logger.Error(utils.CREATION_FAILED, zap.Error(err))
				response := &dto.APIResponse{
					StatusCode: utils.CREATION_FAILED_CODE,
					Message:    utils.CREATION_FAILED,
					Data:       make([]interface{}, 0),
				}
				return response, err
			}
			batchData = []*models.Data{}
			continue
		}
		continue
	}

	// insert remaining data from batch data
	err = d.repo.Create(batchData)
	if err != nil {
		d.logger.Error(utils.CREATION_FAILED, zap.Error(err))
		response := &dto.APIResponse{
			StatusCode: utils.CREATION_FAILED_CODE,
			Message:    utils.CREATION_FAILED,
			Data:       make([]interface{}, 0),
		}
		return response, err
	}

	response := &dto.APIResponse{
		StatusCode: utils.SUCCESSFUL_CODE,
		Message:    utils.SUCCESSFUL,
		Data:       make([]interface{}, 0),
	}

	return response, nil
}

func extractData(file string) ([][]string, error) {
	fmt.Println("file path: ", file)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	headers, err := csvReader.Read()
	if errors.Is(err, io.EOF) {
		return nil, err
	}

	if err = utils.ValidateCSVHeaders(headers); err != nil {
		return nil, err
	}

	csvDataList, err := csvReader.ReadAll()
	if errors.Is(err, io.EOF) {
		return nil, err
	}
	return csvDataList, nil
}

func (d *dataSvc) Retrieve(queryParams *dto.SearchParams) (res *dto.APIResponse, err error) {
	var response dto.APIResponse

	results, err := d.repo.Retrieve(queryParams)
	if err != nil {
		d.logger.Error(utils.RETRIEVE_FAILED, zap.Error(err))

		response.StatusCode = utils.RETRIEVE_FAILED_CODE
		response.Message = utils.RETRIEVE_FAILED
		response.Data = make([]interface{}, 0)

		return &response, err
	}

	response.StatusCode = utils.SUCCESSFUL_CODE
	response.Message = utils.SUCCESSFUL
	response.Data = results

	return &response, nil
}
