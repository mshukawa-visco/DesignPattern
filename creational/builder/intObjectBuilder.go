package builder

import (
	"strconv"
)

type IntObjectBuilder struct {
	object IntObject
}

func newIntObjectBuilder() *IntObjectBuilder {
	return &IntObjectBuilder{}
}

func (b *IntObjectBuilder) buildStep() {
	b.object.val = 1
}

func (b *IntObjectBuilder) string() string {
	return strconv.Itoa(b.object.val)
}
