package main

import (
	"encoding/json"
	"examples/text_json"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
)

const (
	userName = "jcc"
	password = "jcc"
	ip       = "172.16.0.203"
	port     = "3306"
	dbName   = "city_name"
)

type List struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Fullname string   `json:"fullname"`
	Pinyin   []string `json:"pinyin"`
	Location Location `json:"location"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type SqlInstall struct {
	Code     int64   `json:"code"`
	FullName string  `json:"full_name"`
	Pinyin   string  `json:"pinyin"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	ParentId int64   `json:"parent_id"`
}

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open("mysql", "jcc:jcc@tcp(172.16.0.203)/city_name?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func main() {
	List := make([]List, 0)
	err := json.Unmarshal([]byte(text_json.TextJson), &List)
	if err != nil {
		log.Fatal(err)
		return
	}
	c := 1
	i := 0
	m := make(map[int]int, 100)
	n := make(map[int]int)
	for _, v := range List {
		var provinceCode, _ = strconv.Atoi(v.Id[:2])
		var cityCode, _ = strconv.Atoi(v.Id[2:4])
		var areaCode, _ = strconv.Atoi(v.Id[4:])
		var f4Code, _ = strconv.Atoi(v.Id[:4])
		for province := 11; province < 72; province++ {
			i++
			ok := m[province]
			if ok == 0 {
				m[province] = i
			}
			for city := 0; city < 100; city++ {
				for area := 0; area < 100; area++ {
					if province == provinceCode {
						if city == cityCode {
							if area == areaCode {
								// 计算区的父id
								if areaCode == 0 {
									logs.Debug(c)
									n[f4Code] = c
									//logs.Debug(n)
								}
								s := new(SqlInstall)
								if len(v.Pinyin) > 0 {
									pinyin := ""
									for _, value := range v.Pinyin {
										pinyin += value
									}
									s.Pinyin = pinyin
								}
								s.Lat = v.Location.Lat
								s.Lng = v.Location.Lng
								s.FullName = v.Fullname
								Code, _ := strconv.Atoi(v.Id)
								s.Code = int64(Code)
								if city == 0 && area == 0 {
									s.ParentId = 0
								}
								if city != 0 && area == 0 {
									s.ParentId = int64(m[province])
								}
								if city != 0 && area != 0 {
									s.ParentId = int64(n[f4Code])
									if provinceCode == 11 {
										s.ParentId = 1
									}
									if provinceCode == 12 {
										s.ParentId = 2
									}
								}
								//logs.Debug(s)
								db.Table("city11").Create(&s)
								c++
								logs.Debug(c)
							}
						}
					}
				}
			}
		}
	}

}
