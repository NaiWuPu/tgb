package main

import (
	"fmt"
	"strings"
)

var str = `{\"request_list\":{\"head\":{\"current_vsys_name\":\"root-vsys\",\"target_vsys_name\":\"root-vsys\",\"from\":\"webui\",\"user\":\"admin\",\"language\":\"CN\",\"sessionid\":0},\"body\":{\"request\":[{\"head\":{\"module\":\"info_collect\",\"function\":\"set_send_policy\"},\"body\":{\"info_collect\":[{\"cmdId\":\"asd332123456qw4\",\"cmdType\":2,\"devId\":\"73\",\"timestamp\":\"1598425575696\",\"name\":\"sftp_sender\",\"desc\":\"sftp\",\"status\":\"0\",\"logtype\":[],\"sendserver\":[{\"addr\":\"1.2.3.43\",\"port\":\"1234\"}],\"regular\":\"\",\"mode\":\"bjwa\",\"auth\":\"\",\"compress\":\"\",\"encrypt\":\"\",\"encoding\":\"\",\"enc_key\":\"\",\"transproto\":\"udp\",\"protocol\":\"\",\"transfer\":\"ftp\",\"fileType\":\"\",\"username\":\"admin22\",\"password\":\"123422\",\"devName\":\"admin2\",\"devPwd\":\"12345678\",\"vendorId\":2}]}}]}}}`

func main() {
	addInfoCollect([]byte(str))
}

func addInfoCollect(aesBody []byte) {
	collectAddString := string(aesBody)
	fmt.Println(collectAddString)

	pdf := strings.Replace(collectAddString, "set_send_policy", "add_send_policy", -1)
	// mgd 解析
	fmt.Println(pdf)
}

type MgdRefundType struct {
	Head struct {
		Smacid      int64  `json:"smacid"`
		Function    string `json:"function"`
		Module      string `json:"module"`
		ErrorCode   int64  `json:"error_code"`
		ErrorString string `json:"error_string"`
		PageIndex   int64  `json:"page_index"`
		PageSize    int64  `json:"page_size"`
		Total       int64  `json:"total"`
	} `json:"head"`
	Body struct {
	}
}
