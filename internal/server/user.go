package server

import (
	"net/http"

	usersDomain "github.com/Wookkie/notes-g2/internal/domain/users"
	"github.com/gin-gonic/gin"
)

func (s *Server) login(ctx *gin.Context) {
	var uReq usersDomain.UserReqest

	if err := ctx.ShouldBindBodyWithJSON(&uReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	//TODO: вызов бизнес логики авторизации

	ctx.JSON(http.StatusOK, nil)
}
