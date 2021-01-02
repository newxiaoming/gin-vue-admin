package middleware

import (
	"gin-vue-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type xfyunHeaderParams struct {
	Header struct {
		AppID  string `json:"app_id" binding:"required"`
		Status int    `json:"status" binding:"required,max=3,min=1"`
	} `josn:"header"`
	Parameter struct {
		S9a87e3ec struct {
			Result struct {
				EnCoding string `json:"encoding" binding:"required"`
				Compress string `json:"comporess" bind:"required"`
				Format   string `json:"format" binding:"required"`
			} `json:"result" binding:"required"`
		} `json:"s9a87e3ec" binding:"required"`
	} `json:"parameter" binding:"required"`
	Payload struct {
		Input struct {
			EnCoding string `json:"encoding" binding:"required"`
			Compress string `json:"comporess" bind:"required"`
			Format   string `json:"format" binding:"required"`
			Status   string `json:"status" binding:"required"`
			Text     string `json:"text" binding:"required"`
		}
	}
}

func CheckXfyunParams() gin.HandlerFunc {
	var xfp xfyunHeaderParams

	return func(c *gin.Context) {
		if err := c.ShouldBind(&xfp); err != nil {
			response.FailResult(500, err.Error(), c)
			c.Abort()
		}
		c.Next()
	}
}
