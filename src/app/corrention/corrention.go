package corrention

import (
	"fmt"
	"gin-vue-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

// textCorrection 接口
type textCorrection interface {
	corrention() string
}

// provider 服务提供者
type provider struct {
	name    string
	text    string
	service string
}

// corrention 文本纠错函数
func (p *provider) corrention() string {
	return p.name
}

// NlpTextCorrention 入口
func NlpTextCorrention(c *gin.Context) {
	name := c.DefaultQuery("name", "xfyun")
	newProvider := &provider{
		name:    name,
		text:    "时代财经智能化",
		service: "textcorrention",
	}

	var newTextCorrection textCorrection
	newTextCorrection = newProvider
	response.SuccessResult(gin.H{
		"text": newProvider.name,
	}, c)
	fmt.Println(newTextCorrection.corrention())
}
