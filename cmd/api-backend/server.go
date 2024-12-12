package server

import (
  "net/http"
)

func Start() {
  // Configura as rotas usando a função routes
  mux := routes()

  // Inicia o servidor na porta 8080 com o roteador mux
  http.ListenAndServe(":8080", mux) 
}
