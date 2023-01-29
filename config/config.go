package config

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	config = new(Config)
	client *minio.Client
	err    error
	engine *gin.Engine     = gin.Default()
	ctx    context.Context = context.Background()
)

type Config struct {
	MinioConfig MinioConfig `yaml:"MinioConfig"`
	MysqlConfig MysqlConfig `yaml:"DataConfig"`
	JudgeConfig JudgeConfig `yaml:"JudgeConfig"`
	Port        string      `yaml:"Port"`
}

func Init() {
	file, errs := ioutil.ReadFile("./config.yaml")
	if errs != nil {
		log.Fatalf("解析配置文件失败 %#v", errs.Error())
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("解析配置文件失败 %#v", err.Error())
	}
	config.MysqlConfig.Init()
	config.MinioConfig.Init()
	//engine = gin.Default()
}
func GetGin() *gin.Engine {
	//if engine == nil {
	//	gin.Default()
	//}
	return engine
}
func GetMinIO() *minio.Client {
	return client
}
func GetConfig() *Config {
	return config
}
