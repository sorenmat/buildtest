package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestDatabase(t *testing.T) {
	DB_USER := "postgres"
	DB_PASSWORD := "postgres"
	DB_NAME := "postgres"
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	pinged := false
	counter := 0
	for !pinged {
		if db.Ping() == nil {
			pinged = true
			fmt.Println("Database pinged !!")
		}
		time.Sleep(1 * time.Second)
		counter++
		if counter >= 10 {
			t.Error("not able to ping db")
		}
	}

}
