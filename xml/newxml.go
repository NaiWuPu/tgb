package main

import (
	"encoding/json"
	"encoding/xml"
	"examples/text_json"
	"fmt"
	"log"
	"os"
)


type Servers struct {
	XMLName xml.Name   `xml:"address,omitempty"`
	Svs     []Province `xml:"province,omitempty"`
}

type Province struct {
	Name string `xml:"name,attr,omitempty"`
	Code string `xml:"code,attr,omitempty"`
	City []City `xml:"city,omitempty"`
}
type City struct {
	Name   string   `xml:"name,attr,omitempty"`
	Code   string   `xml:"code,attr,omitempty"`
	County []County `xml:"county,omitempty"`
}
type County struct {
	Name string `xml:"name,attr,omitempty"`
	Code string `xml:"code,attr,omitempty"`
}
type Lists struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
}

func main() {
	// 解析text_json
	List := make([]Lists, 0)
	var TextJson = text_json.TextJson
	err := json.Unmarshal([]byte(TextJson), &List)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 构建结构数据
	v := Servers{}
	P := make([]Province, 0)
	for _, firstList := range List {
		var f02 = firstList.Id[:2]
		var f26 = firstList.Id[2:]
		// 第一层，拿出省的地址
		if f26 == "0000" {
			C := make([]City, 0)
			for _, twoList := range List {
				//var t04 = twoList.Id[:4]
				var t02 = twoList.Id[:2]
				var t04 = twoList.Id[:4]
				var t26 = twoList.Id[2:]
				var t46 = twoList.Id[4:]
				if t02 == f02 && t46 == "00" && t26 != "0000"{
					D := make([]County, 0)
					for _, thirdList := range List{
						var th04 = thirdList.Id[:4]
						var th46 = thirdList.Id[4:]
						if t04 == th04 && th46 != "00" {
							D = append(D, County{thirdList.Fullname, thirdList.Id})
						}
					}
					C = append(C, City{twoList.Fullname, twoList.Id, D})
				}
			}
			if f02 == "11" || f02 == "12" || f02 =="50" || f02 =="81" || f02 == "82" || f02 == "31"  {
				D := make([]County, 0)
				for _, ts := range List {
					var ts02 = ts.Id[0:2]
					var ts26 = ts.Id[2:]
					if ts02 == f02 && ts26 != "0000"{
						D = append(D, County{ts.Fullname, ts.Id})
					}
				}
				C = append(C, City{firstList.Fullname, f02+"0100", D})
			}
			// 最外层无误
			P = append(P, Province{firstList.Fullname, firstList.Id, C})
		}
		v.Svs = P


		// 输出xml 文件
		data, _ := xml.MarshalIndent(v, "", "		")
		myString := []byte(xml.Header + string(data))
		f, err := os.Create("dz.xml")
		defer f.Close()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_, err = f.Write(myString)
		}
	}
}
