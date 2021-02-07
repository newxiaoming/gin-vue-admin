package router

import (
	"gin-vue-admin/app/corrention"
	"gin-vue-admin/app/tts"
	"gin-vue-admin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func noCheckSignRouterV2(r *gin.RouterGroup) {
	v2 := r.Group("/api/v2")

	v2.POST("/text/corrention", middleware.CheckCorrentionParams(), corrention.NlpTextCorrention)
	v2.POST("/tts/onlinetts", middleware.CheckTTSParams(), tts.NlpTTS)
}
