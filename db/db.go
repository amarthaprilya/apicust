package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := "freedb_customeruser:pHQQcE@ttPyB#5Z@tcp(sql.freedb.tech:3306)/freedb_customerdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	runMigration(db)
	return db
}

func runMigration(db *sql.DB) {
	file, err := os.ReadFile("ddl.sql")
	if err != nil {
		log.Fatal("Failed to open ddl.sql", err)
	}

	statements := strings.Split(string(file), ";")
	for _, statement := range statements {
		trimmed := strings.TrimSpace(statement)
		if trimmed != "" {
			_, err := db.Exec(trimmed)
			if err != nil {
				log.Fatal("Failed to execute migration: ", err)
			}
		}
	}
	log.Println("Database migration completed successfully")
}
