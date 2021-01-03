package router

import (
	"gin-vue-admin/app/corrention"
	"gin-vue-admin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func noCheckSignRouterV2(r *gin.RouterGroup) {
	v2 := r.Group("/api/v2")

	v2.POST("/text/corrention", middleware.CheckCorrentionParams(), corrention.NlpTextCorrention)
}
