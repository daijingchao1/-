package main

import (
	"fmt"
	"io/ioutil")


func main() {
	content,_:=ioutil.ReadFile("aaa.xml")
	fmt.Printf("content:%s\n",content)
}