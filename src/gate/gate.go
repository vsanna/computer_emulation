package gate

// 入出力がそれぞれ微妙に異なるからなぁ
// TODO: goにマーカーメソッド的なものあるのかな
type Gate interface {
	AsGate() bool
}
