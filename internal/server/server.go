package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Wookkie/notes-g2/internal"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg       *internal.Config
	httpServe *http.Server
}

func New(cfg *internal.Config) *Server {
	httpServe := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	myServer := Server{
		httpServe: &httpServe,
		cfg:       cfg,
	}
	myServer.configRoutes()

	return &myServer
}

func (s *Server) configRoutes() {
	router := gin.Default()
	router.GET("/")
	users := router.Group("/users")
	{
		users.GET("/porfile")
		users.POST("/register")
		users.POST("/login", s.login)
	}

	notes := router.Group("/notes")
	{
		notes.GET("/")
		notes.GET("/:id")
		notes.POST("/")
		notes.PUT("/:id")
		notes.DELETE("/:id")
	}

	s.httpServe.Handler = router
}

func (s *Server) Run() error {
	return s.httpServe.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServe.Shutdown(ctx)
}
