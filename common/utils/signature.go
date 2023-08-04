package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RandomSignature() (string, error) {
	url := "https://v1.hitokoto.cn/?encode=text"

	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求发生错误:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容错误:", err)
		return "", err
	}

	// 将响应内容（个性签名）转换为字符串并打印
	signature := string(body)
	fmt.Println("个性签名:", signature)
	return signature, nil
}
