package userManagement

import (
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type User struct {
	gorm.Model

	Name           string `gorm:"column:name"`
	UserID         string `gorm:"column:user_id;unique_index"` //微信中的ID
	SapID          string `gorm:"column:sapid"`
	IdNumber       string `gorm:"column:id_number"` //身份证号
	Birthday       int    `gorm:"column:birthday"`  //19830701
	Gender         string `gorm:"column:gender"`
	Mobile         string `gorm:"column:mobile"`
	Email          string `gorm:"column:email"`
	Department     string `gorm:"column:department"`
	Position       string `gorm:"column:position"`
	PositionWeight int    `gorm:"column:position_weight"`

	//显示的姓名，例如：陈翔（主任）、马迅飞（信息技术室主任）、李华松（信息技术室）
	NameDisplay string `gorm:"column:name_display"`
}

func InitTableFromFile(db *gorm.DB, csvFilePath string) (users map[string]User, err error) {
	users = make(map[string]User)
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("open csc file: %s error: %v", csvFilePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reader.ReadAll error: %v", err)
	}

	for _, item := range record {
		//从scv文件中逐条解析
		user := User{
			Model:          gorm.Model{},
			Name:           item[0],
			UserID:         item[1],
			SapID:          item[2],
			IdNumber:       item[3],
			Birthday:       0,
			Gender:         item[5],
			Mobile:         item[6],
			Email:          item[7],
			Department:     item[8],
			Position:       item[9],
			PositionWeight: 0,
			NameDisplay:    "",
		}

		birthday, err := strconv.Atoi(item[4])
		if err != nil {
			return nil, fmt.Errorf("Birthday strconv.Atoi  error, Name = %s, Birthday = %s: %v", user.Name, item[4], err)
		} else {
			user.Birthday = birthday
		}

		//合成显示姓名
		user.NameDisplay = fmt.Sprintf("%s(%s %s)", user.Name, user.Department, user.Position)

		//插入map
		users[user.UserID] = user
	}

	//写入数据库
	if err := CreateUsers(db, users); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("create users in db error")
		return nil, fmt.Errorf("create users in db error: %v", err)
	} else {
		log.WithFields(log.Fields{
			"count": len(users),
		}).Info("create users in db success")
	}

	return users, nil
}

func CreateUsers(db *gorm.DB, users map[string]User) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, user := range users {
			if err := tx.Create(&user).Error; err != nil {
				return fmt.Errorf("create user %s %s error: %v", user.UserID, user.Name, err)
			}
		}
		return nil
	})
}
