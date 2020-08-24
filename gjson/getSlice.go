package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const jsonArray = `{
  "names": [
    {
      "name": "zhangsan"
    },
    {
      "name": "lisi"
    }
  ]
}`

const jsonArray2 = `[
    {
      "name": "zhangsan",
      "address": "beijing"
    },
    {
      "name": "lisi",
      "address": "tianjin"
    }
  ]
`

func main() {
	value := gjson.Get(jsonArray, "names.0.name")
	println(value.String())
	value = gjson.Get(jsonArray, "names.1.name")
	println(value.String())

	valur := gjson.Get(jsonArray2, "0.name")
	fmt.Println(valur.String())
}
