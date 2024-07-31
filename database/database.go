package database

import (
	"educative-rest-api-course/models"
	"educative-rest-api-course/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// create the variable to store the database instance.
// Using gorm library

var DB *gorm.DB

// init database creates a connection to the database
func InitDatabase(dbName string) {

	// initialize some variables for the
	// MySQL datastore
	var (
		databaseUser     string = utils.GetValue("DB_USER")
		databasePassword string = utils.GetValue("DB_PASSWORD")
		databaseHost     string = utils.GetValue("DB_HOST")
		databasePort     string = utils.GetValue("DB_PORT")
		databaseName     string = dbName
	)

	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	var err error

	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the database")

	DB.AutoMigrate(&models.User{}, &models.Item{})
}
