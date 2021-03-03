package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//SConfig 指定数据类型
type SConfig struct {
	XMLName      xml.Name   `xml:"config"`     // 指定最外层的标签为config
	SMTPServer   string     `xml:"smtpServer"` //SmtpServer 读取smtpServer配置项，并将结果保存到SmtpServer变量中
	SMTPPort     int        `xml:"smtpPort"`
	Sender       string     `xml:"sender"`
	SenderPasswd string     `xml:"senderPasswd"`
	Receivers    SReceivers `xml:"receivers"` // 读取receivers标签下的内容，以结构方式获取
}

//SReceivers 指定数据类型
type SReceivers struct {
	Flag string   `xml:"flag,attr"` // 读取flag属性
	User []string `xml:"user"`      // 读取user数组
}

func main() {
	file, err := os.Open("aaa.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := SConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
