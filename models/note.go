package models

import (
	"github.com/electro98/noteapp/utils"
)

// Would be lovely if id was random
type Note struct {
	utils.BaseModel `gorm:"embedded"`
	Title           string `json:"title"`
	Content         string `json:"content"`
}
