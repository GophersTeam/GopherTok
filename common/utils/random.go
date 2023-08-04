package utils

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io"
	"net/http"
	"strings"
)

func GetRandomYiYan() string {
	resp, err := httpc.Do(context.Background(), http.MethodGet, "https://v1.hitokoto.cn/?encode=text", nil)
	if err != nil {
		logx.Errorf("获取一言失败: %v", err)
		return ""
	}
	defer resp.Body.Close()
	builder := strings.Builder{}
	io.Copy(&builder, resp.Body)
	return builder.String()
}

func GetRandomImageUrl() string {
	res := ""
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			res = req.URL.String()
			return nil
		},
	}
	resp, err := client.Get("https://source.unsplash.com/random")
	if err != nil {
		logx.Errorf("获取图片失败: %v", err)
		return ""
	}
	resp.Body.Close()
	return res
}
