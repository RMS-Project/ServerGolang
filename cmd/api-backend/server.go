package server

import (
  "net/http"
)

func Start() {

  routes()

  http.ListenAndServe(":8080", nil)
}