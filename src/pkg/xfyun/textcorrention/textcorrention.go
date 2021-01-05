package textcorrention

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gin-vue-admin/pkg/response"
	"gin-vue-admin/tool"
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Header struct {
		AppID  string `json:"app_id"`
		Status int    `json:"status"`
	} `json:"header"`
	Parameter struct {
		S9a87e3ec struct {
			Result struct {
				Encoding string `json:"encoding"`
				Compress string `json:"compress"`
				Format   string `json:"format"`
			} `json:"result"`
		} `json:"s9a87e3ec"`
	} `json:"parameter"`
	Payload struct {
		Input struct {
			Encoding string `json:"encoding"`
			Compress string `json:"compress"`
			Format   string `json:"format"`
			Status   int    `json:"status"`
			Text     string `json:"text"`
		} `json:"input"`
	} `json:"payload"`
}

type textData struct {
	Char   interface{} `json:"char"`
	Word   interface{} `json:"word"`
	Redund interface{} `json:"redund"`
	Miss   interface{} `json:"miss"`
	Order  interface{} `json:"order"`
	Dapei  interface{} `json:"dapei"`
	Punc   interface{} `json:"punc"`
	Idm    interface{} `json:"idm"`
	Org    interface{} `json:"org"`
	Leader interface{} `json:"leader"`
	Number interface{} `json:"number"`
}

type responseData struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		SID     string `json:"sid"`
	} `json:"header"`
	Payload struct {
		Result struct {
			Compress string `json:"compress"`
			Encoding string `json:"encoding"`
			Format   string `json:"format"`
			Seq      string `json:"seq"`
			Status   string `json:"status"`
			Text     string `json:"text"`
		} `json:"result"`
	} `json:"payload"`
}

var (
	respData     responseData
	textOrigData textData
)

// PostData 提交数据
func PostData(c *gin.Context, text string) string {
	var (
		data []byte
	)
	uri := "https://api.xf-yun.com/v1/private/s9a87e3ec?authorization=%s&host=%s&date=%s"
	contentType := "application/json"
	reqURL := fmt.Sprintf(uri, c.GetString("authorization"), c.GetString("host"), url.QueryEscape(c.GetString("date")))
	fmt.Println(reqURL)

	reqBody := &requestBody{}
	reqBody.Header.AppID = tool.GetConfig("xfyun.textcorrention.appid")
	reqBody.Header.Status = 3
	reqBody.Parameter.S9a87e3ec.Result.Encoding = "utf8"
	reqBody.Parameter.S9a87e3ec.Result.Compress = "raw"
	reqBody.Parameter.S9a87e3ec.Result.Format = "json"
	reqBody.Payload.Input.Encoding = "utf8"
	reqBody.Payload.Input.Compress = "raw"
	reqBody.Payload.Input.Format = "json"
	reqBody.Payload.Input.Status = 3
	reqBody.Payload.Input.Text = base64.StdEncoding.EncodeToString([]byte(text))

	r, err := json.Marshal(reqBody)
	if err != nil {
		response.FailResult(500, err.Error(), c)
	}

	result := tool.HttpPost(reqURL, contentType, string(r))
	if data, err = formatData([]byte(result)); err != nil {
		fmt.Println(err)
	}

	// if v, err = json.Marshal(data); err != nil {
	// 	fmt.Println(err)
	// }

	return string(data)
}

// formatData 格式化返回的数据
func formatData(data []byte) ([]byte, error) {
	var v []byte
	// item := map[string]interface{}{}
	// items := response.TextCorrentionResponseItem{}

	if err := json.Unmarshal(data, &respData); err != nil {
		return nil, errors.New("responseData is no json")
	}

	if respData.Header.Code != 0 {
		return nil, errors.New("request error")
	}

	textBase64Data, err := base64.StdEncoding.DecodeString(respData.Payload.Result.Text)
	fmt.Println(string(textBase64Data))
	if err != nil {
		return nil, errors.New("responseData text is no base64")
	}

	if err := json.Unmarshal(textBase64Data, &textOrigData); err != nil {
		fmt.Println(err)
		return nil, errors.New("textOrigData is no json")
	}

	// 别字纠错
	// if reflect.ValueOf(textOrigData.Char).Len() > 0 {
	// 	for _, value := range textOrigData.Char.([]interface{}) {
	// 		item["OriFrag"] = tool.Strval(value.([]interface{})[1])
	// 		item["BeginPos"] = tool.Strval(value.([]interface{})[0])
	// 		item["CorrectFrag"] = tool.Strval(value.([]interface{})[2])
	// 		item["EndPos"] = 0
	// 	}
	// }

	// getItem(textOrigData.Char)
	// for _, t := range textOrigData. {
	// 	getItem(t)
	// }

	k := reflect.ValueOf(textOrigData)
	count := k.NumField()
	for i := 0; i < count; i++ {
		f := k.Field(i)
		switch f.Kind() {
		case reflect.Interface:
			fmt.Println(f)
			getItem(f.Interface)
		}
	}
	// getItem(textOrigData.Word)
	// getItem(textOrigData.Redund)
	// getItem(textOrigData.Redund)
	// getItem(textOrigData.Redund)
	// getItem(textOrigData.Redund)
	// getItem(textOrigData.Redund)
	// getItem(textOrigData.Redund)
	// getItem(textOrigData.Redund)
	// getItem(textOrigData.Punc)

	if v, err = json.Marshal(textOrigData); err != nil {
		fmt.Println(err)
	}

	return v, nil
}

func getItem(data interface{}) []interface{} {

	items := []interface{}{}

	if data != nil {
		for _, value := range data.([]interface{}) {
			item := map[string]interface{}{}

			item["OriFrag"] = tool.Strval(value.([]interface{})[1])
			item["BeginPos"] = tool.Strval(value.([]interface{})[0])
			item["CorrectFrag"] = tool.Strval(value.([]interface{})[2])
			item["EndPos"] = 0
			fmt.Println(value)
			items = append(items, item)
		}
	}
	return items
}
