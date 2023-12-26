package configs

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)
import "gorm.io/driver/mysql"
import "golang-restful-api/models"

var DB *gorm.DB

func ConnectDatabase() {
	godotenv.Load()
	var connectionString string = os.Getenv("MYSQL_USERNAME") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(localhost:3306)/" + os.Getenv("MYSQL_DBNAME")
	db, err := gorm.Open(mysql.Open(connectionString))
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Product{})

	DB = db
}
