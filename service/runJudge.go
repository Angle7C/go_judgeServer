package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"testts/config"
	"testts/dal"
	"testts/dal/model"
)

/*
 负责代码的编译，执行。
 当前数据库存在run记录,但是等待执行
*/
//编译
func compiler(run model.Run, config *config.Config) bool {
	lang := run.GetLangType()
	target := getTarget(run, config)
	src := getSrc(run, config, lang)
	command := exec.Command(config.JudgeConfig.CompilerPath, run.Content, config.JudgeConfig.JudgePath, lang.GetCmd(src, target), config.JudgeConfig.Work, strconv.Itoa(int(run.ID)), lang.GetType())
	output, err := command.CombinedOutput()
	if err != nil {
		log.Printf("%v编译执行失败,err:%v", run.ID, err.Error())
		return false
	}

	result := model.GetResult(output)
	if result.State == "Access" {
		log.Printf("编译成功")

		return true
	} else {
		return false
	}
}

// 获取可执行程序的位置
func getTarget(run model.Run, config *config.Config) string {
	target := config.JudgeConfig.Work + string(os.PathSeparator) + strconv.Itoa(int(run.ID)) + string(os.PathSeparator) + "Main"
	return target
}
func getSrc(run model.Run, config *config.Config, lang dal.LangAble) string {
	src := config.JudgeConfig.Work + string(os.PathSeparator) + strconv.Itoa(int(run.ID)) + string(os.PathSeparator) + "compiler/Main." + lang.GetSuffix()
	return src
}

// 执行
func Run(run model.Run, config *config.Config) []model.Record {
	check(run)
	if compiler(run, config) == false {
		return nil
	}
	//查询问题的信息
	var problem model.Problem
	problem.QueryId(run.ProblemId)
	//下载测式文件
	DownloadProblem(problem, config.JudgeConfig.Work)
	//执行测试点并存储
	list := runTest(run, config, problem)
	//更新到数据库
	model.SaveRunRecord(list)
	return list
}

// 是否符合编译执行条件
func check(run model.Run) bool {
	if run.IsRunning() || run.IsSuccess() {
		return false
	} else {
		return true
	}
}
func runTest(run model.Run, config *config.Config, problem model.Problem) []model.Record {
	exe := getTarget(run, config)
	filePath := config.JudgeConfig.Work + string(os.PathSeparator) + strconv.Itoa(int(run.ProblemId)) + string(os.PathSeparator)
	outfilePath := config.JudgeConfig.Work + string(os.PathSeparator) + strconv.Itoa(int(run.ID)) + string(os.PathSeparator)
	var recordsList = make([]model.Record, problem.Num)
	for i := 0; i < problem.Num; i++ {
		inPath := filePath + strconv.Itoa(i) + ".in"
		outPath := outfilePath + strconv.Itoa(i) + ".out"
		errPath := outfilePath + strconv.Itoa(i) + ".err"
		logPath := outfilePath + strconv.Itoa(i) + ".log"
		var item = runSingle(config.JudgeConfig.JudgePath, exe, problem.CpuTime, problem.CpuTime*3.0, problem.MemorySize, inPath, outPath, errPath, logPath, i)
		compareSingle(&item, config.JudgeConfig.ComparePath, filePath+strconv.Itoa(i)+".out", outPath)
		recordsList = append(recordsList, item)
	}
	return recordsList
}
func runSingle(judge, exe string, cpuTime, realTime, memory int64, in, out, logs, err string, i int) model.Record {
	arg := fmt.Sprintf("--cpu_max=%d --cpu_real=%d --memory=%d --exe=%s --input=%s --out=%s --error=%s --log=%s", cpuTime, realTime, memory, exe, in, out, err, logs)
	output, er := exec.Command(judge, arg).CombinedOutput()
	var result model.Result
	if er != nil {
		log.Printf("执行测试点失败,err:%v\n", er.Error())
		result = model.Result{
			CpuTime:  0,
			RealTime: 0,
			Memory:   0,
			Signal:   0,
			ExitCode: 1,
			Error:    1,
			State:    "SystemError",
		}
	} else {
		_ = json.Unmarshal(output, &result)
	}
	return model.ToRecord(result, i)

}
func compareSingle(record *model.Record, compare, ans, out string) {
	output, _ := exec.Command(compare, ans, out).CombinedOutput()
	if string(output) == "0" {
		record.State = "AC"
	} else {
		record.State = "WA"
	}
}
