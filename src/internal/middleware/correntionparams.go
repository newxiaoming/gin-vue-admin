package middleware

import (
	"gin-vue-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type CorrentionParams struct {
	Content struct {
		Text string `json:"text" binding:"required"`
		Line int    `json:"line" binding:"required,min=1"`
	} `josn:"content"`
	Total int `json:"total" binding:"required,min=1"`
}

type publicCorrentionParams struct {
	Charset      string `form:"charset" json:"charset" binding:"required,min=1"`
	ProviderName string `form:"provider_name" json:"provider_name" `
	Signature    string `form:"signature" json:"signature"  binding:"required"`
	Date         string `form:"date" format_date:"200601020000" json:"date" binding:"required"`
}

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

		if err := c.ShouldBind(&cp); err != nil {
			response.FailResult(500, err.Error(), c)
			c.Abort()
		}
		c.Next()
	}
}
