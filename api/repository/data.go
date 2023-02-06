package repository

import (
	"ohlc-data-api/api/models"

	"gorm.io/gorm"
)

type DataRepo interface {
	Create(data *models.Data) error
	Retrieve() (dataList []*models.Data, err error)
}

type repo struct {
	db *gorm.DB
}

func NewDataRepository(db *gorm.DB) DataRepo {
	return &repo{db: db}
}

func (r *repo) Create(data *models.Data) error {

	return nil
}

func (r *repo) Retrieve() (dataList []*models.Data, err error) {

	return nil, nil
}
