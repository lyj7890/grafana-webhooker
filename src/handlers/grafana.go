package handlers

import (
	"log"
	"webhooker/src/grafana"
	"webhooker/src/qywchat"

	"github.com/gin-gonic/gin"
)

func Do(c *gin.Context) error {
	alertbody := grafana.Alert{}
	c.BindJSON(&alertbody)
	log.Printf("receive message from grafana with alert rule[%s] and status[%s]\n", alertbody.RuleName, alertbody.State)

	return qywchat.SendMsg(&alertbody)
}
