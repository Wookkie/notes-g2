package server

import (
	"net/http"

	usersDomain "github.com/Wookkie/notes-g2/internal/domain/users"
	usersService "github.com/Wookkie/notes-g2/internal/services/user"
	"github.com/gin-gonic/gin"
)

func (s *NotesAPI) login(ctx *gin.Context) {
	var uReq usersDomain.UserRequest

	if err := ctx.ShouldBindBodyWithJSON(&uReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	UserService := usersService.New(s.repo)

	userID, err := UserService.LoginUser(uReq)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "user logined: %s", userID)
}

func (s *NotesAPI) register(ctx *gin.Context) {
	var uReq usersDomain.User

	if err := ctx.ShouldBindBodyWithJSON(&uReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	UserService := usersService.New(s.repo)

	userID, err := UserService.RegisterUser(uReq)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "user registered: %s", userID)
}
