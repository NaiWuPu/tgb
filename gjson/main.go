package main

import "github.com/tidwall/gjson"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

// 获取值
func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())
	results := gjson.GetMany(json, "name.last", "age")
	for _, result := range results {
		println(result.String())
	}
	value = gjson.GetBytes([]byte(json), "name.last")
	println(value.String())
}
