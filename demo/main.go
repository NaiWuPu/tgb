package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var str = `{
	"head": {
		"smacid": -1,
		"function": "get_auto_upgrade_module",
		"module": "auto_upgrade",
		"error_code": 0,
		"error_string": "执行成功",
		"page_index": 1,
		"page_size": 0,
		"total": -1
	},
	"body": {
		"signature": [{
			"module": "ips",
			"state": "disable",
			"cur_version": "2008311908",
			"new_version": "2008311908",
			"license": "2021-07-24 15:58:22",
			"cnname": "入侵检测库",
			"enname": "ips signature",
			"sig_count": "4161"
		}, {
			"module": "ioc",
			"state": "disable",
			"cur_version": "2008121255",
			"new_version": "2008121255",
			"license": "2021-07-24 15:58:22",
			"cnname": "威胁情报库",
			"enname": "ioc signature",
			"update_time": "2020-08-13 15:20:32",
			"update_num": "83",
			"ioc_num": "1815063",
			"sig_count": "1815063"
		}, {
			"module": "app_pro",
			"state": "disable",
			"cur_version": "2008031841_I",
			"new_version": "2008031841_I",
			"license": "2021-07-24 15:58:22",
			"cnname": "应用识别库",
			"enname": "app signature",
			"sig_count": "54478"
		}, {
			"module": "av",
			"state": "disable",
			"cur_version": "2007221759",
			"new_version": "2007221759",
			"license": "2021-07-24 15:58:22",
			"cnname": "恶意文件库",
			"enname": "vicious files signature",
			"sig_count": "36401906"
		}, {
			"module": "area",
			"state": "disable",
			"cur_version": "2006081809",
			"new_version": "2006081809",
			"license": "-",
			"cnname": "区域库",
			"enname": "area",
			"sig_count": "3373"
		}]
	}
}`

func main() {

	mdsRefundType := new(MgdRefundFeatureNds)
	err := json.Unmarshal([]byte(DecodeStr(str)), &mdsRefundType)
	fmt.Println(err)
	fmt.Println(mdsRefundType)
}

// 特征库上报mgd返回结构体
type MgdRefundFeatureNds struct {
	Head struct {
		ErrorCode   int64  `json:"error_code"`
		ErrorString string `json:"error_string"`
		Function    string `json:"function"`
		Module      string `json:"module"`
		PageIndex   int64  `json:"page_index"`
		PageSize    int64  `json:"page_size"`
		Smacid      int64  `json:"smacid"`
		Total       int64  `json:"total"`
	} `json:"head"`
	Body struct {
		Signature []struct {
			Module     string `json:"module"`
			State      string `json:"state"`
			CurVersion string `json:"cur_version"`
			NewVersion string `json:"new_version"`
			Cnname     string `json:"cnname"`
			Enname     string `json:"enname"`
			License    string `json:"license"`
			SigCount   string `json:"sig_count"`
		} `json:"signature"`
	} `json:"body"`
}

func DecodeStr(str string) string {
	return strings.Replace(str, `\`, "", -1)
}
