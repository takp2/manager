package server

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/takp2/manager/mtea"
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
	p := tea.NewProgram(mtea.New())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
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
