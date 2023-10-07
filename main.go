/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/cphovo/note/cmd"
	"github.com/cphovo/note/db"
)

const (
	dsn = "data.db"
)

func main() {
	db.InitDB(dsn)
	cmd.Execute()
}
