package middleware

import (
	"fmt"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/gin-gonic/gin"
)

func ValidateRoleAccess(c *gin.Context) {
	tokenData, exist := c.Get("user")
	if !exist {
		util.MiddlewareCallUserUnauthorized(c, "invalid token data", fmt.Errorf("missing token data"))
	}
	data := tokenData.(*model.TokenUserData)
	if data.RoleName != "admin" {
		util.MiddlewareCallUserUnauthorized(c, "invalid token data", fmt.Errorf("you are not an admin"))
	}
	c.Next()
}
