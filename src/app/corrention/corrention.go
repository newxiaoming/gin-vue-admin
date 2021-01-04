package corrention

import (
	"errors"
	"fmt"
	"gin-vue-admin/internal/middleware"
	"gin-vue-admin/pkg/response"
	"net/http"
	"reflect"
	"time"

	xfyunauthorization "gin-vue-admin/pkg/xfyun/authorization"
	xfyuntextcorrention "gin-vue-admin/pkg/xfyun/textcorrention"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// textCorrection 接口
type textCorrection interface {
	corrention() string
}

// provider 服务提供者
type provider struct {
	name string
	text string
}

type CorrentionParams struct {
	Content []struct {
		Text string `json:"text" binding:"required"`
		Line string `json:"line" binding:"required"`
	} `json:"content" binding:"required"`
	Total string `json:"total" binding:"required"`
}

var textArr CorrentionParams
var (
	cp middleware.CorrentionParams
)

// corrention 文本纠错函数
func (p *provider) corrention() string {
	return p.name
}

// 讯飞文本纠错具体实现方法
func xfyun(c *gin.Context, text string) int {
	fmt.Println("running function xfyun!", c.Query("charset"))

	host := "api.xf-yun.com"
	date := time.Now().UTC().Format(http.TimeFormat)
	c.Set("host", host)
	c.Set("date", date)

	authorization := xfyunauthorization.Authorization(c)
	c.Set("authorization", authorization)
	rst := xfyuntextcorrention.PostData(c, text)
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

// Handle
func Handle(c *gin.Context, providerName string, text string) {
	funcs := map[string]interface{}{
		"xfyun": xfyun,
	}

	if result, err := Call(funcs, providerName, c, text); err == nil {
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
	providerName := c.DefaultQuery("provider_name", "xfyun")
	if err := c.ShouldBindBodyWith(&cp, binding.JSON); err != nil {
		response.FailResult(501, err.Error(), c)
	}
	// bodyData, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	response.FailResult(501, err.Error(), c)
	// }
	fmt.Println(&cp)

	// err = json.Unmarshal(bodyData, &textArr)
	// if err != nil {
	// 	response.FailResult(500, err.Error(), c)
	// }
	// fmt.Println(textArr)
	newProvider := &provider{
		name: providerName,
		text: cp.Total,
	}

	for _, v := range cp.Content {
		Handle(c, providerName, v.Text)
	}

	// var newTextCorrection textCorrection
	// newTextCorrection = newProvider
	response.SuccessResult(gin.H{
		"text": newProvider.text,
		"name": newProvider.name,
	}, c)
	// fmt.Println(newTextCorrection.corrention())
}
