package model

import (
	"github.com/thoseJanes/tinyblog/pkg/util/id"
	"gorm.io/gorm"
)


type Post struct{
	// Id int64 `gorm:"column:id;primary_key"`
	// CreatedAt time.Time `gorm:"column:createdAt"`
	// UpdatedAt time.Time `gorm:"column:UpdatedAt"`
	TimeAndId
	Username string `gorm:"column:username;not null"`
	Title string `gorm:"column:title;not null"`
	PostId string `gorm:"column:postId;not null"`
	Content string `gorm:"column:content"`
}


func (*Post) TableName() string {
	return "post_table"
}

func (p *Post) BeforeCreate(db *gorm.DB) error {
	p.PostId = "post-" + id.GenShortId()
	return nil
}