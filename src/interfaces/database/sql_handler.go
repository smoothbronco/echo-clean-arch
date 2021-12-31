package database

type SqlHandler interface {
	Create(object interface{})
	FindAll(object interface{})
	FindOne(object interface{}, id int)
	DeleteById(object interface{}, id int)
}
