package router

import (
	"gin-vue-admin/app/corrention"
	"gin-vue-admin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func noCheckSignRouterV1(r *gin.RouterGroup) {
	v1 := r.Group("/api/v2")

	v1.POST("/text/corrention", middleware.CheckCorrentionParams(), corrention.NlpTextCorrention)
}
