package models

import "github.com/marceloagmelo/cursogomvc/lib"

//Usuarios representa a tabela de usuários na base de dados
type Usuarios struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}

//UsuarioModel recebe a tabela do banco de dados
var UsuarioModel = lib.Sess.Collection("usuarios")
