package middleware

import (
	"gin-vue-admin/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CorrentionParams struct {
	Content []struct {
		Text string `json:"text" binding:"required"`
		Line string `json:"line" binding:"required"`
	} `josn:"content"`
	Total string `json:"total" binding:"required"`
}

type publicCorrentionParams struct {
	Charset      string `form:"charset" json:"charset" binding:"required"`
	ProviderName string `form:"provider_name" json:"provider_name" `
	Signature    string `form:"signature" json:"signature"  binding:"required"`
	Date         string `form:"date" format_date:"200601020000" json:"date" binding:"required"`
}

func CheckCorrentionParams() gin.HandlerFunc {
	var (
		cp CorrentionParams
		// pp publicCorrentionParams
	)

	return func(c *gin.Context) {

		// if err := c.ShouldBindQuery(&pp); err != nil {
		// 	response.FailResult(500, err.Error(), c)
		// 	c.Abort()
		// }

		if err := c.ShouldBindBodyWith(&cp, binding.JSON); err != nil {
			response.FailResult(502, err.Error(), c)
			c.Abort()
		}
		// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(&cp))
		c.Next()
	}
}
