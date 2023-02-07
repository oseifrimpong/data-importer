package service

import (
	"encoding/csv"
	"errors"
	"io"
	"ohlc-data-api/api/dto"
	"ohlc-data-api/api/models"
	"ohlc-data-api/api/repository"
	"os"
	"strconv"
)

type DataService interface {
	Create(targetDestination string) (res *dto.APIResponse, err error)
	// Get(map[string]interface{})
}

type dataSvc struct {
	repo repository.DataRepo
}

func NewDataService(r repository.DataRepo) DataService {
	return &dataSvc{repo: r}
}

func (d *dataSvc) Create(file string) (res *dto.APIResponse, err error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, errors.New("failed to load uploaded file")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	_, err = csvReader.Read() //TODO: validate headers
	if errors.Is(err, io.EOF) {
		return nil, errors.New("csv file is missing headers")
	}

	csvDataList, err := csvReader.ReadAll()
	if errors.Is(err, io.EOF) {
		return nil, errors.New("failed to read csv file")
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
		if len(batchData) == 100 {
			err = d.repo.Create(batchData)
			if err != nil {
				return nil, errors.New("failed to create data from csv")
			}
			continue
		}
		continue
	}

	// insert remaining data from batch data
	err = d.repo.Create(batchData)
	if err != nil {
		return nil, errors.New("failed to create data from csv")
	}

	return res, nil
}
