package model

import (
	"log"
	"testts/config"
)

type Record struct {
	RecordId   uint   `gorm:"Column:record_id" json:"recordId"`
	NumId      int    `gorm:"Column:unm_id" json:"numId"`
	CpuTime    int64  `gorm:"Column:cpu_time" json:"cpuTime"`
	MemorySize int64  `gorm:"Column:memory_size" json:"memorySize"`
	State      string `gorm:"Column:state" json:"state"`
	err        string `gorm:"Column:error" json:"err"`
}

func ToRecord(result Result, index int) Record {
	var record = Record{
		RecordId:   0,
		NumId:      index,
		CpuTime:    result.CpuTime,
		MemorySize: 0,
		State:      result.State,
		err:        "",
	}
	if result.State == "RealTimeLimitExceeded" || result.State == "CpuTimeLimitExceeded" {
		record.State = "TLE"
	} else if result.State == "RuntimeError" {
		record.State = "RE"
	} else if result.State == "SystemError" {
		record.State = "WE"
	} else if result.State == "MemoryLimitExceeded" {
		record.State = "MLE"
	} else {
		record.State = "AC"
	}
	return record
}
func SaveRunRecord(list []Record) {
	db := config.GetDb()
	tx := db.Create(list)
	if tx.Error != nil {
		log.Fatalf("保存运行结果失败")
	}
}
