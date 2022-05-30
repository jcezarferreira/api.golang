package main

import (
	"fmt"

	"golang.api/database"
	"golang.api/routes"
)

func main() {

	// models.Personalities = []models.Personality{
	// 	{Id: 1, Name: "name 1", About: "Blei"},
	// 	{Id: 2, Name: "name 2", About: "Blou"},
	// }

	database.ConnectDB()
	fmt.Println("iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
