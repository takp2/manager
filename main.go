package main

import (
	"fmt"
	"os"

	"github.com/takp2/manager/server"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func run() error {
	s := server.New()
	err := s.Start()
	if err != nil {
		return fmt.Errorf("server start: %w", err)
	}
	return nil
}
