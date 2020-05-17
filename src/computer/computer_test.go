package computer

import (
	"testing"
)

func TestComputer(t *testing.T) {
	computer := NewComputer()
	computer.Run()
}
