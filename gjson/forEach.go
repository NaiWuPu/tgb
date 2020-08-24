package main

import "github.com/tidwall/gjson"

const json1 = `{"name":{"name":"zhangsan","age":47},"name1":{"name":"lisi","age":42}}`

// 遍历
func main() {
	gjson.Get(json1, "name1").ForEach(printKeyValue())
	var json = `{"name":"zhangsan","age":47}`
	gjson.Parse(json).ForEach(printKeyValue1())

}

func printKeyValue() func(key gjson.Result, value gjson.Result) bool {
	return func(key, value gjson.Result) bool {
		println(key.String(), ":", value.String())
		return true
	}
}

func printKeyValue1() func(key gjson.Result, value gjson.Result) bool {
	return func(key, value gjson.Result) bool {
		println(key.String(), ":", value.String())
		return true
	}
}
