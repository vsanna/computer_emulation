package translator

import (
	parser "computer_emulation/src/assembler/parser"
	"computer_emulation/src/assembler/tokenizer"
	"strings"
	"testing"
)

func TestTranslator_Translate(t1 *testing.T) {
	input := `
	@i
	M=1;
	@sum
	M=0;
	D;JMP
	ADM=D&M;JNE
`
	t2 := tokenizer.New(input)
	p := parser.New(t2)
	program := p.ParseProgram()

	tests := []struct {
		expectedLine string
	}{
		{expectedLine: "0000010000000000"}, // 最初の変数はアドレス1024に保存される
		{expectedLine: "1110111111001000"}, // 111 0 111111 001 000
		{expectedLine: "0000010000000001"}, // 最初の変数はアドレス1024に保存される
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

func Test_intSliceToString(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test", args{bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, "0000000000000000"},
		{"test", args{bits: []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0}}, "0101010101010100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intSliceToString(tt.args.bits); got != tt.want {
				t.Errorf("intSliceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToBinaryString(t *testing.T) {
	type args struct {
		address int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test", args{address: 0}, "0000000000000000"},
		{"test", args{address: 5}, "0000000000000101"},
		{"test", args{address: 31}, "0000000000011111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToBinaryString(tt.args.address); got != tt.want {
				t.Errorf("intToBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
