package models

import (
	"log"
	"time"

	"../helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var DB *gorm.DB

func init() {
	initDB()
}

func initDB() {
	//获取数据库配置,没有配置直接panic
	db_conf, err := helper.ReadJsonFile("./config/database.json")

	if err != nil {
		panic("json file not found")
	}

	db, err := gorm.Open("mysql", db_conf["db_user"]+":"+db_conf["db_password"]+"@/"+db_conf["db_name"]+"?parseTime=true&loc=Local")

	log.Println(db_conf["db_user"] + ":" + db_conf["db_password"] + "@/" + db_conf["db_name"] + "?parseTime=true")
	if err != nil {
		log.Println(err)
	}

	DB = db
}
