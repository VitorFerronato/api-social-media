package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuarios"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuário por ID"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuario"))
}
