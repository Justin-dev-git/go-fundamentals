package main

import (
	"fmt"

)
type Server struct{
	Name string
	IP string
	Active  bool
}
func (s *Server) Stop(){
	s.Active=false
}
func (s *Server) Start(){
	s.Active=true
}
func main(){
	servers:=[]Server{
		{Name:"Server1", IP: "127.0.0.1", Active: true},
		{Name:"Server2", IP: "192.168.0.1", Active: true},
	}
	for _, s:= range servers{
		fmt.Println(s.Name, s.Active)
	}
	servers[0].Stop()
	fmt.Println(servers[0].Name, servers[0].Active)
}
