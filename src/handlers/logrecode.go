package handlers

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// LogRecodes recode minio logs
func LogRecodes(c *gin.Context) error {

	// alertbody := grafana.Alert{}
	// c.BindJSON(&alertbody)
	// log.Printf("receive message from grafana with alert rule[%s] and status[%s]\n", alertbody.RuleName, alertbody.State)
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))

	// log := make(map[string]interface{})
	// c.BindJSON(log)
	// fmt.Println(log)
	return nil
	// return qywchat.SendMsg(&alertbody)
}
