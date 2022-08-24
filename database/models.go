package database

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB, err = gorm.Open("sqlite3", "db.db")

var dns = "creativesspace:WeWantPeace!@(db4free.net:3306)/creativesspace?charset=utf8&parseTime=True&loc=Local"
var MySQLDB, err2 = gorm.Open("mysql", dns)

type User struct {
	gorm.Model
	IsAdmin  uint8 `gorm:"DEFAULT:0"`
	Nick     string
	Password string
	Email    string
	LoginKey string
}

type AboutUser struct {
	gorm.Model
	UserId uint
	About  string
	Avatar string `gorm:"DEFAULT:'static/images/avatar.png'"`
	Cover  string `gorm:"DEFAULT:'static/images/cover.jpg'"`
}

type Video struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      uint64
	PathToVideo string
	Name        string
	Cover       string
	About       string
	Category    uint
	Views       uint
}

type Subscribers struct {
	gorm.Model
	UserId      uint
	FromId      uint
	ToId        uint
	WhenStarted *time.Time
}

type VideoEmotions struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    uint
	VideoId   uint
	Emotion   bool `gorm:"DEFAULT:false; type:boolean;"`
}

type Comments struct {
	gorm.Model
	UserId      uint64
	VideoId     uint64
	TextComment string
}

type History struct {
	gorm.Model
	UserId  uint64
	VideoId uint64
}

func Connect() {
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&VideoEmotions{})
	DB.AutoMigrate(&Subscribers{})
	DB.AutoMigrate(&AboutUser{})
	DB.AutoMigrate(&Comments{})
	DB.AutoMigrate(&History{})
	DB.AutoMigrate(&Video{})
	MySQLDB.AutoMigrate(&User{})
}

func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

type Result struct {
	Status  int
	Comment string
}

type EmbdedResults struct {
	Status  int
	Comment string
	Id      uint
	Nick    string
	Avatar  string
}

type UserData struct {
	Id    int    `json:"id"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
	About string `json:"about"`
}
