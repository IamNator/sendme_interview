package repository

type Transaction interface {
	CustomQuery(query interface{}, data ...interface{}) (interface{}, error)
}
