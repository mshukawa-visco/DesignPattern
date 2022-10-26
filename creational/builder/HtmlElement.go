package builder

import (
	"fmt"
	"strings"
)

const indentSize = 2

type HtmlElement struct {
	tagName  string
	content  string
	children []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}

	// インデントを作成
	indentStr := strings.Repeat(" ", indentSize*indent)

	// 開始タグの作成
	sb.WriteString(fmt.Sprintf("%s<%s>\n", indentStr, e.tagName))

	// コンテンツを入れ込む
	if len(e.content) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.content)
		sb.WriteString("\n")
	}

	// 子要素を作成する
	for _, child := range e.children {
		sb.WriteString(child.string(indent + 1))
	}

	// 閉じタグの作成
	sb.WriteString(fmt.Sprintf("%s</%s>\n", indentStr, e.tagName))
	return sb.String()
}
