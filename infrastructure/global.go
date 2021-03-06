package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ErrLog *log.Logger
var InfoLog *log.Logger
var db *gorm.DB

func init() {
	//InfoColor    := "\033[1;34m%s\033[0m"
	//ca:=color.New(color.FgCyan)
	InfoLog = log.New(os.Stdout, "\u001B[1;34m[INFO]\u001B[0m ", log.Ldate|log.Ltime|log.Llongfile)
	//color.Set(color.FgRed)
	ErrLog = log.New(os.Stderr, "\033[1;31m[ERROR]\033[0m ", log.Ldate|log.Ltime|log.Llongfile)
	DevelopEnv = getStringEnvParameter(DevelopEnv, "local")

	if DevelopEnv == "local1" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		log.Print(DevelopEnv == "local1")
	}
	loadEnvParameters()

	dbO, err := OpenPostgreConnection()
	if err != nil {
		InfoLog.Print("Erro: ", err)
	}
	db = dbO
	/*	err = db.AutoMigrate(&model.User{}, &model.Student{}, &model.Employer{}, &model.School{}, &model.Post{}).Error
		db.Model(&model.Student{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.Employer{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.School{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")*/

}
