package main

import (
	"fmt"
	"os/exec"
)
var tet = "你晋级了，下次再来"

func main() {
	//fmt.Printf("%#v\n",os.Args)
	//fmt.Printf("%T/n",os.Args)
	cmd := exec.Command("a.exe")
	bt, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bt))
}
