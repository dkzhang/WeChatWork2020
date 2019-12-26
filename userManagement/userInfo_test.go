package userManagement

import (
	"github.com/jinzhu/gorm"
	"testing"
)

func TestInitTableFromFile(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	users, err := InitTableFromFile(db, `./PersonInfoForTest.csv`)
	if err != nil {
		t.Errorf("InitTableFromFile error: %v", err)
	} else {
		t.Logf("InitTableFromFile success. Create %d users.", len(users))
	}
}
