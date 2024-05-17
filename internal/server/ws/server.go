package ws

import (
	"context"
	"net/http"
	"time"
)

type WebsocketServer struct {
	httpServer *http.Server
	port       string
	handler    http.Handler
}

func NewWebsocketServer(port string, handler http.Handler) *WebsocketServer {
	return &WebsocketServer{port: port, handler: handler}
}

func (s *WebsocketServer) Run() error {
	s.httpServer = &http.Server{
		Addr:           ":" + s.port,
		Handler:        s.handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *WebsocketServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
