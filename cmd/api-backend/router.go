package server

import (
	"net/http" // Pacote para manipulação de HTTP.

	"github.com/RodrigoMS/App_Go/cmd/api-backend/controllers" // Importa os controladores.
	"github.com/RodrigoMS/App_Go/cmd/api-backend/views"
)

func routes() *http.ServeMux {

    router := http.NewServeMux()

	router.HandleFunc("GET /user", controllers.GetUser)
	router.HandleFunc("POST /user", controllers.PostUser)
	router.HandleFunc("PUT /user", controllers.PutUser)
	router.HandleFunc("DELETE /user", controllers.DeleteUser)
	router.HandleFunc("/", views.HandleNotFound)

    return router
}