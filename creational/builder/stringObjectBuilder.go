package builder

type StringObjectBuilder struct {
	object StringObject
}

func newStringObjectBuilder() *StringObjectBuilder {
	return &StringObjectBuilder{}
}

func (b *StringObjectBuilder) buildStep() {
	b.object.val = "a"
}

func (b *StringObjectBuilder) string() string {
	return b.object.val
}
