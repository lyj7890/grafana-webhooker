package qywchat

import "webhooker/src/grafana"

const (
	CorpID = "000000000000"

	Corpsecret = "MODIFY-TO-YOURSELF"
	AgentId    = 99999
)

var (
	URLGetToken    = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	URLGetUID      = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserid"
	URLGetUserInfo = "https://qyapi.weixin.qq.com/cgi-bin/user/get"
	URLSendMsg     = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
	URLSendWebHook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	AlertStatusMap = map[string]string{
		grafana.AlertStatusOk:    "已恢复",
		grafana.AlertStatusAlert: "告警中",
	}
)
