package articleManagement

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model

	Title        string `gorm:"column:title"`
	Author       string `gorm:"column:author;index"`
	Abstract     string `gorm:"column:abstract"`
	FilePath     string `gorm:"column:filepath"`
	InitialGrade int    `gorm:"column:initial_grade"`
}

func CreateArticle(db *gorm.DB, article Article) (err error) {
	err = db.Create(&article).Error
	if err != nil {
		return fmt.Errorf("create artcile %v error: %v", article, err)
	} else {
		return nil
	}
}

func GetLatestNArticle(db *gorm.DB, n int) (articles []Article, err error) {
	err = db.Order("created_at desc").Limit(n).Find(&articles).Error
	if err != nil {
		return nil, fmt.Errorf("GetLatestNArticle n=%d error: %v", n, err)
	}
	return articles, nil
}
