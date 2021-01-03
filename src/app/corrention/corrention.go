package corrention

import (
	"errors"
	"fmt"
	"gin-vue-admin/pkg/response"
	"net/http"
	"reflect"
	"time"

	xfyunauthorization "gin-vue-admin/pkg/xfyun/authorization"
	xfyuntextcorrention "gin-vue-admin/pkg/xfyun/textcorrention"

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

// 讯飞文本纠错具体实现方法
func xfyun(c *gin.Context) int {
	fmt.Println("running function xfyun!", c.Query("charset"))

	host := "api.xf-yun.com"
	date := time.Now().UTC().Format(http.TimeFormat)
	c.Set("host", host)
	c.Set("date", date)

	authorization := xfyunauthorization.Authorization(c)
	c.Set("authorization", authorization)
	rst := xfyuntextcorrention.PostData(c)
	fmt.Println("authorization=", authorization)
	fmt.Println(rst)
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
