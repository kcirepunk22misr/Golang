package conexion

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "kcirepunk:1007787011@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error()) // Error Handling manejo de errores
	}
	return databaseConnection
}