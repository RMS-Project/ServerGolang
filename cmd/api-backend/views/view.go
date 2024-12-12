package views

import (
	"net/http"
)

// handleNotFound responde com status 404 quando o recurso não é encontrado.
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Not found", http.StatusNotFound)
}

// handleMethodNotAllowed responde com status 405 quando o método HTTP não é permitido.
func HandleMethodNotAllowed(w http.ResponseWriter) {
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}