package article

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Body  []byte
}

func init() {

}
