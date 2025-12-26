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
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
func main() {
	serverName := getEnv("SERVER_NAME", "default-server")
	server := &Server{Name: serverName}
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
		fmt.Println(server.Name, "Status", server.Active)
	default:
		fmt.Println("Unbekanntes Kommando")
	}

}
