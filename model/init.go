package model

import(
	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"fmt"
)
//DB database connection
var DB *gorm.DB

//InitDatabase function
func InitDatabase(){

	var err error
	DB,err =gorm.Open("mysql","username:password@tcp(host:port)/database's name?charset=utf8&parseTime=True&loc=Local")
	DB.LogMode(true)

	if err != nil{
		fmt.Println(err)
		panic(err)
	}

	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(time.Second * 30)
	DB.AutoMigrate(&User{},&TodoModel{})
}