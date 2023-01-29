package model

import (
	"testts/dal"
)

type Run struct {
	//gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserId    uint   `gorm:"Column:user_id" json:"userId"`
	ProblemId uint   `gorm:"Column:problem_id" json:"problemId"`
	RecordId  uint   `gorm:"Column:record_id" json:"recordId"`
	Content   string `gorm:"Colum:content" json:"content"`
	Type      string `gorm:"Colum:types" json:"type"`
	State     string `gorm:"Colum:state" json:"state"`
}

func (run Run) IsWait() bool {
	return run.State == RunWait
}
func (run Run) IsRunning() bool {
	return run.State == Running
}
func (run Run) IsSuccess() bool {
	return run.State == RunSuccess
}
func (run Run) GetLangType() dal.LangAble {
	if run.Type == "C" {
		return dal.CLang{}
	}
	return nil
}
