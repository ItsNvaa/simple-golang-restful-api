package configs

import "gorm.io/gorm"
import "gorm.io/driver/mysql"
import "golang-restful-api/models"

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/devify"))
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Product{})

	DB = db
}
