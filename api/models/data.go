package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Data struct {
	Base
	Unix   int64  `json:"unix" gorm:"int8;"`
	Symbol string `json:"symbol" gorm:"type:varchar(20);"`
	Open   string `json:"open" gorm:"type:varchar(25)"`
	High   string `json:"high" gorm:"type:varchar(25)"`
	Low    string `json:"low" gorm:"type:varchar(25)"`
	Close  string `json:"close" gorm:"type:varchar(25)"`
}

func (d *Data) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	d.ID = id.String()
	return nil
}
