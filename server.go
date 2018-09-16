package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger *log.Logger

func _setupHandler() *Handler {

	logger = log.New(os.Stdout, "", 0)

	handler := &Handler{mux: http.NewServeMux()}
	handler.logger = logger

	return handler
}

// Handler -- server handler
type Handler struct {
	logger *log.Logger
	mux    *http.ServeMux
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", "example Go server")

	h.mux.ServeHTTP(res, req)
}

func _startServer(server *http.Server) {
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal(err)
	}
}

// StartServer -- run the server
func StartServer(handler *Handler) {
	server := &http.Server{Addr: ":3000", Handler: handler}
	logger.Printf("Starting server on port 3000 ...")
	go _startServer(server)

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	logger.Printf("\nShutting down\n")

	if err := server.Shutdown(ctx); err != nil {
		logger.Printf("Error: %v\n", err)
	} else {
		logger.Println("Server stopped")
	}
}

// SetupServer -- exported func for setup
func SetupServer() *Handler {
	return _setupHandler()
}
