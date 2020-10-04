package server

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/xerrors"
)

type Server struct {
	*http.Server
	localFile string
}

func New(localFilePath string) *Server {
	s := &Server{localFile: localFilePath}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", s.handleHello)
	mux.HandleFunc("/gocon", s.handleGocon)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	s.Server = &http.Server{
		Addr:    ":8999",
		Handler: mux,
	}

	return s
}

func (s *Server) handleHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello GoConference20!")
}

func (s *Server) handleGocon(w http.ResponseWriter, _ *http.Request) {
	b, err := ioutil.ReadFile(s.localFile)
	if err != nil {
		log.Printf("Failed load file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (s *Server) Start() error {
	log.Printf("Start server %s", s.Server.Addr)
	err := s.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (s *Server) Stop() error {
	return s.Server.Shutdown(context.Background())
}
