package builder

// 共通のオブジェクト組み立てルーチン
// 必要な項目(変数, 手数)が多ければ多いほどBuilderを使う価値がある。
type Builder interface {
	buildStep()
	string() string
}

func NewBuilder(builderType string) Builder {
	if builderType == "Int" {
		return newIntObjectBuilder()
	}

	if builderType == "String" {
		return newStringObjectBuilder()
	}
	return nil
}
