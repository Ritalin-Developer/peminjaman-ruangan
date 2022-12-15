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

func SubmissionList(c *gin.Context) {
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
	submissions := []*model.Submission{}
	err = db.
		Model(&submissions).
		Limit(limit).
		Find(&submissions).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	util.CallSuccessOK(c, "success", submissions)
}

type submissionRequest struct {
	RoomID       uint64 `json:"room_id"`
	RoomNumber   string `json:"room_number"`
	Remark       string `json:"remark"`
	StartUseDate string `json:"start_use_date"`
	EndUseDate   string `json:"end_use_date"`
}

func SubmissionCreate(c *gin.Context) {
	request := &submissionRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallUserError(c, "invalid request", err)
		return
	}

	if request.RoomID == 0 || request.RoomNumber == "" || request.Remark == "" || request.StartUseDate == "" || request.EndUseDate == "" {
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

	room := &model.Room{}
	err = db.
		Model(&room).
		Where("id = ?", request.RoomID).
		First(&room).
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
	if !room.IsAvailable {
		err = fmt.Errorf("room is not available")
		util.CallUserError(c, "make sure to check list room available", err)
		return
	}

	submission := &model.Submission{
		RoomNumber:   request.RoomNumber,
		Remark:       request.Remark,
		StartUseDate: request.StartUseDate,
		EndUseDate:   request.EndUseDate,
		IsApproved:   false,
		RoomID:       request.RoomID,
	}
	err = db.
		Create(&submission).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}

	util.CallSuccessOK(c, "success", nil)
}

func SubmissionUpdate(c *gin.Context) {
	submissionID := c.Query("submission_id")
	request := &submissionRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallUserError(c, "invalid request", err)
		return
	}
	if submissionID == "" || request.RoomID == 0 || request.RoomNumber == "" || request.Remark == "" || request.StartUseDate == "" || request.EndUseDate == "" {
		err = fmt.Errorf("submission_id, room_id, room_number, remark, start_use_date, end_use_date field cannot be empty")
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
	submission := &model.Submission{}
	err = db.
		Model(&submission).
		Where("id = ?", submissionID).
		Updates(map[string]interface{}{
			"room_id":        request.RoomID,
			"room_number":    request.RoomNumber,
			"remark":         request.Remark,
			"start_use_date": request.StartUseDate,
			"end_use_date":   request.EndUseDate,
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

func SubmissionDelete(c *gin.Context) {
	submissionID := c.Query("submission_id")
	if submissionID == "" {
		err := fmt.Errorf("submission_id field cannot be empty")
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
	submission := &model.Submission{}
	err = db.
		Where("id = ?", submissionID).
		Delete(&submission).
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
