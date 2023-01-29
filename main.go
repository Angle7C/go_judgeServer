package main

import (
	GlobalConfig "testts/config"
	"testts/controller"
)

var (
	config *GlobalConfig.Config
)

func main() {
	GlobalConfig.Init()
	gin := GlobalConfig.GetGin()
	controller.Init()
	gin.Run()
	//fmt.Println("%s\n%s", string(output), err.Error())
}
