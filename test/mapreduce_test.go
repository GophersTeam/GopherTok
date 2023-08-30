package test

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"testing"
)

func TestMapReduce(t *testing.T) {
	// 数据源，生成0-9的整数
	genF := func(source chan<- int) {
		for i := 0; i < 10; i++ {
			source <- i
		}
		// 数据出错时，可以调用cancel，这样会中断整个流程
		//close(source)
	}

	// 数据处理，将每个整数转换为字符串
	mapF := func(item int, writer mr.Writer[string], cancel func(error)) {
		// 数据处理出错时，可以调用cancel，这样会中断整个流程
		//if item == 5 {
		//	cancel(errors.New("item is 5"))
		//}

		// 处理后的字符串写进writer里
		idStr := fmt.Sprintf("%d", item)
		// mapF是并发执行的，所以这里的输出顺序是不确定的
		writer.Write(idStr)
	}

	// 做聚合处理，变成字符串数组，从pipe取出每个处理后的数，结果写进writer里
	reducerF := func(pipe <-chan string, writer mr.Writer[[]string], cancel func(error)) {
		var result []string
		for item := range pipe {
			result = append(result, item)
		}
		writer.Write(result)
	}

	result, err := mr.MapReduce(genF, mapF, reducerF)

	if err != nil {
		fmt.Println(err)
		return
	}
	// 会输出0-9的字符串数组，顺序不确定
	fmt.Println(result)
}
