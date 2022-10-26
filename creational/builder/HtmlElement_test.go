package builder

import (
	"testing"
)

func TestHtmlElement(t *testing.T) {
	testcases := []struct {
		html     HtmlElement
		expected string
	}{
		{
			HtmlElement{
				tagName: "p",
				content: "sentence",
			},
			`<p>
  sentence
</p>
`,
		},
		{
			HtmlElement{
				tagName: "ul",
				children: []HtmlElement{
					{
						tagName: "li",
						content: "golang",
					},
					{
						tagName: "li",
						content: "python",
					},
				},
			},
			`<ul>
  <li>
    golang
  </li>
  <li>
    python
  </li>
</ul>
`,
		},
	}

	for _, tt := range testcases {
		if output := tt.html.String(); tt.expected != output {
			t.Errorf("failed to decode html element. expected=%s, but got=%s", tt.expected, output)
		}
	}
}
