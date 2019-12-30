package articleManagement

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
)

func TestArticle(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Errorf("gorm.Open error: %v", err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})

	for i := 0; i < 10; i++ {
		err := CreateArticle(db, Article{
			Title:        fmt.Sprintf("article%03d", i),
			Author:       fmt.Sprintf("a%d", i),
			Abstract:     "xxx",
			FilePath:     "http://001",
			InitialGrade: 111,
		})
		if err != nil {
			t.Errorf("CreateArticle i=%d error: %v", i, err)
		}
	}

	n := 5
	articles, err := GetLatestNArticle(db, n)
	if err != nil {
		t.Errorf("GetLatestNArticle n=%d error:%v", n, err)
	} else {
		t.Logf("GetLatestNArticle n=%d success: %v", n, articles)
	}
}
