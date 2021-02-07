package tts

import (
	"errors"
	"fmt"
	"gin-vue-admin/internal/middleware"
	"gin-vue-admin/pkg/response"
	xfyunauthorization "gin-vue-admin/pkg/xfyun/authorization"
	xfyuntts "gin-vue-admin/pkg/xfyun/xfyuntts"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// tts 接口
type tts interface {
	tts() string
}

// provider 服务提供者
type provider struct {
	name string
	text string
}

// tts 文本纠错函数
func (p *provider) tts() string {
	return p.name
}

var (
	cp middleware.TTSParams
)

// 讯飞TTS具体实现方法
func xfyun(c *gin.Context, text string) response.TextCorrentionResponseItem {
	fmt.Println("running function xfyun!", c.Query("charset"))

	host := "api.xf-yun.com"
	date := time.Now().UTC().Format(http.TimeFormat)
	c.Set("host", host)
	c.Set("date", date)

	authorization := xfyunauthorization.Authorization(c)
	c.Set("authorization", authorization)
	rst := xfyuntts.TTSRequest(c, text)
	fmt.Println("authorization=", authorization)
	return rst
}

// Call 调用
func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	return
}

// Handle 函数调用入口
func Handle(c *gin.Context, providerName string, text string) interface{} {
	var rst interface{}
	funcs := map[string]interface{}{
		"xfyun": xfyun,
	}

	if result, err := Call(funcs, providerName, c, text); err == nil {
		for _, r := range result {
			rst = r.Interface()
		}
	} else {
		fmt.Println("Handle error:", err)
	}

	return rst
}

// NlpTTS 入口
func NlpTTS(c *gin.Context) {
	providerName := c.DefaultQuery("provider_name", "xfyun")

	if err := c.ShouldBindBodyWith(&cp, binding.JSON); err != nil {
		response.FailResult(501, err.Error(), c)
	}

	fmt.Println(&cp)
	fmt.Println(cp.Data.Text)

	// items := []interface{}{}

	// rst := Handle(c, providerName, cp.Data.Text)

	// items = append(items, rst)

	// response.SuccessResult(items, c)
}
