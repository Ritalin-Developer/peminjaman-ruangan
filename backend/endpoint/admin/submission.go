package endpoint

import (
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/external"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SubmissionList(c *gin.Context) {
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
		Find(&submission).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	util.CallSuccessOK(c, "success", submission)
}

func SubmissionApprove(c *gin.Context) {
	submissionID := c.Query("submission_id")

	db, err := external.GetPostgresClient()
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	tx := db.Begin()
	defer tx.Rollback()

	submission := &model.Submission{}
	err = tx.
		Model(&submission).
		Where("id = ?", submissionID).
		Update("is_approved", true).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}

	room := &model.Room{}
	err = tx.
		Model(&room).
		Where("id = ?", submission.RoomID).
		Where("is_available = true").
		Update("is_available", false).
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
	tx.Commit()

	util.CallSuccessOK(c, "success", nil)
}

func SubmissionReject(c *gin.Context) {
	submissionID := c.Query("submission_id")

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
		Update("is_approved", false).
		Error
	if err != nil {
		log.Error(err)
		util.CallServerError(c, "something wrong, please try again later", err)
		return
	}
	util.CallSuccessOK(c, "success", nil)
}
