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

	router := gin.Default()
	configRoutes(router)

	httpServe.Handler = router

	return &Server{
		httpServe: &httpServe,
		cfg:       cfg,
	}
}

func configRoutes(router *gin.Engine) {
	router.GET("/")
	users := router.Group("/users")
	{
		users.GET("/porfile")
		users.POST("/register")
		users.POST("/login")
	}

	notes := router.Group("/notes")
	{
		notes.GET("/")
		notes.GET("/:id")
		notes.POST("/")
		notes.PUT("/:id")
		notes.DELETE("/:id")
	}
}

func (s *Server) Run() error {
	return s.httpServe.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServe.Shutdown(ctx)
}
