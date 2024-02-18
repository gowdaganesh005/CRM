package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func Connect() (*gorm.DB, error) {

	godotenv.Load()
	dburl := os.Getenv("DBURL")
	db, err := sql.Open("pgx", dburl)
	if err != nil {
		log.Fatalf("could not connect to database:%v", err)
	}
	gormdb, err := gorm.Open(postgres.New(
		postgres.Config{
			Conn: db,
		}), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database:%v", err)
	}

	return gormdb, err

}
