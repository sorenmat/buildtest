package main

import (
	"database/sql"
	"fmt"
	"testing"

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
	if db.Ping() != nil {
		t.Error("unable to ping database")
	}
}
