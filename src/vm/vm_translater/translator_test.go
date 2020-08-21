package vm_translater

import (
	"computer_emulation/src/vm/vm_parser"
	"computer_emulation/src/vm/vm_tokenizer"
	"strings"
	"testing"
)

func TestTranslator_translatePushStatement(t1 *testing.T) {
	input := `
	push constant 2
`
	t2 := vm_tokenizer.New(input)
	p := vm_parser.New(t2)
	program := p.ParseProgram()

	tests := []struct {
		expectedLine string
	}{
		{expectedLine: "0000000000010000"}, // 最初の変数はアドレス16にマッピングされる
		{expectedLine: "1110111111001000"}, // 111 0 111111 001 000
		{expectedLine: "0000000000010001"}, // ２つ目の変数はアドレス17にマッピングされる
		{expectedLine: "1110101010001000"}, // 111 0 101010 001 000
		{expectedLine: "1110001100000111"}, // 111 0 001100 000 111
		{expectedLine: "1111000000111101"}, // 111 1 000000 111 101
		// 1110011000000111
	}

	translator := New(program)
	translatedLines := translator.Translate()

	for i, line := range strings.Split(translatedLines, "\n") {
		if line != tests[i].expectedLine {
			t1.Fatalf("invalid translation. expected=%q, actual=%q", tests[i].expectedLine, line)
		}
	}
}
