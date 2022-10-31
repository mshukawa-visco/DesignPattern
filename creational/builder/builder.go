package builder

type HtmlBuilder struct {
	rootTagName string
	root        HtmlElement
}

func NewHtmlBuilder(rootTagName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootTagName: rootTagName,
		root: HtmlElement{
			rootTagName,
			"",
			[]HtmlElement{},
		},
	}
}

func (hb *HtmlBuilder) String() string {
	return hb.root.String()
}

func (hb *HtmlBuilder) AddChildTag(tagName string, content string) {
	hb.root.children = append(hb.root.children, HtmlElement{tagName, content, []HtmlElement{}})
}
