package main

import (
	"api_filmes/database"
	"api_filmes/routes"
)

func main() {
	database.ConexaoDatabase()
	routes.HandleRequest()
}
