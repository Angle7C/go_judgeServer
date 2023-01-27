package untils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	MinioConfig MinioConfig `yaml:"MinioConfig"`
	MysqlConfig MysqlConfig `yaml:"DataConfig"`
}

func (config *Config) Init() {
	file, errs := ioutil.ReadFile("F:\\Dept\\testts\\config.yaml")
	if errs != nil {
		log.Fatalf("解析配置文件失败 %#v", errs.Error())
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("解析配置文件失败 %#v", err.Error())
	}
}
