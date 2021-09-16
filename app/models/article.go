package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Body  []byte
}

func (article *Article) Create() {
	db.Create(&Article{Title: article.Title, Body: article.Body})
}
