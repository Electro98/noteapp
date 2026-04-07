package models

type Query[T any] interface {
	// SELECT * FROM @@table WHERE id=@id
	GetById(id uint) (T, error)
}
