package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg       any
	httpServe *http.Server
}

func New(host string, port string) *Server {
	httpServe := http.Server{
		Addr: host + ":" + port,
	}
	//fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	router := gin.Default()
	configRoutes(router)

	httpServe.Handler = router

	return &Server{
		httpServe: &httpServe,
		//cfg:       cfg,
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
