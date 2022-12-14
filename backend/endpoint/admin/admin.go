package endpoint

import (
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/util"
	"github.com/gin-gonic/gin"
)

func AdminChangeInfo(c *gin.Context) {
	util.CallSuccessOK(c, "success", nil)
}
