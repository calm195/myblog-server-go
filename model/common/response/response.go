package response

import (
	"myblog-server-go/model/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response
//
//	@Description: 通用返回结构体
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

// Result
//
//	@Description: 将返回信息写入gin.Context，类型为application/json
//	@param code http状态码 by http
//	@param data 返回数据
//	@param msg 返回信息
//	@param c *gin.Context
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// Ok
//
//	@Description: 仅返回成功状态，msg为"success"，data为空map
//	@param c *gin.Context
func Ok(c *gin.Context) {
	Result(SUCCESS, common.JSONMap{}, "success", c)
}

// OkWithMessage
//
//	@Description: 返回成功状态，msg自定义，data为空map
//	@param message 自定义返回信息
//	@param c *gin.Context
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, common.JSONMap{}, message, c)
}

// OkWithData
//
//	@Description: 返回成功状态，msg为"success"，data自定义
//	@param data 自定义返回数据
//	@param c *gin.Context
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

// OkWithDetailed
//
//	@Description: 返回成功状态，msg自定义，data自定义
//	@param data 自定义返回数据
//	@param message 自定义返回信息
//	@param c *gin.Context
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// Fail
//
//	@Description: 返回失败状态，msg为"failure"，data为空map
//	@param c *gin.Context
func Fail(c *gin.Context) {
	Result(ERROR, common.JSONMap{}, "failure", c)
}

// FailWithMessage
//
//	@Description: 返回失败状态，msg自定义，data为空map
//	@param message 自定义返回信息
//	@param c *gin.Context
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, common.JSONMap{}, message, c)
}

// FailWithDetailed
//
//	@Description: 返回失败状态，msg自定义，data自定义
//	@param data 自定义返回数据
//	@param message 自定义返回信息
//	@param c *gin.Context
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
