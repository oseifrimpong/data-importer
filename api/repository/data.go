package repository

import (
	"math"
	"ohlc-data-api/api/dto"
	"ohlc-data-api/api/models"

	"gorm.io/gorm"
)

type DataRepo interface {
	Create(data []*models.Data) error
	Retrieve(queryParams *dto.SearchParams) (dataList map[string]interface{}, err error)
}

type repo struct {
	db *gorm.DB
}

func NewDataRepository(db *gorm.DB) DataRepo {
	return &repo{db: db}
}

func (r *repo) Create(data []*models.Data) error {
	if err := r.db.CreateInBatches(data, len(data)).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) Retrieve(queryParams *dto.SearchParams) (dataList map[string]interface{}, err error) {

	var data []models.Data

	offset := (queryParams.PageNum - 1) * queryParams.PageSize
	sqlBuilder := r.db.Limit(queryParams.PageSize).Offset(offset).Order(queryParams.Sort)
	if err = sqlBuilder.Model(&models.Data{}).
		Where("high IN (?)", queryParams.High).
		Or("open IN (?)", queryParams.Open).
		Or("unix IN (?)", queryParams.Unix).
		Or("close IN (?)", queryParams.Close).
		Or("low IN (?)", queryParams.Low).
		Or("symbol IN (?)", queryParams.Symbol).
		Find(&data).Error; err != nil {
		return nil, err
	}

	var totalCount int64
	r.db.Model(&models.Data{}).Count(&totalCount)

	res := make(map[string]interface{})
	res["data"] = data
	res["page_size"] = queryParams.PageSize
	res["sort"] = queryParams.Sort
	res["total_pages"] = int(math.Ceil(float64(totalCount) / float64(queryParams.PageSize)))
	res["total_count"] = totalCount

	return res, nil
}
