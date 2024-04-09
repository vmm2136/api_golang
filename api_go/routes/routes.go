package routes

import (
	"api_filmes/controllers"
	"api_filmes/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/api/listarFilmes", controllers.ListarFilmes).Methods("GET")
	r.HandleFunc("/api/listarFilmes/{id}", controllers.ListarFilmeEspecifico).Methods("GET")
	r.HandleFunc("/api/cadastrarFilme", controllers.CadastrarFilme).Methods("POST")
	r.HandleFunc("/api/deletarFilme/{id}", controllers.DeletarFilme).Methods("DELETE")
	r.HandleFunc("/api/editarFilme/{id}", controllers.EditarFilme).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))

}
