package textcorrention

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gin-vue-admin/pkg/response"
	"gin-vue-admin/tool"
	"net/url"

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
	Char   [][]interface{} `json:"char"`
	Word   [][]interface{} `json:"word"`
	Redund [][]interface{} `json:"redund"`
	Miss   [][]interface{} `json:"miss"`
	Order  [][]interface{} `json:"order"`
	Dapei  [][]interface{} `json:"dapei"`
	Punc   [][]interface{} `json:"punc"`
	Idm    [][]interface{} `json:"idm"`
	Org    [][]interface{} `json:"org"`
	Leader [][]interface{} `json:"leader"`
	Number [][]interface{} `json:"number"`
}

type textDataItem struct {
	Pos         int
	Cur         string
	Corrent     string
	Description string
}

type textDataItemJSON struct {
	*textDataItem
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
	textArr      []map[string]interface{}
)

// PostData 提交数据
func PostData(c *gin.Context, text string) response.TextCorrentionResponseItem {
	var (
		data response.TextCorrentionResponseItem
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
	fmt.Println(result)
	lists, err := formatData([]byte(result))

	if err != nil {
		fmt.Println("PostData:", err)
	}
	data.Text = text
	data.VecFragment = lists

	return data
}

// formatData 格式化返回的数据
func formatData(data []byte) ([]map[string]interface{}, error) {
	var items response.TextCorrentionResponseItem
	var lists []map[string]interface{}

	if err := json.Unmarshal(data, &respData); err != nil {
		return nil, errors.New("responseData is no json")
	}

	if respData.Header.Code != 0 {
		return nil, errors.New("request error")
	}

	textBase64Data, err := base64.StdEncoding.DecodeString(respData.Payload.Result.Text)
	fmt.Println("textBase64Data:", string(textBase64Data))
	if err != nil {
		return nil, errors.New("responseData text is no base64")
	}
	fmt.Println("====")
	fmt.Println(string(textBase64Data))
	fmt.Println("====")
	if err := json.Unmarshal(textBase64Data, &textOrigData); err != nil {
		fmt.Println("textOrigData error :", err)
		return nil, errors.New("textOrigData is no json")
	}

	fmt.Println("--------textArr----------")
	textDataItems := &textData{}
	err = json.Unmarshal(textBase64Data, textDataItems)
	fmt.Println(textDataItems)

	getItem(textDataItems.Char)
	getItem(textDataItems.Word)
	getItem(textDataItems.Redund)
	getItem(textDataItems.Miss)
	getItem(textDataItems.Order)
	getItem(textDataItems.Dapei)
	getItem(textDataItems.Punc)
	getItem(textDataItems.Idm)
	getItem(textDataItems.Org)
	getItem(textDataItems.Leader)
	getItem(textDataItems.Number)

	lists = textArr
	textArr = make([]map[string]interface{}, 0)

	fmt.Println("--------lists start----------")
	fmt.Println(lists)
	fmt.Println("--------lists end----------")
	d, _ := json.Marshal(items)
	fmt.Println(string(d))
	fmt.Println("---------textArr---------")

	fmt.Println(items)

	return lists, nil
}

func getItem(inputItem [][]interface{}) []map[string]interface{} {
	var lists []map[string]interface{}
	item := map[string]interface{}{}

	// items := response.TextCorrentionResponseItem{}
	if len(inputItem) > 0 {
		for _, value := range inputItem {
			item = map[string]interface{}{
				"OriFrag":     tool.Strval(value[1]),
				"BeginPos":    tool.Strval(value[0]),
				"CorrectFrag": tool.Strval(value[2]),
				"EndPos":      "0",
			}

			textArr = append(textArr, item)
		}
	}

	fmt.Println("========================================================================")
	fmt.Println(textArr)
	// fmt.Println(items)
	fmt.Println("========================================================================")
	return lists
}
