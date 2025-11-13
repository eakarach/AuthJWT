package database

import (
	"fmt"
	"strconv"

	"github.com/eakarach/AuthJWT/config"
	"github.com/eakarach/AuthJWT/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var gdb *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	gdb, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	gdb.AutoMigrate(&models.User{})
	fmt.Println("Database Migrated")
}
