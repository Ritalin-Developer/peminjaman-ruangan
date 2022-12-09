package endpoint

import (
	"fmt"
	"time"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/external"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RealName string `json:"real_name"`
}

func Register(c *gin.Context) {
	request := &registerRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallUserError(c, "invalid request", err)
		return
	}

	db, err := external.GetPostgresClient()
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "server busy, please try again later", err)
		return
	}

	user := &model.User{}
	err = db.
		Model(&user).
		Where("username = ?", request.Username).
		Find(&user).
		Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Error(err)
			sentry.CaptureException(err)
			util.CallServerError(c, "fail to generate user", err)
			return
		}
	}
	if user.Username != "" {
		err = fmt.Errorf("user already registered")
		util.CallUserError(c, "user with specified username already registered", err)
		return
	}

	hash := util.HashAndSalt([]byte(request.Password))
	err = db.
		Create(&model.User{
			Username: request.Username,
			Password: hash,
			RealName: request.RealName,
			RoleID:   2,
		}).
		Error
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "fail to generate user", err)
		return
	}

	util.CallSuccessOK(c, "success", nil)
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func Login(c *gin.Context) {
	request := &loginRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallUserError(c, "invalid request", err)
		return
	}

	db, err := external.GetPostgresClient()
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "server busy, please try again later", err)
		return
	}

	tx := db.Begin()
	defer tx.Rollback()

	user := &model.User{}
	err = tx.
		Model(&user).
		Where("username = ?", request.Username).
		Find(&user).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("user is not registered")
			util.CallUserError(c, "user is not registered", err)
			return
		}
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "fail to login", err)
		return
	}
	isPasswordMatch := util.ComparePasswords(user.Password, request.Password)
	if !isPasswordMatch {
		err = fmt.Errorf("user is not registered")
		util.CallUserError(c, "user is not registered", err)
		return
	}

	role := &model.Role{}
	err = tx.
		Model(&role).
		Where("id = ?", user.RoleID).
		Find(&role).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("user is not registered")
			util.CallUserError(c, "user is not registered", err)
			return
		}
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "fail to login", err)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "server busy, please try again later", err)
		return
	}

	// Generate token
	tokenData := jwt.MapClaims{
		"username":  user.Username,
		"issuer":    "peminjaman-ruangan",
		"issued_at": time.Now().Unix(),
		"role_id":   user.RoleID,
		"role_name": role.RoleName,
	}
	token, err := util.GenerateToken(tokenData)
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "fail to generate token", err)
		return
	}
	response := &loginResponse{
		Username: request.Username,
		Token:    token,
	}
	util.CallSuccessOK(c, "success", response)
}

func UserValidateToken(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallUserError(c, "fail to validate token", err)
		return
	}
	tokenValid, tokenData, err := util.ValidateToken(token)
	if !tokenValid || tokenData == nil {
		err = fmt.Errorf("invalid token or mismatch token data")
		util.CallUserError(c, "invalid token", err)
		return
	}
	if err != nil {
		util.CallUserError(c, "invalid token", err)
		return
	}
	if _, exist := tokenData["username"]; !exist {
		err = fmt.Errorf("missing username in token")
		util.CallUserError(c, "token missing username", err)
		return
	}
	if _, exist := tokenData["role_id"]; !exist {
		err = fmt.Errorf("missing role_id in token")
		util.CallUserError(c, "token missing role_id", err)
		return
	}
	util.CallSuccessOK(c, "success", map[string]interface{}{
		"username": tokenData["username"],
		"role_id":  tokenData["role_id"],
	})
}
