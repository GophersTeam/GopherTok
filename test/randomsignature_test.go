package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"GopherTok/common/utils"
)

func TestSignature(t *testing.T) {
	url := "https://v1.hitokoto.cn/?encode=text"

	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求发生错误:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容错误:", err)
		return
	}

	// 将响应内容（个性签名）转换为字符串并打印
	signature := string(body)
	fmt.Println("个性签名:", signature)
}

func TestNameRandom(t *testing.T) {
	fmt.Println(utils.GetRandomYiYan())
	fmt.Println(utils.GetRandomImageUrl())
}
