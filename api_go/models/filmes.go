package models

type Filme struct {
	Id         int    `json:"id"`
	Nome       string `json:"nome"`
	Genero     string `json:"genero"`
	Descricao  string `json:"descricao"`
	ImagemCapa []byte `json:"imagemCapa"`
}

var Filmes []Filme
