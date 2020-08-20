package assembler

import (
	"testing"
)

func TestAssembler_FromString(t *testing.T) {
	type args struct {
		assemblerCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{assemblerCode: "0;"}, "1110101010000000"},
		{"test2", args{assemblerCode: "ADM=M+1;"}, "1111110111111000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Assembler{}
			if got := a.FromString(tt.args.assemblerCode); got != tt.want {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
