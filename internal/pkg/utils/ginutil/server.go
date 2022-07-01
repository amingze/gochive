package ginutil

import (
	"github.com/amingze/gochive/internal/pkg/utils/httputil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server
}

func NewServer(addr string) *Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: gin.Default(),
	}
	return &Server{
		srv: srv,
	}
}

func (s *Server) Start() {
	go s.Run()
}

func (s *Server) Run() {
	log.Printf("[rest server listen at %s]", s.srv.Addr)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}

func (s *Server) Stop() {
	httputil.Shutdown(s.srv)
}
