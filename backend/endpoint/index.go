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

	if request.Username == "" || request.Password == "" {
		err = fmt.Errorf("username and password field cannot be empty")
		log.Error(err)
		util.CallUserError(c, "invalid request", err)
		return
	}

	db, err := external.GetPostgresClient()
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "something wrong, please try again later", err)
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
		util.CallServerError(c, "something wrong, please try again later", err)
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
