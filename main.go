package main

import (
	"fmt"
	"GolangRestApi/conexion"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := conexion.InitDB();
	defer databaseConnection.Close()
	fmt.Println(databaseConnection)

}