package transport

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/pintro-golang/data/model"
	"github.com/uchupx/pintro-golang/helper/crypt"
)

type Middleware struct {
	CryptService crypt.CryptService
	// Redis        *redis.Client
}

func (m Middleware) Authorization(c *gin.Context) {
	// Authorization
	var user model.User

	header := c.Request.Header

	if len(header["Authorization"]) < 1 {
		c.JSON(http.StatusUnauthorized, "unauthenticated")
		c.Abort()
		return
	}

	authStr := header["Authorization"][0]
	tokenArry := strings.Split(authStr, " ")

	if len(tokenArry) < 2 {
		c.JSON(http.StatusUnauthorized, "unauthenticated")
		c.Abort()
		return
	}

	token := tokenArry[1]

	userStr, err := m.CryptService.VerifyJWTToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthenticated")
		c.Abort()
		return
	}

	err = json.Unmarshal([]byte(userStr.(string)), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	// key := fmt.Sprintf(constant.RedisTokenKey, user.ID)

	// // redisToken, err := m.Redis.Get(c, key).Result()
	// // if err == redis.Nil {
	// // 	err = errors.ErrAuthTokenExpired

	// // 	handler.WriteErrorResponse(c, err)
	// // 	return
	// // }

	// if redisToken != token {
	// 	err = errors.ErrUnauthorized
	// 	handler.WriteErrorResponse(c, err)
	// 	return
	// }

	c.Set("user-key", user)

	c.Next()
}
