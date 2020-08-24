package main

import "github.com/tidwall/gjson"

const jsonRow = `{"name": "Gilbert", "age": 61}
{"name": "Alexa", "age": 34}
{"name": "May", "age": 57}
{"name": "Deloise", "age": 44}`

func main() {
	println(gjson.Get(jsonRow, "..1").String())

	gjson.ForEachLine(jsonRow, func(line gjson.Result) bool {
		println(line.Get("name").String())
		return true
	})
}
