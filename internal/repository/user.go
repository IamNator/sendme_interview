package repository

type User interface {
	CustomQuery(query interface{}, data ...interface{}) (interface{}, error)
}
