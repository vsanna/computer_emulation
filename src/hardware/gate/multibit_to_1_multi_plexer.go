package gate

import . "computer_emulation/src/hardware/bit"

// 1bus -> 1bit
type MultibitTo1MultiPlexer struct {
	multi_plexer *MultiPlexer
	or           *Or
	and          *And
	not          *Not
}

func NewMultibitTo1MultiPlexer() *MultibitTo1MultiPlexer {
	return &MultibitTo1MultiPlexer{multi_plexer: NewMultiPlexer(), or: NewOr(), and: NewAnd(), not: NewNot()}
}

// from 1 bus and 4 bits, specify one bit of the input-bus as output.
// ex. 0010 -> a.Bits[2]
// TODO: enjoy improving!
func (gate *MultibitTo1MultiPlexer) Pass(a *Bus, s1 *Bit, s2 *Bit, s3 *Bit, s4 *Bit) (out *Bit) {
	return gate.or.Pass(
		// 1, 1, 1, 1 -> MSB
		gate.and.Pass(s1, gate.and.Pass(s2, gate.and.Pass(s3, gate.and.Pass(s4, a.Bits[0])))),
		gate.or.Pass(
			gate.and.Pass(s1, gate.and.Pass(s2, gate.and.Pass(s3, gate.and.Pass(gate.not.Pass(s4), a.Bits[1])))),
			gate.or.Pass(
				gate.and.Pass(s1, gate.and.Pass(s2, gate.and.Pass(gate.not.Pass(s3), gate.and.Pass(s4, a.Bits[2])))),
				gate.or.Pass(
					gate.and.Pass(s1, gate.and.Pass(s2, gate.and.Pass(gate.not.Pass(s3), gate.and.Pass(gate.not.Pass(s4), a.Bits[3])))),
					gate.or.Pass(
						gate.and.Pass(s1, gate.and.Pass(gate.not.Pass(s2), gate.and.Pass(s3, gate.and.Pass(s4, a.Bits[4])))),
						gate.or.Pass(
							gate.and.Pass(s1, gate.and.Pass(gate.not.Pass(s2), gate.and.Pass(s3, gate.and.Pass(gate.not.Pass(s4), a.Bits[5])))),
							gate.or.Pass(
								gate.and.Pass(s1, gate.and.Pass(gate.not.Pass(s2), gate.and.Pass(gate.not.Pass(s3), gate.and.Pass(s4, a.Bits[6])))),
								gate.or.Pass(
									gate.and.Pass(s1, gate.and.Pass(gate.not.Pass(s2), gate.and.Pass(gate.not.Pass(s3), gate.and.Pass(gate.not.Pass(s4), a.Bits[7])))),
									gate.or.Pass(
										gate.and.Pass(gate.not.Pass(s1), gate.and.Pass(s2, gate.and.Pass(s3, gate.and.Pass(s4, a.Bits[8])))),
										gate.or.Pass(
											gate.and.Pass(gate.not.Pass(s1), gate.and.Pass(s2, gate.and.Pass(s3, gate.and.Pass(gate.not.Pass(s4), a.Bits[9])))),
											gate.or.Pass(
												gate.and.Pass(gate.not.Pass(s1), gate.and.Pass(s2, gate.and.Pass(gate.not.Pass(s3), gate.and.Pass(s4, a.Bits[10])))),
												gate.or.Pass(
													gate.and.Pass(gate.not.Pass(s1), gate.and.Pass(s2, gate.and.Pass(gate.not.Pass(s3), gate.and.Pass(gate.not.Pass(s4), a.Bits[11])))),
													gate.or.Pass(
														gate.and.Pass(gate.not.Pass(s1), gate.and.Pass(gate.not.Pass(s2), gate.and.Pass(s3, gate.and.Pass(s4, a.Bits[12])))),
														gate.or.Pass(
															gate.and.Pass(gate.not.Pass(s1), gate.and.Pass(gate.not.Pass(s2), gate.and.Pass(s3, gate.and.Pass(gate.not.Pass(s4), a.Bits[13])))),
															gate.or.Pass(
																gate.and.Pass(gate.not.Pass(s1), gate.and.Pass(gate.not.Pass(s2), gate.and.Pass(gate.not.Pass(s3), gate.and.Pass(s4, a.Bits[14])))),
																// 0, 0, 0, 0 -> LSB
																gate.and.Pass(gate.not.Pass(s1), gate.and.Pass(gate.not.Pass(s2), gate.and.Pass(gate.not.Pass(s3), gate.and.Pass(gate.not.Pass(s4), a.Bits[15])))),
															),
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			),
		),
	)
}

func (gate *MultibitTo1MultiPlexer) AsGate() bool { return true }
