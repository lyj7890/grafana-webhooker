package grafana

import (
	"fmt"
	"strings"
)

type EvalMatche struct {
	Metric string            `json:"metric"`
	Value  float64           `json:"value"`
	Tags   map[string]string `json:"tags"`
}

var (
	AlertStatusOk    string = "ok"
	AlertStatusAlert string = "alerting"
)

type Alert struct {
	DashboardID int               `json:"dashboard"`
	EvalMatches []EvalMatche      `json:"evalMatches"`
	ImageURL    string            `json:"imageUrl"`
	OrgID       int               `json:"orgId"`
	PanelID     int               `json:"panelId"`
	RuleID      int               `json:"ruleId"`
	RuleName    string            `json:"ruleName"`
	RuleURL     string            `json:"ruleUrl"`
	State       string            `json:"state"`
	Tags        map[string]string `json:"tags"`
	Title       string            `json:"title"`
	Message     string            `json:"message"`
}

// GetReceiverArray 根据 receivers tag 获取uid
func (a *Alert) GetReceiverArray() []string {
	var res []string
	if v, ok := a.Tags["receivers"]; ok {
		res = strings.Split(v, ",")
	}
	return res
}

func (a *Alert) GetEvalMatche() string {
	if len(a.EvalMatches) == 0 {
		return ""
	}
	var res string
	for _, evalMatche := range a.EvalMatches {
		line := fmt.Sprintf("\tMatrics:%s\t\tValue:%.2f\n", evalMatche.Metric, evalMatche.Value)
		res = fmt.Sprintf(res + line)
	}
	return res
}

// GetBotHooks 根据groups tag获取机器人webhooks
// func (a *Alert) GetBotHook() string {
// 	var res string
// 	if v, ok := a.Tags["groups"]; ok {
// 		res = paseUrl(v)
// 	}
// 	return res
// }

// func paseUrl(url string) string {
// 	if govalidator.IsURL(url) {
// 		return url
// 	}
// 	return ""
// }

func (a *Alert) GetGroup() string {
	var res string
	if v, ok := a.Tags["group"]; ok {
		res = v
	}
	return res
}

func (a *Alert) GetAlertRule() string {
	return a.RuleName
}

func (a *Alert) GetImageURL() string {
	return a.ImageURL
}

func (a *Alert) GetAlertMsg() string {
	return a.Message
}

// webhook body example:
// {
// 	"dashboardId":1,
// 	"evalMatches":[
// 	  {
// 		"value":1,
// 		"metric":"Count",
// 		"tags":{}
// 	  }
// 	],
// 	"imageUrl":"https://grafana.com/assets/img/blog/mixed_styles.png",
// 	"message":"Notification Message",
// 	"orgId":1,
// 	"panelId":2,
// 	"ruleId":1,
// 	"ruleName":"Panel Title alert",
// 	"ruleUrl":"http://localhost:3000/d/hZ7BuVbWz/test-dashboard?fullscreen\u0026edit\u0026tab=alert\u0026panelId=2\u0026orgId=1",
// 	"state":"alerting",
// 	"tags":{
// 	  "tag name":"tag value"
// 	},
// 	"title":"[Alerting] Panel Title alert"
// }
