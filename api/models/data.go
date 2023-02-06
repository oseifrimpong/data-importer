package models

import (
	"ohlc-data-api/api/internal/utils"

	"gorm.io/gorm"
)

type Data struct {
	Base
	Unix   int64   `json:"unix"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
}

func (d *Data) BeforeCreate(tx *gorm.DB) error {
	id, err := utils.GenerateSnowflakeID()
	if err != nil {
		return err
	}
	d.ID = id.Int64()
	return nil
}
