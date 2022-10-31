package builder

import (
	"testing"
)

func TestHtmlBuilder(t *testing.T) {
	expected := `<ul>
  <li>
    golang
  </li>
  <li>
    python
  </li>
</ul>
`
	builder := NewHtmlBuilder("ul")
	builder.AddChildTag("li", "golang")
	builder.AddChildTag("li", "python")

	if expected != builder.String() {
		t.Errorf("TestHtmlBuilder error: expected=%s. but got%s", expected, builder.String())
	}
}
