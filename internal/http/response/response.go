package response

import "github.com/gin-gonic/gin"

var (

	// OK is a successful response
	// 20000 - 29999 表示请求成功
	OK = NewStatus(20000, "操作成功")

	// BadRequest is a bad request
	// 40000 - 49999 为用户端错误
	BadRequest         = NewStatus(40000, "请求错误")
	UnprocessableJSON  = NewStatus(40001, "JSON 格式错误")
	UnsatisfiedRequest = NewStatus(40002, "请求参数不满足要求")

	// InternalServerError is a internal server error
	// 50000 - 59999 为服务器端错误
	InternalServerError = NewStatus(50000, "服务器内部错误")
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

func EndWithUnprocessableJSON(c *gin.Context, data interface{}) {
	EndWithJSON(c, UnprocessableJSON, data)
}

func EndWithUnsatisfiedRequest(c *gin.Context, data interface{}) {
	EndWithJSON(c, UnsatisfiedRequest, data)
}

func EndWithInternalServerError(c *gin.Context, data interface{}) {
	EndWithJSON(c, InternalServerError, data)
}
