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
	if command == "start" {
		server.Start()
		fmt.Println("Server wird gestartet")
	} else if command == "stop" {
		server.Stop()
		fmt.Println("Server wird gestoppt")
	} else {
		fmt.Println("Unbekanntes Kommando", command)
	}
}
