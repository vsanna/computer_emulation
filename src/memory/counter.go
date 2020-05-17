package memory

import (
	"computer_emulation/src/adder"
	. "computer_emulation/src/bit"
)

// 1word相当のデータのカウント処理を行う論理ゲート.
// PCは都度ここに問い合わせして新しい値を取得する
type Counter struct {
	word        *Word
	load        *RegisterCell
	inc         *RegisterCell
	reset       *RegisterCell
	incrementer *adder.Incrementer
}

func NewCounter() *Counter {
	return &Counter{
		word:        NewWord(),
		load:        NewRegisterCell(),
		reset:       NewRegisterCell(),
		inc:         NewRegisterCell(),
		incrementer: adder.NewIncrementer(),
	}
}

// load:1で新しい値に上書き(writeモード)
func (counter *Counter) Pass(in *Bus, load *Bit, reset *Bit, inc *Bit) *Bus {
	if reset == ON {
		// 0で更新
		return counter.word.Pass(NewBus(BusOption{}), reset)
	}

	if load == ON {
		// inで更新
		return counter.word.Pass(in, load)
	}

	if inc == ON {
		return counter.word.Pass(
			counter.incrementer.Pass(counter.word.GetPrevious()),
			inc,
		)
	}

	return counter.word.Pass(nil, OFF)
}
