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
	@SP
	@LCL
	@ARG
	@THIS
	@THAT
	@R0
	@R1
	@R2
	@R3
	@R4
	@R5
	@R6
	@R7
	@R8
	@R9
	@R10
	@R11
	@R12
	@R13
	@R14
	@R15
`
	t2 := tokenizer.New(input)
	p := parser.New(t2)
	program := p.ParseProgram()

	tests := []struct {
		expectedLine string
	}{
		{expectedLine: "0000000000010000"},
		{expectedLine: "1110111111001000"},
		{expectedLine: "0000000000010001"},
		{expectedLine: "1110101010001000"},
		{expectedLine: "1110001100000111"},
		{expectedLine: "1111000000111101"},

		{expectedLine: "0000000000000000"},
		{expectedLine: "0000000000000001"},
		{expectedLine: "0000000000000010"},
		{expectedLine: "0000000000000011"},
		{expectedLine: "0000000000000100"},

		{expectedLine: "0000000000000000"},
		{expectedLine: "0000000000000001"},
		{expectedLine: "0000000000000010"},
		{expectedLine: "0000000000000011"},
		{expectedLine: "0000000000000100"},
		{expectedLine: "0000000000000101"},
		{expectedLine: "0000000000000110"},
		{expectedLine: "0000000000000111"},
		{expectedLine: "0000000000001000"},
		{expectedLine: "0000000000001001"},
		{expectedLine: "0000000000001010"},
		{expectedLine: "0000000000001011"},
		{expectedLine: "0000000000001100"},
		{expectedLine: "0000000000001101"},
		{expectedLine: "0000000000001110"},
		{expectedLine: "0000000000001111"},
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
