package server

import (
	"AppMagicTestTask/internal/service"
	"net/http"
	"time"
)

func NewServer(addr string, srv *service.Service) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", srv.GetAllResults)

	s := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s
}
