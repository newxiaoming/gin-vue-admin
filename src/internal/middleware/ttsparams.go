package middleware

import (
	"fmt"
	"gin-vue-admin/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type TTSParams struct {
	Business struct {
		Aue   string `json:"aue" binding:"required"`
		Vcn   string `json:"vcn" binding:"required"`
		Pitch int    `json:"pitch" binding:"required"`
		Speed int    `json:"speed" binding:"required"`
	} `json:"business" binding:"required"`
	Data struct {
		Status int    `json:"status" binding:"required"`
		Text   string `json:"text" binding:"required"`
	} `json:"data" binding:"required"`
}

// CheckTTSParams 检查语音合成（流式版）WebAPI提交的字段
func CheckTTSParams() gin.HandlerFunc {
	var (
		cp TTSParams
		pp publicParams
	)

	return func(c *gin.Context) {

		if err := c.ShouldBindQuery(&pp); err != nil {
			response.FailResult(500, err.Error(), c)
			c.Abort()
		}

		if err := c.ShouldBindBodyWith(&cp, binding.JSON); err == nil {
			validate := validator.New()
			log.Printf("%+v", cp)
			if err := validate.Struct(&cp); err != nil {
				response.FailResult(503, err.Error(), c)
				c.Abort()
				fmt.Println("503")
			}
		} else {
			response.FailResult(502, err.Error(), c)
			c.Abort()
			return
		}

		c.Next()
	}
}
