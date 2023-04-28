package transport

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/model"
	"github.com/uchupx/pintro-golang/helper/crypt"
	"github.com/uchupx/pintro-golang/transport/payload"
)

type UserHandler struct {
	UserRepository data.UserRepoitory
	CryptService   crypt.CryptService
}

func (h UserHandler) Login(c *gin.Context) {
	var request payload.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.UserRepository.FindByUsername(c, request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, "username / password invalid")
		return
	}

	isSame, err := h.CryptService.Verify(request.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "username / password invalid")
		return
	}

	if !isSame {
		c.JSON(http.StatusBadRequest, "username / password invalid")
		return
	}

	userByte, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	token, err := h.CryptService.CreateJWTToken(time.Hour*time.Duration(1), string(userByte))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, payload.TokenReponse{
		Token: *token,
	})
}

func (h UserHandler) Post(c *gin.Context) {
	var request payload.UserPostRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.CryptService.CreateSignPSS(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user := model.User{
		Name:     request.Name,
		Password: result,
		Username: request.Username,
	}

	_, err = h.UserRepository.Insert(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "success")
	return
}
