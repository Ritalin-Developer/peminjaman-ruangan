package middleware

import (
	"fmt"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func ValidateRoleAccess(c *gin.Context) {
	tokenData, exist := c.Get("data")
	if !exist {
		util.MiddlewareCallUserUnauthorized(c, "invalid token data", fmt.Errorf("missing token data"))
		return
	}
	data := &model.TokenUserData{}
	mapstructure.Decode(tokenData, &data)
	if data.RoleName != "admin" {
		util.MiddlewareCallUserUnauthorized(c, "invalid token data", fmt.Errorf("you are not an admin"))
		return
	}
	c.Next()
}
