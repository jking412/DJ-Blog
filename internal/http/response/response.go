package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (

	// OK is a successful response
	OK = NewStatus(http.StatusOK, "操作成功")

	// BadRequest is a bad request
	BadRequest         = NewStatus(http.StatusBadRequest, "请求错误")
	UnprocessableData  = NewStatus(http.StatusUnprocessableEntity, "数据格式错误")
	UnsatisfiedRequest = NewStatus(http.StatusUnprocessableEntity, "请求参数不满足要求")
	Unauthorized       = NewStatus(http.StatusUnauthorized, "未授权")

	// InternalServerError is a internal server error
	InternalServerError = NewStatus(http.StatusInternalServerError, "服务器内部错误")
)

type Status struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewStatus(code int, msg string) *Status {
	return &Status{
		Code: code,
		Msg:  msg,
	}
}

func EndWithJSON(c *gin.Context, status *Status, data interface{}) {
	status.Data = data
	c.JSON(status.Code, status)
}

func EndWithOK(c *gin.Context, data interface{}) {
	EndWithJSON(c, OK, data)
}

func EndWithBadRequest(c *gin.Context, data interface{}) {
	EndWithJSON(c, BadRequest, data)
}

func EndWithUnauthorized(c *gin.Context, data interface{}) {
	EndWithJSON(c, Unauthorized, data)
}

func EndWithUnprocessableData(c *gin.Context, data interface{}) {
	EndWithJSON(c, UnprocessableData, data)
}

func EndWithUnsatisfiedRequest(c *gin.Context, data interface{}) {
	EndWithJSON(c, UnsatisfiedRequest, data)
}

func EndWithInternalServerError(c *gin.Context, data interface{}) {
	EndWithJSON(c, InternalServerError, data)
}
