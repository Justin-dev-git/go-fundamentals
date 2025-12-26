package main

import (
	"fmt"
	"os"
)

type Server struct {
	Name   string
	Active bool
}

func (s *Server) Start() {
	s.Active = true
}
func (s *Server) Stop() {
	s.Active = false
}
func main() {
	server := &Server{Name: "Server 1"}
	if len(os.Args) < 2 {
		fmt.Println("Usage:app start|stop")
		return
	}
	command := os.Args[1]
	switch command {
	case "Start":
		server.Start()
		fmt.Println("Server wird gestartet")
	case "Stop":
		server.Stop()
		fmt.Println("Server wird gestoppt")
	case "Status":
		fmt.Println(server.Active)
	default:
		fmt.Println("Unbekanntes Kommando")
	}

}
