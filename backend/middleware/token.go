package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

func MiddlewareValidateToken(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.MiddlewareCallUserUnauthorized(c, "fail to validate token", err)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if err != nil {
			log.Error(err)
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		token.Valid = true
		return token, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var exp int64
		switch e := claims["expire_at"].(type) {
		case string:
			conv, err := strconv.Atoi(e)
			if err != nil {
				log.Error(err)
				break
			}
			exp = int64(conv)
		case float64:
			exp = int64(e)
		}
		if exp < time.Now().Unix() {
			err = fmt.Errorf("expired token")
			util.MiddlewareCallUserUnauthorized(c, "something wrong with the token", err)
			c.Abort()
		}
		data := &model.TokenUserData{
			Username: claims["username"].(string),
			Issuer:   claims["issuer"].(string),
			RoleID:   claims["role_id"].(float64),
			RoleName: claims["role_name"].(string),
		}
		c.Set("data", data)
		c.Next()
	} else {
		err = fmt.Errorf("invalid token claims")
		util.MiddlewareCallUserUnauthorized(c, "invalid token", err)
	}
	c.Next()
}
