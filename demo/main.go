package main

import (
	"encoding/json"
	"fmt"
)

var j = `{
    "id": "20160630200323423999", 
"rule_id": 100,
    "sip": "192.168.0.12", 
    "sport": 21023, 
    "smac": "21:4c:37:52:3c:37", 
        "dip": "159.226.19.23", 
    "dport": 53, 
    "dmac": "21:4c:37:52:3c:38",
    "time": "2016-9-24 15:08:24", 
"risk": 5,
        "sender": "abc1@sina.com", 
    "receiver": "abc2@sina.com", 
    "cc": "abc3@sina.com",
"bcc": "abc4@sina.com", 
"subject": "test",
"mail_content": "This is a test.",
"attachment": ["file1.txt","file2.doc"]
}`

type AccountListenAlarmJson struct {
	MessageType string                 `json:"message_type"`
	MessageBody AccountListenAlarmBody `json:"message_body"`
}
type AccountListenAlarmBody struct {
	Id         string `json:"id" description:"告警号id"`
	RuleId     uint64 `json:"rule_id" description:"规则id"`
	Sip        string `json:"sip" description:"源ip地址"`
	Sport      uint64 `json:"sport" description:"端口号"`
	Smac       string `json:"smac" description:"源主机MAC地址"`
	Dip        string `json:"dip" description:"目的地址"`
	Dport      int64  `json:"dport" description:"目的端口号"`
	Dmac       string `json:"dmac" description:"会话目的主机MAC地址"`
	Time       string `json:"time" description:"告警发生时间"`
	Risk       int8   `json:"risk" description:"下发规则告警级别"`
	Sender     string `json:"sender" description:"发件人"`
	Receiver   string `json:"receiver" description:"收件人"`
	Cc         string `json:"cc" description:"抄送"`
	Bcc        string `json:"bcc" description:""`
	Subject    string `json:"subject" description:"邮件主题"`
	MailConte  string `json:"mail_conte" description:"邮件内容"`
	Attachment string `json:"attachment" description:"附件名列表"`
}

func main() {
	var AccountListenAlarmJson = new(AccountListenAlarmBody)
	err := json.Unmarshal([]byte(j), &AccountListenAlarmJson)
	fmt.Println(err)
	fmt.Println(AccountListenAlarmJson)
}
