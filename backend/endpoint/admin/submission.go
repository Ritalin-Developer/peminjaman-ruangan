package endpoint

import (
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/gin-gonic/gin"
)

func SubmissionList(c *gin.Context) {
	util.CallSuccessOK(c, "success", nil)
}
func SubmissionApprove(c *gin.Context) {
	util.CallSuccessOK(c, "success", nil)
}

func SubmissionReject(c *gin.Context) {
	util.CallSuccessOK(c, "success", nil)
}
