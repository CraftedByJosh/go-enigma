package rotor

type RotorI struct {
	Rotor // Compose this object from base Rotor struct
}

func NewRotorI(initialPos Position) IRotor {
	return &RotorI{
		Rotor: Rotor{
			name:          "RotorI",
			position:      initialPos,
			turnoverNotch: Q,
			keymap: map[Position]Position{
				A: E,
				B: K,
				C: M,
				D: F,
				E: L,
				F: G,
				G: D,
				H: Q,
				I: V,
				J: Z,
				K: N,
				L: T,
				M: O,
				N: W,
				O: Y,
				P: H,
				Q: X,
				R: U,
				S: S,
				T: P,
				U: A,
				V: I,
				W: B,
				X: R,
				Y: C,
				Z: J,
			},
		},
	}
}

type Position int

func (p Position) ToString() string {
	return string((int(p) + 97))
}

const (
	A Position = 0
	B Position = 1
	C Position = 2
	D Position = 3
	E Position = 4
	F Position = 5
	G Position = 6
	H Position = 7
	I Position = 8
	J Position = 9
	K Position = 10
	L Position = 11
	M Position = 12
	N Position = 13
	O Position = 14
	P Position = 15
	Q Position = 16
	R Position = 17
	S Position = 18
	T Position = 19
	U Position = 20
	V Position = 21
	W Position = 22
	X Position = 23
	Y Position = 24
	Z Position = 25
)
