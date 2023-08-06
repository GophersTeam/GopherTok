package utils

import (
	"regexp"
	"strings"
)

type SensitiveWordFilter interface {
	Filter(content string) string
}

// SensitiveTrie 敏感词前缀树
type SensitiveTrie struct {
	replaceChar rune // 敏感词替换的字符
	root        *TrieNode
}

func (st *SensitiveTrie) Filter(content string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	content = re.ReplaceAllString(content, "")      // 去除html标签
	content = strings.Replace(content, " ", "", -1) // 去除空格
	_, replaceText := st.Match(content)             // 过滤敏感词
	return replaceText
}

// NewSensitiveTrie 构造敏感词前缀树实例
func NewSensitiveTrie() *SensitiveTrie {
	return &SensitiveTrie{
		replaceChar: '*',
		root:        &TrieNode{End: false},
	}
}

// FilterSpecialChar 过滤特殊字符
func (st *SensitiveTrie) FilterSpecialChar(text string) string {
	text = strings.ToLower(text)
	text = strings.Replace(text, " ", "", -1) // 去除空格

	// 过滤除中英文及数字以外的其他字符
	otherCharReg := regexp.MustCompile("[^\u4e00-\u9fa5a-zA-Z0-9]")
	text = otherCharReg.ReplaceAllString(text, "")
	return text
}

// AddWord 添加敏感词
func (st *SensitiveTrie) AddWord(sensitiveWord string) {
	// 添加前先过滤一遍
	sensitiveWord = st.FilterSpecialChar(sensitiveWord)

	// 将敏感词转换成utf-8编码后的rune类型(int32)
	tireNode := st.root
	sensitiveChars := []rune(sensitiveWord)
	for _, charInt := range sensitiveChars {
		// 添加敏感词到前缀树中
		tireNode = tireNode.AddChild(charInt)
	}
	tireNode.End = true
	tireNode.Data = sensitiveWord
}

// AddWords 批量添加敏感词
func (st *SensitiveTrie) AddWords(sensitiveWords []string) {
	for _, sensitiveWord := range sensitiveWords {
		st.AddWord(sensitiveWord)
	}
}

// Match 查找替换发现的敏感词
func (st *SensitiveTrie) Match(text string) (sensitiveWords []string, replaceText string) {
	if st.root == nil {
		return nil, text
	}

	// 过滤特殊字符
	filteredText := st.FilterSpecialChar(text)
	sensitiveMap := make(map[string]*struct{}) // 利用map把相同的敏感词去重
	textChars := []rune(filteredText)
	textCharsCopy := make([]rune, len(textChars))
	copy(textCharsCopy, textChars)
	for i, textLen := 0, len(textChars); i < textLen; i++ {
		trieNode := st.root.FindChild(textChars[i])
		if trieNode == nil {
			continue
		}

		// 匹配到了敏感词的前缀，从后一个位置继续
		j := i + 1
		for ; j < textLen && trieNode != nil; j++ {
			if trieNode.End {
				// 完整匹配到了敏感词
				if _, ok := sensitiveMap[trieNode.Data]; !ok {
					sensitiveWords = append(sensitiveWords, trieNode.Data)
				}
				sensitiveMap[trieNode.Data] = nil

				// 将匹配的文本的敏感词替换成 *
				st.replaceRune(textCharsCopy, i, j)
			}
			trieNode = trieNode.FindChild(textChars[j])
		}

		// 文本尾部命中敏感词情况
		if j == textLen && trieNode != nil && trieNode.End {
			if _, ok := sensitiveMap[trieNode.Data]; !ok {
				sensitiveWords = append(sensitiveWords, trieNode.Data)
			}
			sensitiveMap[trieNode.Data] = nil
			st.replaceRune(textCharsCopy, i, textLen)
		}
	}

	if len(sensitiveWords) > 0 {
		// 有敏感词
		replaceText = string(textCharsCopy)
	} else {
		// 没有则返回原来的文本
		replaceText = text
	}

	return sensitiveWords, replaceText
}

// replaceRune 字符替换
func (st *SensitiveTrie) replaceRune(chars []rune, begin int, end int) {
	for i := begin; i < end; i++ {
		chars[i] = st.replaceChar
	}
}

// TrieNode 敏感词前缀树节点
type TrieNode struct {
	childMap map[rune]*TrieNode // 本节点下的所有子节点
	Data     string             // 在最后一个节点保存完整的一个内容
	End      bool               // 标识是否最后一个节点
}

// AddChild 前缀树添加字节点
func (tn *TrieNode) AddChild(c rune) *TrieNode {

	if tn.childMap == nil {
		tn.childMap = make(map[rune]*TrieNode)
	}

	if trieNode, ok := tn.childMap[c]; ok {
		// 存在不添加了
		return trieNode
	} else {
		// 不存在
		tn.childMap[c] = &TrieNode{
			childMap: nil,
			End:      false,
		}
		return tn.childMap[c]
	}
}

// FindChild 前缀树查找字节点
func (tn *TrieNode) FindChild(c rune) *TrieNode {
	if tn.childMap == nil {
		return nil
	}

	if trieNode, ok := tn.childMap[c]; ok {
		return trieNode
	}
	return nil
}
