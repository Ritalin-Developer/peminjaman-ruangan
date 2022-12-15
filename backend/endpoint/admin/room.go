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
	"gorm.io/gorm"
)

func ListRoom(c *gin.Context) {
	// TODO: Implement search by query/filter
	_limit := c.Query("limit")
	limit, err := strconv.Atoi(_limit)
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallUserError(c, "must provide limit query", err)
		return
	}

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
		Limit(limit).
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
	Capacity    uint   `json:"capacity"`
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
	if request.RoomNumber == "" || request.Remark == "" || request.Capacity == 0 {
		err = fmt.Errorf("room_number, remark, and capacity field cannot be empty or 0")
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
		Capacity:    request.Capacity,
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
			"capacity":     request.Capacity,
		}).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			util.CallUserError(c, "room doesn't exist", err)
			return
		}
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}

	util.CallSuccessOK(c, "success", nil)
}

func DeleteRoom(c *gin.Context) {
	_roomID := c.Query("room_id")
	roomID, err := strconv.Atoi(_roomID)
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
		Delete(&room).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			util.CallUserError(c, "room doesn't exist", err)
			return
		}
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}

	util.CallSuccessOK(c, "success", nil)
}
