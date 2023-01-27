package untils

type MysqlConfig struct {
	Url      string `yaml:"Url"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"PassWord"`
	DataName string `yaml:"DataName"`
}

var ()

func (mysqlConfig MysqlConfig) Init() {

}
