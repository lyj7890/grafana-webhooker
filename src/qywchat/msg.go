package qywchat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"webhooker/src/grafana"
)

type MsgToUser struct {
	User    string `json:"touser"`
	AgentID int    `json:"agentid"`
	Msg
}

type MsgToWebhook struct {
	Msg
}

type Msg struct {
	Type     string       `json:"msgtype"`
	Text     TextContent  `json:"text,omitempty"`
	Markdown TextContent  `json:"markdown,omitempty"`
	Image    ImageContent `json:"image,omitempty"`
}

type TextContent struct {
	Content string `json:"content"`
}
type ImageContent struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

func GetMsg2User(msg string, users string) *MsgToUser {
	return &MsgToUser{
		User:    users,
		AgentID: AgentId,
		Msg: Msg{
			Type: "text",
			Text: TextContent{
				Content: msg,
			},
		},
	}
}

func GetMsg2Hook(msg string) *MsgToWebhook {
	return &MsgToWebhook{
		Msg: Msg{
			Type: "text",
			Text: TextContent{
				Content: msg,
			},
		},
	}
}

func SendMsg(a *grafana.Alert) (err error) {
	log.Printf("receive alert[%s] from grafana.\n", a.RuleName)
	receivers := a.GetReceiverArray()
	if len(receivers) > 0 {
		err = sendMsgToUser(a)
	} else {
		log.Println("no receivers")
	}

	group := a.GetGroup()
	if group != "" {
		err = sendMsgToGroup(a, group)
	} else {
		log.Println("no group")
	}

	if err != nil {
		log.Printf("send all alert[%s] finished, but got error %+v\n", a.RuleName, err)
		return err
	}

	return nil
}

func sendMsgToUser(a *grafana.Alert) error {
	query := url.Values{}
	query.Set("access_token", Token)

	u := fmt.Sprintf(URLSendMsg + "?" + query.Encode())
	log.Printf("send alert[%s] to user\n", a.RuleName)

	var users string
	for _, user := range a.GetReceiverArray() {
		users = fmt.Sprintf(user + "|" + users)
	}

	msg := GetMsg2User(alertMsg(a), users)
	body, _ := json.Marshal(msg)
	_, err := HTTPRequest("POST", u, string(body))
	if err != nil {
		return fmt.Errorf("send message to user err: %v", err)
	}

	return nil
}

func sendMsgToGroup(a *grafana.Alert, group string) error {
	query := url.Values{}
	query.Set("key", group)

	u := fmt.Sprintf(URLSendWebHook + "?" + query.Encode())
	log.Printf("send alert[%s] to group[%s]\n", a.RuleName, group)

	msg := GetMsg2Hook(alertMsg(a))
	body, _ := json.Marshal(msg)
	_, err := HTTPRequest("POST", u, string(body))
	if err != nil {
		return fmt.Errorf("send message to group err: %v", err)
	}

	return nil
}

func alertMsg(a *grafana.Alert) string {
	data := "Grafana 告警通知\n告警状态：%s\n告警主题：%s\n告警指标：\n%s告警详情：%s"
	return fmt.Sprintf(data, AlertStatusMap[a.State], a.RuleName, a.GetEvalMatche(), a.Message)
}
