package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ismail118/booklending/util"
	"net/http"
	"time"
)

type userResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	User        userResponse
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) loginUser(c *gin.Context) {
	var req loginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.querier.GetUser(context.Background(), req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	token, err := server.paseto.CreateToken(user.Email, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := loginResponse{
		AccessToken: token,
		User: userResponse{
			Email: user.Email,
			Name:  user.Name,
		},
	}

	c.JSON(http.StatusOK, resp)
}
