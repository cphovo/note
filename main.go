/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/cphovo/note/cmd"
	"github.com/cphovo/note/constants"
	"github.com/cphovo/note/db"
)

func main() {
	database, err := db.GetDB(constants.DB_CONNECTION_STRING)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	cmd.Execute()
}
