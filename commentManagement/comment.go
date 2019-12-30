package commentManagement

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model

	ArticleID         uint   `gorm:"column:article_id;index"`
	CommentatorUserID string `gorm:"column:commentator_user_id;index"`
	Grade             int    `gorm:"column:abstract"`
	Content           string `gorm:"column:filepath"`
}
