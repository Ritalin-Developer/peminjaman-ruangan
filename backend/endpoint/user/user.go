package endpoint

import (
	"fmt"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/external"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
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
	if request.Username == "" || request.Password == "" || request.RealName == "" {
		err = fmt.Errorf("username, password, and real_name field cannot be empty")
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

func UserChangeInfo(c *gin.Context) {
	util.CallSuccessOK(c, "success", nil)
}
