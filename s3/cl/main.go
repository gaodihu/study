package main

import (
	"fmt"

	"github.com/silenceper/wechat/template"
	"github.com/silenceper/wechat/util"
	"encoding/json"
	"strings"
)

func main() {

	value := `�contentType   "text/plain"originalContentType    "application/json;charset=UTF-8"{"msgType":"FresherGiftUsed","mallCode":"0301A404","userCode":"dca3f4f9259a4a10ba326b2f2c52334c","data":{"giftName":"新人礼","giftId":"954083908fb649248a3dd39f2fc0ba51","couponNo":"SH22c12200007160","status":2}}`
	first := strings.Index(value,"{")
	value = value[first:]
	fmt.Println(value)


	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(value),&m)
	if err == nil {
		fmt.Println(m)
		if d,ok := m["data"];ok {
			fmt.Println(d)
			dd := d.(map[string]interface{})
			if g,ok := dd["giftId"];ok {
				wxm := template.Message{}
				wxm.TemplateID = "jw6dYTSWX6-sbgX7FX9bhgQzKCj3mnSRHUvnHIGShHU"
				wxm.ToUser = "ogjRFwZwoX8VXV1Rxip9LqMsLj54"
				wxm.Data = nil
				wxm.URL = "http://www.qq.com"

				WechatTempltenSend(&wxm,"9_TJbyA3uqH9p56KkBo6qSx15s1b90fMEhr1oBWqC-2IaTli_ubD9sqGHe7Q4lcrK6SzmhEK_7mbmjrvhTIrAn52mFwtiAjIj-9fCqIdMZXbvn_VYsIC3ZIZKgUBSTyGS_25LA14pDFeYaAwE2WHQbAHABMV")

			}
		}
	}



	/*
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s,%s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			value := string(msg.Value)

			first := strings.Index(value,"{")
			value = value[first:]

			m := make(map[string]interface{})
			err = json.Unmarshal([]byte(value),&m)
			if err == nil {

			}


		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
	*/

}
func WechatTempltenSend(msg *template.Message, accessToken string)  (msgID int64, err error) {
	templateSendURL := "https://api.weixin.qq.com/cgi-bin/message/template/send"

	uri := fmt.Sprintf("%s?access_token=%s", templateSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)


	type resTemplateSend struct {
		util.CommonError

		MsgID int64 `json:"msgid"`
	}
	var result resTemplateSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	msgID = result.MsgID
	return
}