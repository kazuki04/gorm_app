package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Body  []byte
}

func (article *Article) Create() *Article {
	article = &Article{Title: article.Title, Body: article.Body}
	db.Create(&article)
	return article
}

func (article *Article) Update() {
	db.Model(&article).Updates(Article{Title: article.Title, Body: article.Body})
}

func FindArticle(id int) Article {
	var article Article
	db.Table("articles").First(&article, id).Scan(&article)
	return article
}
