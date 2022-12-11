package endpoint

import (
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/external"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SubmissionList(c *gin.Context) {
	isApproved := c.Query("is_approved")

	db, err := external.GetPostgresClient()
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	submission := []*model.Submission{}
	err = db.
		Model(&submission).
		Where("is_approved = ?", isApproved).
		Find(&submission).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	util.CallSuccessOK(c, "success", nil)
}

type submissionCreateRequest struct {
	RoomID       uint64 `json:"room_id"`
	RoomNumber   string `json:"room_number"`
	Remark       string `json:"remark"`
	StartUseDate string `json:"start_use_date"`
	EndUseDate   string `json:"end_use_date"`
}

func SubmissionCreate(c *gin.Context) {
	request := &submissionCreateRequest{}
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
		util.CallServerError(c, "something wrong, please try again later", err)
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

func SubmissionApprove(c *gin.Context) {
	util.CallSuccessOK(c, "success", nil)
}

func SubmissionReject(c *gin.Context) {
	util.CallSuccessOK(c, "success", nil)
}
