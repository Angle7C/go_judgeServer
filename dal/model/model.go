package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"Column:user_name"`
	Email    string `gorm:"Column:email"`
	PassWord string `gorm:"Column:password"`
	Avatar   string `gorm:"Column:avatar"`
}
type Problem struct {
	gorm.Model
	ProblemName string `gorm:"Column:problem_name"`
	ProblemDesc string `gorm:"Column:problem_desc"`
	CpuTime     int64  `gorm:"Column:cpu_time"`
	MemorySize  int64  `gorm:"Column:memory_size"`
	Num         int32  `gorm:"Column:num"`
}
type Comments struct {
	gorm.Model
	Title   string `gorm:"Column:title"`
	Context int64  `gorm:"Column:context"`
	UserId  uint   `gorm:"Column:user_id"`
}

type Enter struct {
	RaceId uint `gorm:"Column:race_id"`
	UserId uint `gorm:"Column:user_id"`
}

type Race struct {
	gorm.Model
	RaceName  string `gorm:"Column:race_name"`
	RaceDesc  string `gorm:"Column:race_desc"`
	StartTime int64  `gorm:"Column:start_time"`
	EndTime   int64  `gorm:"Column:end_time"`
}
type Record struct {
	RecordId   uint   `gorm:"Column:record_id"`
	NumId      uint   `gorm:"Column:unm_id"`
	CpuTime    int64  `gorm:"Column:cpu_time"`
	MemorySize int64  `gorm:"Column:memory_size"`
	State      string `gorm:"Column:state"`
	err        string `gorm:"Column:error"`
}
type Category struct {
	ProblemId uint `gorm:"Column:problem_id"`
	TagId     uint `gorm:"Column:tag_id"`
}

type Run struct {
	gorm.Model
	UserId    uint `gorm:"Column:user_id"`
	ProblemId uint `gorm:"Column:problem_id"`
	RecordId  uint `gorm:"Column:record_id"`
}
type Tag struct {
	gorm.Model
	TagName uint `gorm:"Column:tag_name"`
}
