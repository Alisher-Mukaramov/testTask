package server

import (
	"github.com/gin-gonic/gin"
	"interviewTask/internal/handler"
	"net/http"
	"os"
	"time"
)

const version = "v1"

type Server struct {
	handler *handler.Handler
	engine  *gin.Engine
}

func New(h *handler.Handler) *Server {
	return &Server{handler: h}
}

func (s *Server) Run() {
	s.engine = gin.New()
	s.engine.Use(gin.Recovery())
	s.engine.Use(gin.Logger())

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if host == "" {
		host = "127.0.0.1"
	}

	if port == "" {
		port = "8080"
	}

	ep := NewEndpoint(version, s)
	ep.Init()

	srv := &http.Server{
		Addr:           port,
		Handler:        s.engine,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	srv.ListenAndServe()
	s.engine.Run()
}
