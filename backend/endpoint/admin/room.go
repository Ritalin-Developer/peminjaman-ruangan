package endpoint

import (
	"fmt"
	"strconv"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/external"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ListRoom(c *gin.Context) {
	db, err := external.GetPostgresClient()
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	rooms := []*model.Room{}
	err = db.
		Model(&rooms).
		Find(&rooms).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	util.CallSuccessOK(c, "success", rooms)
}

type roomRequest struct {
	RoomNumber  string `json:"room_number"`
	Remark      string `json:"remark"`
	IsAvailable bool   `json:"is_available"`
}

func RegisterRoom(c *gin.Context) {
	request := &roomRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallUserError(c, "invalid request", err)
		return
	}

	if request.RoomNumber == "" || request.Remark == "" {
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
	room := &model.Room{
		RoomNumber:  request.RoomNumber,
		Remark:      request.Remark,
		IsAvailable: request.IsAvailable || false,
	}
	err = db.
		Create(&room).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}

	util.CallSuccessOK(c, "success", nil)
}

func UpdateRoomInformation(c *gin.Context) {
	_roomID := c.Query("room_id")
	roomID, err := strconv.Atoi(_roomID)
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallUserError(c, "invalid request", err)
		return
	}
	request := &roomRequest{}
	err = c.BindJSON(&request)
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
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	room := &model.Room{}
	err = db.
		Model(&room).
		Where("id = ?", roomID).
		Updates(map[string]interface{}{
			"room_number":  request.RoomNumber,
			"remark":       request.Remark,
			"is_available": request.IsAvailable,
		}).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}

	util.CallSuccessOK(c, "success", nil)
}
