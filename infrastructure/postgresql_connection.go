package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func OpenPostgreConnection() (*gorm.DB, error) {
	dbO, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=student-management password=123456")
	if err != nil {
		return nil, err
	}
	db = dbO
	db.LogMode(true)
	return db, nil
}

func GetPostgreDatabase() *gorm.DB {
	return db
}
