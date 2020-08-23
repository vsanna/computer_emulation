package bit

import (
	"fmt"
)

/************************
* Bit
**************************/
type Bit struct {
	val int
}

func NewBit(n int) *Bit {
	if n > 1 || n < 0 {
		panic("invalid int")
	}
	return &Bit{val: n}
}

func (bit *Bit) GetVal() int {
	return bit.val
}

func (bit *Bit) String() string {
	return fmt.Sprintf("%d", bit.val)
}

// TODO: これであってる？定数にしたいんだけど。
var ON *Bit = NewBit(1)
var OFF *Bit = NewBit(0)

func ToBit(n int) *Bit {
	if n > 1 || n < 0 {
		panic("invalid int")
	}
	if n == 0 {
		return OFF
	} else {
		return ON
	}
}

/************************
* Bus
**************************/
// TODO: consider better name
const BUS_WIDTH = 16

// TODO: how to constraint the width of Bits
type Bus struct {
	Bits []*Bit
}

type BusOption struct {
	Bits []int
}

func zeroBus() *Bus {
	bits := make([]*Bit, BUS_WIDTH)
	for i := 0; i < BUS_WIDTH; i++ {
		bits[i] = OFF
	}
	return &Bus{Bits: bits}
}

func NewBus(option BusOption) *Bus {
	bus := zeroBus()
	if len(option.Bits) == 0 {
		return bus
	}

	for idx, n := range option.Bits {
		bus.Bits[idx] = ToBit(n)
	}
	return bus
}

func (bus *Bus) String() string {
	output := ""
	for _, bit := range bus.Bits {
		output += bit.String()
	}
	return output
}

func (bus *Bus) Equals(other *Bus) bool {
	for idx, _ := range bus.Bits {
		if bus.Bits[idx].GetVal() != other.Bits[idx].GetVal() {
			return false
		}
	}
	return true
}

func (bus *Bus) ToInt() int {
	n := 0
	for idx, b := range bus.Bits {
		n = n + b.GetVal()
		if idx != len(bus.Bits)-1 {
			n = n << 1
		}
	}

	return n
}

func IntToBus(i int) *Bus {
	if i >= 1024*128 {
		panic("too large")
	}
	result := NewBus(BusOption{Bits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}})

	pos := 0
	for i > 0 {
		lsb := i & 1
		result.Bits[BUS_WIDTH-1-pos] = ToBit(lsb)
		i = i >> 1
		pos += 1
	}

	return result
}
