package sql

type DataInter interface {
	queryId()
	SaveModel(model interface{})
	DeleteModel(model interface{})
	UpdateModel(model interface{})
}
