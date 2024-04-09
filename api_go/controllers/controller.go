package controllers

import (
	"api_filmes/database"
	"api_filmes/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ListarFilmes(w http.ResponseWriter, r *http.Request) {
	var f []models.Filme
	database.DB.Find(&f)
	json.NewEncoder(w).Encode(f)
}

func ListarFilmeEspecifico(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var filme models.Filme

	database.DB.First(&filme, id)
	json.NewEncoder(w).Encode(filme)
}

func CadastrarFilme(w http.ResponseWriter, r *http.Request) {
	var novoFilme models.Filme
	json.NewDecoder(r.Body).Decode(&novoFilme)
	database.DB.Create(&novoFilme)
	json.NewEncoder(w).Encode(novoFilme)
}

func DeletarFilme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var filme models.Filme

	database.DB.Delete(&filme, id)
	json.NewEncoder(w).Encode(filme)
}

func EditarFilme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var filme models.Filme

	database.DB.First(&filme, id)
	json.NewDecoder(r.Body).Decode(&filme)

	database.DB.Save(&filme)
	json.NewEncoder(w).Encode(filme)

}
