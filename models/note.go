package models

import (
	"github.com/electro98/noteapp/utils"
)

// Would be lovely if id was random
type Note struct {
	utils.BaseModel
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Query[T any] interface {
	// SELECT * FROM @@table WHERE id=@id
	GetById(id uint) (T, error)
}
