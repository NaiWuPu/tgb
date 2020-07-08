package main

import (
	"encoding/xml"

	"fmt"
)

type Servers struct {
	XMLName xml.Name   `xml:"address,omitempty"`
	Svs     []Province `xml:"province,omitempty"`
}

type Province struct {
	City City `xml:"city,omitempty"`
}

type City struct {
	Name   string `xml:"name,attr,omitempty"`
	County County `xml:"county,omitempty"`
}
type County struct {
	Name string `xml:"name,attr,omitempty"`
	Code string `xml:"code,attr,omitempty"`
}

func main() {
	v := Servers{}
	P := make([]Province, 0)
	P = append(P, Province{City{"测试1", County{"白沙乡1", "1116622"}}})
	P = append(P, Province{City{"测试2", County{"白沙乡2", "1116622"}}})
	v.Svs = P

	output, err := xml.MarshalIndent(v, "", "	")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	myString := []byte(xml.Header + string(output))
	//将字节流转换成string输出
	fmt.Println(string(myString))

}
