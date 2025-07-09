package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Wookkie/notes-g2/internal"
	usersDomain "github.com/Wookkie/notes-g2/internal/domain/users"
	"github.com/gin-gonic/gin"
)

type Repository interface { // описываем те методы, которые мы ожидаем от хранилища (от реализации БД)
	SaveUser(user usersDomain.User) error //string- это ID
	GetUser(login string) (usersDomain.User, error)
}

type NotesAPI struct {
	cfg       *internal.Config
	httpServe *http.Server
	repo      Repository
}

func New(cfg *internal.Config, repo Repository) *NotesAPI {
	httpServe := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	NotesAPI := NotesAPI{
		httpServe: &httpServe,
		cfg:       cfg,
		repo:      repo,
	}
	NotesAPI.configRoutes()

	return &NotesAPI
}

func (nApi *NotesAPI) Run() error {
	return nApi.httpServe.ListenAndServe()
}

func (nApi *NotesAPI) Stop(ctx context.Context) error {
	return nApi.httpServe.Shutdown(ctx)
}

func (nApi *NotesAPI) configRoutes() {
	router := gin.Default()
	router.GET("/")
	users := router.Group("/users")
	{
		users.GET("/porfile")
		users.POST("/register", nApi.register)
		users.POST("/login", nApi.login)
	}

	notes := router.Group("/notes")
	{
		notes.GET("/")
		notes.GET("/:id")
		notes.POST("/")
		notes.PUT("/:id")
		notes.DELETE("/:id")
	}

	nApi.httpServe.Handler = router
}
