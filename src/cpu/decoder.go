package cpu

import (
	. "computer_emulation/src/bit"
	"computer_emulation/src/gate"
)

type Decoder struct {
	multibit_multi_plexer *gate.MultibitMultiPlexer
	not                   *gate.Not
}

func NewDecoder() *Decoder {
	return &Decoder{multibit_multi_plexer: gate.NewMultibitMultiPlexer(), not: gate.NewNot()}
}

func (decoder *Decoder) Pass(in *Bus) (
	isCommandA *Bit,
	address *Bus,
	opsBus *Bus,
	destBus *Bus,
	jumpBus *Bus,
) {
	isCommandA = decoder.not.Pass(in.Bits[0])
	address = NewBus(BusOption{
		Bits: []int{
			0,
			in.Bits[1].GetVal(),
			in.Bits[2].GetVal(),
			in.Bits[3].GetVal(),
			in.Bits[4].GetVal(),
			in.Bits[5].GetVal(),
			in.Bits[6].GetVal(),
			in.Bits[7].GetVal(),
			in.Bits[8].GetVal(),
			in.Bits[9].GetVal(),
			in.Bits[10].GetVal(),
			in.Bits[11].GetVal(),
			in.Bits[12].GetVal(),
			in.Bits[13].GetVal(),
			in.Bits[14].GetVal(),
			in.Bits[15].GetVal(),
		},
	})

	opsBus = decoder.multibit_multi_plexer.Pass(
		NewBus(BusOption{}),
		NewBus(BusOption{
			Bits: []int{
				in.Bits[3].GetVal(),
				in.Bits[4].GetVal(),
				in.Bits[5].GetVal(),
				in.Bits[6].GetVal(),
				in.Bits[7].GetVal(),
				in.Bits[8].GetVal(),
				in.Bits[9].GetVal(),
				0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		}),
		isCommandA,
	)

	destBus = decoder.multibit_multi_plexer.Pass(
		NewBus(BusOption{}),
		NewBus(BusOption{
			Bits: []int{
				in.Bits[10].GetVal(),
				in.Bits[11].GetVal(),
				in.Bits[12].GetVal(),
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		}),
		isCommandA,
	)

	jumpBus = decoder.multibit_multi_plexer.Pass(
		NewBus(BusOption{}),
		NewBus(BusOption{
			Bits: []int{
				in.Bits[13].GetVal(),
				in.Bits[14].GetVal(),
				in.Bits[15].GetVal(),
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		}),
		isCommandA,
	)

	return isCommandA, address, opsBus, destBus, jumpBus
}
