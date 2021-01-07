package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result 结构
type Result struct {
	Status  string        `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

func response(status string, code int, msg string, data []interface{}, c *gin.Context) {
	fmt.Println(data)
	r := Result{status, code, msg, data}
	d, _ := json.Marshal(r)
	fmt.Println(string(d))
	c.JSON(http.StatusOK, r)
}

func successResponse(data []interface{}, c *gin.Context) {
	response("success", 0, "请求成功", data, c)
}

func failResponse(code int, msg string, c *gin.Context) {
	response("fail", code, msg, []interface{}{}, c)
}

// SuccessResultWithEmptyData 空数据
func SuccessResultWithEmptyData(c *gin.Context) {
	successResponse(nil, c)
}

// SuccessResult 成功响应的数据
func SuccessResult(data []interface{}, c *gin.Context) {
	successResponse(data, c)
}

// FailResultWithDefaultMsg 失败响应返回空数据
func FailResultWithDefaultMsg(code int, c *gin.Context) {
	failResponse(code, "请求失败", c)
}

// FailResult 失败响应
func FailResult(code int, msg string, c *gin.Context) {
	failResponse(code, msg, c)
}
