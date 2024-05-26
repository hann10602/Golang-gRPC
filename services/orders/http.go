package main

import (
	"log"
	"net/http"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}