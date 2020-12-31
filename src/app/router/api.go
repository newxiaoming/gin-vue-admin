package router

import (
	"gin-vue-admin/app/corrention"

	"github.com/gin-gonic/gin"
)

func noCheckSignRouterV1(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.POST("/nlp/corrention", corrention.NlpTextCorrention)
}
