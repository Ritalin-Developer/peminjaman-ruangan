package endpoint

import (
	"strconv"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/external"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UserRoomList(c *gin.Context) {
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
