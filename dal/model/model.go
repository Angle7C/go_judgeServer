package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"Column:user_name"`
	Email    string `gorm:"Column:email"`
	PassWord string `gorm:"Column:password"`
	Avatar   string `gorm:"Column:avatar"`
}
type Comments struct {
	CommentsId int    `gorm:"primaryKey"`
	Title      string `gorm:"Column:title"`
	Context    int64  `gorm:"Column:context"`
	UserId     uint   `gorm:"Column:user_id"`
}

type Enter struct {
	RaceId uint `gorm:"Column:race_id"`
	UserId uint `gorm:"Column:user_id"`
}

type Race struct {
	RaceId    int    `gorm:"primaryKey"`
	RaceName  string `gorm:"Column:race_name"`
	RaceDesc  string `gorm:"Column:race_desc"`
	StartTime int64  `gorm:"Column:start_time"`
	EndTime   int64  `gorm:"Column:end_time"`
}

type Category struct {
	ProblemId uint `gorm:"Column:problem_id"`
	TagId     uint `gorm:"Column:tag_id"`
}
