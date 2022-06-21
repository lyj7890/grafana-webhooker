# 使用指南
当前 Grafana 发送告警通知不支持企业微信，该工具接收 Grafana 发送的告警并发送到企业微信个人或者群组。
# 配置步骤
1. Grafana 上 Alerting -> Notification channels -> New channel 添加 channel，
   - Type: webhook
   - Url:  `http://webhooker.monitor` 
   - Http Method: POST
   - 其他字段按需修改，Name 不能和其他重复
   - 此操作需要管理员权限
   
2. 在 Dashboard 中配置告警条件。
   
   
3. 添加的告警接收人并配置tags：
   
   webhook 根据 tags 的配置将消息通知发送到个人或者群组
   - 接收人为企业微信个人：key=receivers,vlaues=`<user ID>`
   - 接收人为企业微信群组，需添加群机器人：key=group, value=`<bot key>`
   
   注意：

   > `<user ID>`：为企业微信 ID,如果有多个接收人使用','隔开
   >
   > `<bot key>`：机器人账号 webhook 为 `https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxxxxxx` 这种格式，这里只保留 `key` 值，即 `xxxx` 所对应的内容。
4. 当满足告警条件发送告警通知：
   

   - 如果配置了告警恢复，当告警恢复时会立即发送告警恢复通知
