package web

import (
	_ "github.com/beriloqueiroz/desafio-back/core/docs/swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type HandlerFuncMethod struct {
	HandleFunc http.HandlerFunc
	Method     string
}

type WebServer struct {
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddRoute(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() error {
	mux := http.NewServeMux()
	s.AddRoute("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	for path, handler := range s.Handlers {
		mux.Handle(path, handler)
	}
	return http.ListenAndServe(s.WebServerPort, mux)
}

type output struct {
	Message string
}
