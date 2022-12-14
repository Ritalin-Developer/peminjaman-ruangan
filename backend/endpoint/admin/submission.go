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
	util.CallSuccessOK(c, "success", nil)
}

func SubmissionReject(c *gin.Context) {
	util.CallSuccessOK(c, "success", nil)
}
