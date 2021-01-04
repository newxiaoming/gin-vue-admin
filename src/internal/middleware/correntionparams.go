package middleware

import (
	"fmt"
	"gin-vue-admin/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type CorrentionParams struct {
	Content []struct {
		Text string `json:"text" binding:"required"`
		Line string `json:"line" binding:"required"`
	} `json:"content" binding:"required,dive"`
	Total string `json:"total" binding:"required"`
}

type publicCorrentionParams struct {
	Charset      string `form:"charset" json:"charset" binding:"required"`
	ProviderName string `form:"provider_name" json:"provider_name" `
	Signature    string `form:"signature" json:"signature"  binding:"required"`
	Date         string `form:"date" format_date:"200601020000" json:"date" binding:"required"`
}

// CheckCorrentionParams 检查提交的字段
func CheckCorrentionParams() gin.HandlerFunc {
	var (
		cp CorrentionParams
		pp publicCorrentionParams
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
