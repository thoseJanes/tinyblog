package model

import "time"

type TimeAndId struct{
	Id int64 `gorm:"column:id;primary_key"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}