package model

import (
	"gorm.io/gorm"
	"log"
	"testts/config"
)

type Problem struct {
	ProblemId   int    `gorm:"primaryKey"`
	ProblemName string `gorm:"Column:problem_name"`
	ProblemDesc string `gorm:"Column:problem_desc"`
	CpuTime     int64  `gorm:"Column:cpu_time"`
	MemorySize  int64  `gorm:"Column:memory_size"`
	Num         int    `gorm:"Column:num"`
}

func (problem *Problem) QueryId(id uint) {
	tx := config.GetDb().First(&problem, id)
	if tx.Error != nil {
		log.Fatalf("查询问题出错id:%v\nerr:%v", id, tx.Error.Error())
	}
}
func (problem *Problem) TableName() string {
	return "problem"
}
func (problem *Problem) UpdateId() {
	db := config.GetDb()
	db.Transaction(func(tx *gorm.DB) error {
		tx.Updates(problem)
		if tx.Error != nil {
			log.Printf("更新问题出错id:%v\nerr:%v\n", problem.ProblemId, tx.Error.Error())
			return tx.Error
		}
		return nil
	})
}
func (problem *Problem) DeleteId() {
	db := config.GetDb()
	db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(problem)
		if tx.RowsAffected > 1 {
			log.Printf("删除问题出错id:%v\nerr:%v\n", problem.ProblemId, tx.Error.Error())
			return tx.Error
		}
		return nil
	})
}
