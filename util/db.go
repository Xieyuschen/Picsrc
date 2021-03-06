package util

import (
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	"log"
)
type Image struct {
	gorm.Model
	Url string
	IsDelete bool
	Tag string
}
type DbSettings struct{
	Username string `json:"Username"`
	Password string	`json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}
type JsonSettings struct {
	DbSettings DbSettings `json:"DbSettings"`
	Token string `json:"Token"`
}
var ImageType Image
var db *gorm.DB
var Settings JsonSettings
func dsn(settings DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", settings.Username,settings.Password, settings.Hostname,settings.Dbname)
}

func init(){
	Settings = ReadSetting("Config.json")
	connStr := dsn(Settings.DbSettings)
	db, _ = gorm.Open("mysql",connStr)
	if !db.HasTable(&ImageType){
		db.CreateTable(&ImageType)
		log.Println("Create Image table successfully")
	}
}
func AddImage(image *Image){
	db.Create(image)
}
func GetImages(page int, pageSize int,tag string)(images []Image){
	db.Where("tag=?",tag).Limit(pageSize).Offset(page* pageSize).Find(&images)
	return
}
func DeleteImage(id int)(err error){
	err=db.Where("id=?",id).Delete(&Image{}).Error
	return
}