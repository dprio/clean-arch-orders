package webserver

import (
	"fmt"
	"net/http"

	"github.com/dprio/clean-arch-orders/internal/infrastructure/config"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	router        chi.Router
	handlers      map[string]http.HandlerFunc
	webServerPort string
}

type route struct {
	method string
	path   string
}

func New(webConfig *config.Web) *WebServer {
	return &WebServer{
		router:        chi.NewRouter(),
		handlers:      make(map[string]http.HandlerFunc),
		webServerPort: webConfig.Port,
	}
}

func (ws *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	ws.handlers[path] = handler
}

func (ws *WebServer) Start() error {
	for path, handler := range ws.handlers {
		ws.router.Handle(path, handler)
	}

	return http.ListenAndServe(fmt.Sprintf("localhost:%s", ws.webServerPort), ws.router)
}
