package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
    b, err := ioutil.ReadFile("file.txt") 
    if err != nil {
        fmt.Print(err)
    }
	s := strings.Fields(string(b))
	for i:=0;i<len(s);i++ {
		sep:="/"
		res:=strings.Split(s[i],sep)
		fmt.Println(res[len(res)-1])
	}
}