package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Initrouter 初始化路由
func Initrouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(Cors())
	// 抓取网页正文
	// r.POST("api/nlp/corrention", web2text.Web2text)
	g := r.Group("")
	noCheckSignRouterV1(g)
	return r
}

// Cors 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里可以用*，也可以用指定的域名
		c.Header("Access-Control-Allow-Origin", "*")
		// 允许头部参数
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		// 允许的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		method := c.Request.Method
		//放行OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		// 处理请求
		c.Next()
	}
}
