package controllers

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>GET - User</h1>")
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>POST - User</h1>")
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>PUT - User</h1>")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>DELETE - User</h1>")
}