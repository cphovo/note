/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/cphovo/note/cmd"
	"github.com/cphovo/note/constants"
	"github.com/cphovo/note/db"
)

func main() {
	db.InitDB(constants.DB_CONNECTION_STRING)
	cmd.Execute()
}
