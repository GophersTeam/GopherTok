package test

import (
	"fmt"
	"testing"
	"time"
)

func TestUnix(t *testing.T) {
	fmt.Println(time.Now().Unix())
}
