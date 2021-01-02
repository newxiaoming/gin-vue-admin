package corrention

import (
	"errors"
	"fmt"
	"gin-vue-admin/pkg/response"
	"reflect"

	xfyunauthorization "gin-vue-admin/pkg/xfyun/authorization"

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

func xfyun(c *gin.Context) int {
	fmt.Println("running function xfyun!", c.Query("charset"))

	authorization := xfyunauthorization.Authorization()
	fmt.Println("authorization=", authorization)
	return 100
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	return
}
func Handle(c *gin.Context) {
	funcs := map[string]interface{}{
		"xfyun": xfyun,
	}

	if result, err := Call(funcs, "xfyun", c); err == nil {
		fmt.Println("result:", result)
		for _, r := range result {
			fmt.Printf(" return: type=%v, value=[%d]\n", r.Type(), r.Int())
		}
	} else {
		fmt.Println("error:", err)
	}
}

// NlpTextCorrention 入口
func NlpTextCorrention(c *gin.Context) {
	name := c.DefaultQuery("provider_name", "xfyun")
	newProvider := &provider{
		name:    name,
		text:    "时代财经智能化",
		service: "textcorrention",
	}

	Handle(c)

	var newTextCorrection textCorrection
	newTextCorrection = newProvider
	response.SuccessResult(gin.H{
		"text": newProvider.name,
	}, c)
	fmt.Println(newTextCorrection.corrention())
}
