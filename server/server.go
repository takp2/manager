package server

import (
	"os"
	"os/signal"
	"time"
)

// Server handles server processes
type Server struct {
}

// New creates a new server
func New() *Server {
	return &Server{}
}

// Start starts the server
func (s *Server) Start() error {
	go s.poll()
	return nil
}

func (s *Server) poll() error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	ticker := time.NewTicker(12 * time.Second)

	select {
	case <-ticker.C:
	case <-signalChan:
		return nil
	}

	return nil
}
