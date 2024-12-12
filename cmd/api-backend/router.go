package server

import (
	"net/http" // Pacote para manipulação de HTTP.
	"sync"     // Pacote para controle de concorrência.

	"github.com/RodrigoMS/App_Go/cmd/api-backend/controllers" // Importa os controladores.
	"github.com/RodrigoMS/App_Go/cmd/api-backend/views"
)

// HandlerFuncMap é um mapa que associa métodos HTTP a funções manipuladoras.
// type HandlerFuncMap map[string]http.HandlerFunc

// Declarar o map como uma variável global que será inicializada uma vez.
var (
    userHandlers = map[string]http.HandlerFunc{
        "GET":    controllers.GetUser,
        "POST":   controllers.PostUser,
        "PUT":    controllers.PutUser,
        "DELETE": controllers.DeleteUser,
    }

    // Mutex para garantir a segurança em acessos simultâneos.
    mutex sync.RWMutex
)

// userHandler lida com as requisições HTTP para recursos de usuário.
func userHandler(w http.ResponseWriter, r *http.Request) {
    mutex.RLock() // Adquire o mutex para leitura.
    handler, ok := userHandlers[r.Method] // Verifica se há um manipulador para o método HTTP da requisição.
    defer mutex.RUnlock() // Libera o mutex após a leitura.

    if ok {
        handler(w, r) // Chama o manipulador correspondente.
    } else {
        // Status 405 - Método não suportado.
        views.HandleMethodNotAllowed(w)
    }
}

// modifyHandlers adiciona ou modifica um manipulador no map.
func modifyHandlers(method string, handler http.HandlerFunc) {
    mutex.Lock() // Adquire o mutex para escrita.
    userHandlers[method] = handler // Modifica ou adiciona o manipulador no map.
    defer mutex.Unlock() // Libera o mutex após a escrita.
}

func routes() {

	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/", views.HandleNotFound)

	// Exemplo de como modificar o map
    //modifyHandlers("PATCH", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("PATCH handler")) })

}