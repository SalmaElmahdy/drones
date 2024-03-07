package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabse() *gorm.DB {
	dsn := "host=localhost user=hms password=hms dbname=drones port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}
	fmt.Println("Database Name:", db.Migrator().CurrentDatabase())
	fmt.Println("Database DSN:", db.Dialector.Name())
	return db
}
