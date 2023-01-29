package model

import (
	"encoding/json"
	"log"
)

type Result struct {
	CpuTime  int64  `json:"cpuTime"`
	RealTime int64  `json:"realTime"`
	Memory   int64  `json:"memory"`
	Signal   int    `json:"signal"`
	ExitCode int    `json:"exitCode"`
	Error    int    `json:"error"`
	State    string `json:"result"`
}

func GetResult(str []byte) Result {
	var val Result
	err := json.Unmarshal(str, &val)
	if err != nil {
		log.Fatalf("judge 执行成功，无法将结果json转换 err:%v", err.Error())
	}
	return val
}
