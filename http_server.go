package gohelper

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type HTTPServer struct {
	logger  *log.Logger
	server  *http.Server
	handler *http.Handler
}

func NewHTTPServer(handler http.Handler, addr string) *HTTPServer {
	logger := log.New(os.Stdout, "HTTP: ", log.LstdFlags)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &HTTPServer{
		logger:  logger,
		server:  server,
		handler: &handler,
	}
}

func (H HTTPServer) ListenAndServer() {
	if e := H.server.ListenAndServe(); e != http.ErrServerClosed {
		H.logger.Fatalf("Could not listen on %s: %v\n", H.server.Addr, e)
	}
}

func (H HTTPServer) GracefullyShutDown(quit <-chan os.Signal, done chan<- bool) {
	<-quit
	H.logger.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	H.server.SetKeepAlivesEnabled(false)
	if err := H.server.Shutdown(ctx); err != nil {
		H.logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

func (H HTTPServer) Run() {
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go H.GracefullyShutDown(quit, done)
	H.logger.Println("Server is ready to handle requests at", H.server.Addr)
	H.ListenAndServer()
	<-done
	H.logger.Println("Server stopped")
}

type HTTPHandler struct {
}
