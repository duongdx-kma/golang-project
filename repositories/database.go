package repositories

type SqlHandler interface {
	Create(object interface{})
	FindAll(object interface{})
	FindById(object interface{}, id string)
}
