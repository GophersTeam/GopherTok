package main

import (
	"GopherTok/server/chat/mq/internal/config"
	"GopherTok/server/chat/mq/internal/service"
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"sort"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the etc file")

func main() {
	flag.Parse()

	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)

	s := service.NewService(c)

	queue := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(s.Consume))
	defer queue.Stop()

	fmt.Println("chat-mq started!!!")
	queue.Start()
}

func reconstructQueue(people [][]int) [][]int {
	// 题目要求身高是降序，且前面的人数是k，所以先按照身高降序，k升序排序，完成身高降序这个维度
	res := [][]int{}
	// 按照身高降序，k升序排序
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	// 遍历排序后的数组，将元素插入到k位置
	for _, person := range people {
		k := person[1]
		// 插入到k位置
		res = append(res[:k], append([][]int{person}, res[k:]...)...)
	}

	return res
}
