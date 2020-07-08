package main

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

const (
	userName = "jcc"
	password = "jcc"
	ip       = "172.16.0.203"
	port     = "3306"
	dbName   = "city_name"
)

var db1 *gorm.DB

func init() {
	var err error

	db1, err = gorm.Open("mysql", "jcc:jcc@tcp(172.16.0.203)/city_name?charset=utf8")
	if err != nil {
		panic(err)
	}
}

type Servers struct {
	XMLName xml.Name   `xml:"address,omitempty"`
	Svs     []Province `xml:"province,name,code,omitempty"`
}

//注意，结构体中的字段必须是可导出的

type Province struct {
	XMLName xml.Name `xml:"province,omitempty"`
	Name    string   `xml:"name,attr,omitempty"`
	Code    int64    `xml:"code,attr,omitempty"`
	Lat     float64  `xml:"lat,attr,omitempty"`
	Lng     float64  `xml:"lng,attr,omitempty"`
	City    []City  `xml:"city,omitempty"`
}
type City struct {
	XMLName xml.Name `xml:"city,omitempty"`
	Name    string   `xml:"name,attr,omitempty"`
	Code    int64    `xml:"code,attr,omitempty"`
	Lat     float64  `xml:"lat,attr,omitempty"`
	Lng     float64  `xml:"lng,attr,omitempty"`
	Area    []Area  `xml:"area,omitempty"`
}
type Area struct {
	XMLName xml.Name `xml:"county"`
	Name    string   `xml:"name,attr,omitempty"`
	Code    int64    `xml:"code,attr,omitempty"`
	Lat     float64  `xml:"lat,attr,omitempty"`
	Lng     float64  `xml:"lng,attr,omitempty"`
}

type SqlPull struct {
	Id       int64   `json:"id"`
	Code     int64   `json:"code"`
	FullName string  `json:"full_name"`
	Pinyin   string  `json:"pinyin"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	ParentId int64   `json:"parent_id"`
}

func main() {
	var Sql = new([]*SqlPull)
	db1.Table("city11").Find(Sql)
	var ProvinceArr []Province
	for _, v := range *Sql {
		if v.ParentId == 0 {
			var cityArr []City
			for _, twoVlaue := range *Sql {
				if twoVlaue.ParentId == v.Id {
					var areaArr []Area
					for _, thirdValue := range *Sql {
						if thirdValue.ParentId == twoVlaue.Id {
							area := Area{Name: thirdValue.FullName, Code: thirdValue.Code, Lng: thirdValue.Lng, Lat: thirdValue.Lat}
							areaArr = append(areaArr, area)
						}
					}
					CityC := City{Name: twoVlaue.FullName, Code: twoVlaue.Code, Lng: twoVlaue.Lng, Lat: twoVlaue.Lat, Area: areaArr}
					cityArr = append(cityArr, CityC)
				}
			}
			provinceC := Province{Name: v.FullName, Code: v.Code, Lng: v.Lng, Lat: v.Lat, City: cityArr}
			ProvinceArr = append(ProvinceArr, provinceC)
		}
	}
	var servers Servers
	servers.Svs=ProvinceArr
	data, _ := xml.MarshalIndent(servers, "", "		")
	logs.Debug(ProvinceArr)
	myString := []byte(xml.Header + string(data))
	f, err := os.Create("dz.xml")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(myString))
	}

}
