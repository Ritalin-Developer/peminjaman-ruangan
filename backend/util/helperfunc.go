package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Msg     *string     `json:"msg"`
	Data    interface{} `json:"data"`
	Link    *string     `json:"link"`
}

// Contains function is to check item whether is exist or not in a list and will return bool
func Contains(d string, dl []string) bool {
	for _, v := range dl {
		if v == d {
			return true
		}
	}
	return false
}

// CallErrorNotFound is for return API response not found
func CallErrorNotFound(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"err":     err.Error(),
		"msg":     msg,
		"data":    map[string]interface{}{},
	})
	c.Next()
}

// CallServerError is for return API response server error
func CallServerError(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"err":     err.Error(),
		"msg":     msg,
		"data":    map[string]interface{}{},
	})
	c.Next()
}

// MiddlewareCallServerError is for return API response server error
func MiddlewareCallServerError(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"err":     err.Error(),
		"msg":     msg,
		"data":    map[string]interface{}{},
	})
	c.Abort()
}

// CallUserError is for return API response server error
func CallUserError(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"err":     err.Error(),
		"msg":     msg,
		"data":    map[string]interface{}{},
	})
	c.Next()
}

// CallSuccessOK is for return API response with status code 200, you need to specify msg, and data as function parameter
func CallSuccessOK(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   nil,
		"msg":     msg,
		"data":    data,
	})
	c.Next()
}

// CallUserUnauthorized is for return API response user is not authorized
func CallUserUnauthorized(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"err":     err.Error(),
		"msg":     msg,
		"data":    map[string]interface{}{},
	})
	c.Next()
}

// MiddlewareCallUserUnauthorized is for return API response server error
func MiddlewareCallUserUnauthorized(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"err":     err.Error(),
		"msg":     msg,
		"data":    map[string]interface{}{},
	})
	c.Abort()
}

// CallSuccessOkWithLink is for return API response with status code 200, you need to specify msg, and data as function parameter
func CallSuccessOkWithLink(c *gin.Context, msg string, data interface{}, link *string) {
	hashMap := Response{
		Success: true,
		Msg:     &msg,
		Data:    data,
		Link:    link,
	}
	c.JSON(http.StatusOK, hashMap)
	c.Next()
}

// CallUserErrorWithLink is for return API response server error
func CallUserErrorWithLink(c *gin.Context, msg string, err error, data interface{}, link *string) {
	hashMap := Response{
		Success: false,
		Error:   err.Error(),
		Msg:     &msg,
		Data:    data,
		Link:    link,
	}
	c.JSON(http.StatusBadRequest, hashMap)
	c.Next()
}

// CallSuccessOKWithTemplate is for return API response with status code 200, you need to specify msg, and data as function parameter
func CallSuccessOKWithTemplate(c *gin.Context, msg string, data interface{}, template interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"error":    nil,
		"msg":      msg,
		"data":     data,
		"template": template,
	})
	c.Next()
}
