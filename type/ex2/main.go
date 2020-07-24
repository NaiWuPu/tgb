package main

import (
	"encoding/json"
)

type Example struct {
	A []byte `json:"a"`
	B string `json:"b"`
}

func f1() {
	var Ex = new(Example)
	str := "i have an apple"
	strJson, _ := json.Marshal(str)
	Ex.A = strJson
	Ex.B = "i have a banana"
	//fmt.Println(Ex)
}

func f2() {
	var Ex = new(Example)
	str := "i have an apple"
	strJson, _ := json.Marshal(str)
	Ex.A = make([]byte, len([]byte(strJson)))
	Ex.B = "i have a banana"
	copy(Ex.A, []byte(strJson))
	//fmt.Println(Ex)
}

func main() {
	f1()
	f2()
}
