package server

import (
  "net/http"
)

func Start() {
  mux := http.NewServeMux()

  routes(mux)

  http.ListenAndServe(":8080", mux)
}