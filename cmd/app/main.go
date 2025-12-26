package main

import (
	"fmt"
	"log"
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
		fmt.Println("Usage:app start|stop|status")
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "start":
		server.Start()
		log.Println("Server wird gestartet")
	case "stop":
		server.Stop()
		log.Println("Server wird gestoppt")
	case "status":
		log.Println(server.Name, "Status", server.Active)
	default:
		log.Println("Unbekanntes Kommando")
		os.Exit(1)
	}
	os.Exit(0)

}
