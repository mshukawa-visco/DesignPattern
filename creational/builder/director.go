package builder

// DirectorがBuilderを手に取ってObjectを組み立てる
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Director) BuildObject() string {
	d.builder.buildStep()
	return d.builder.string()
}
