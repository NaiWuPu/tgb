package main

import (
	"encoding/json"
	"fmt"
)

type Category struct {
	Id int
	Pid int
	Name string
	Children []*Category
}

var ListArr  = []*Category{
	{1, 0, "001", []*Category{}},
	{2, 1, "001001", []*Category{}},
	{3, 2, "001001001", []*Category{}},
	{4, 2, "001001002", []*Category{}},
	{5, 1, "001002", []*Category{}},
	{6, 5, "001002001", []*Category{}},
	{7, 5, "001002002", []*Category{}},
}

func main() {
	mapArr := make(map[int]*Category)
	for _, value := range ListArr {
		mapArr[value.Id] = value
	}

	for _, value := range mapArr {
		if _, ok := mapArr[value.Pid]; ok {
			t := mapArr[value.Pid]
			t.Children = append(t.Children, value)
			mapArr[value.Pid] = t
			delete(mapArr, value.Id)
		}
	}

	s, _:=json.Marshal(mapArr)

	fmt.Println(string(s))

}