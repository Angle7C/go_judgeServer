package service

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"testts/config"
	"testts/dal/model"
)

/*
	负责测试点文件的下载

*/
func DownloadProblem(problem model.Problem, work string) {
	//删除work/problemId/下的所有文件
	// work/problem/id/
	problemPath := work + string(os.PathSeparator) + "problem" + string(os.PathSeparator) + strconv.Itoa(problem.ProblemId) + string(os.PathSeparator)
	err := os.RemoveAll(problemPath)
	os.Create(problemPath)
	if err != nil {
		log.Fatalf("删除失败")
	}
	for i := 0; i < problem.Num; i += 1 {
		in := config.DownloadFile("problem", strconv.Itoa(i)+".in")
		writerFile(in, problemPath+strconv.Itoa(i)+".in")
		out := config.DownloadFile("problem", strconv.Itoa(i)+".out")
		writerFile(out, problemPath+strconv.Itoa(i)+".out")

	}
}
func getIndexProblem(problemId int, index uint) string {
	return strconv.Itoa(problemId) + string(os.PathSeparator) + strconv.Itoa(int(index))
}
func writerFile(reader io.Reader, filename string) {
	file, err := os.Create(filename)
	//os.Create()
	if err != nil {
		log.Printf("创建文件失败%v", err.Error())
	}
	wr := bufio.NewWriterSize(file, 1000)
	//将reader中的内容写入file
	_, err = bufio.NewReaderSize(reader, 1000).WriteTo(wr)
	if err != nil {
		log.Printf("文件写失败入%v", err.Error())
	}
}
