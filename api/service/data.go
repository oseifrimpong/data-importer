package service

import "ohlc-data-api/api/dto"

type DataService interface {
	// Create()
}

type dataSvc struct {
}

func NewDataService() DataService {
	return &dataSvc{}
}

func Create(file string) (res *dto.APIResponse, err error) {
	return res, nil
}
