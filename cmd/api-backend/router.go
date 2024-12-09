package server

import (
	"net/http"

	"github.com/RodrigoMS/App_Go/cmd/api-backend/controllers"
)

var ( 
	userHandlers = map[string]func(http.ResponseWriter, *http.Request){ 
		"GET": controllers.GetUser, 
		"POST": controllers.PostUser, 
		"PUT": controllers.PutUser, 
		"DELETE": controllers.DeleteUser, 
		} 
	
	mutex = &sync.RWMutex{}
)


func handleNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

func routes(mux *http.ServeMux) {
	
	mux.HandleFunc("GET /user", controllers.GetUser)
	mux.HandleFunc("POST /user", controllers.PostUser)
	mux.HandleFunc("PUT /user", controllers.PutUser)
	mux.HandleFunc("DELETE /user", controllers.DeleteUser)
	mux.HandleFunc("/user", handleNotFound)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	/*var userHandlers = map[string]func(http.ResponseWriter, *http.Request){
		"GET":    controllers.GetUser,
		"POST":   controllers.PostUser,
		"PUT":    controllers.PutUser,
		"DELETE": controllers.DeleteUser,
	}*/

	mutex.RLock()
	handler, ok := userHandlers[r.Method]
	mutex.RUnlock()

	if ok {
		handler(w, r)

	} else {
		// Status 405 - Unsupported Method.
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}


