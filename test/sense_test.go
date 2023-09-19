package test

import (
	"fmt"
	"testing"

	"GopherTok/common/utils"
)

func TestSense(t *testing.T) {
	trie := utils.NewSensitiveTrie()
	trie.AddWords([]string{"傻逼", "死", "你妈", "滚"})
	fmt.Println(trie.Filter("傻逼玩意死马"))
}
