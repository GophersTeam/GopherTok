package test

import (
	"fmt"
	"testing"
	"time"
)

func TestUnix(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

}
func TestU(y *testing.T) {
	// 假设你有一个Unix时间戳（以秒为单位）
	unixTimestamp := int64(1678981123)

	// 将Unix时间戳转换为毫秒
	milliseconds := unixTimestamp * 1000

	// 将毫秒转换为时间
	t := time.Unix(0, milliseconds*int64(time.Millisecond))

	fmt.Println("Unix Timestamp (seconds):", unixTimestamp)
	fmt.Println("Timestamp in milliseconds:", milliseconds)
	fmt.Println("Converted Time:", t)
}
