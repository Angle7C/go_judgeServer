package model

type Tag struct {
	TagId   int  `gorm:"primaryKey"`
	TagName uint `gorm:"Column:tag_name"`
}
