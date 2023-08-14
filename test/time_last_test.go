package test

import (
	"fmt"
	"testing"
)

func TestTimeLast(t *testing.T) {
	timeStr1 := "999"                  // 示例时间字符串
	timeStr2 := "2023-08-15T12:00:00Z" // 另一个示例时间字符串

	if timeStr1 < timeStr2 {
		fmt.Println("时间1在时间2之前")
	} else if timeStr1 > timeStr2 {
		fmt.Println("时间1在时间2之后")
	} else {
		fmt.Println("时间1和时间2相等")
	}
}
