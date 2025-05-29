package rotor

import "github.com/craftedbyjosh/go-engigma/pkg/types"

func NewRotor(typeName string, initialPos types.Position) IRotor {
	switch typeName {
	case "I":
		return newRotorI(initialPos)
	default:
		return nil // or handle error
	}
}

type RotorI struct {
	Rotor // Compose this object from base Rotor struct
}

func newRotorI(initialPos types.Position) IRotor {
	return &RotorI{
		Rotor: Rotor{
			name:          "RotorI",
			position:      initialPos,
			turnoverNotch: types.Q,
			keymap: map[types.Position]types.Position{
				types.A: types.E,
				types.B: types.K,
				types.C: types.M,
				types.D: types.F,
				types.E: types.L,
				types.F: types.G,
				types.G: types.D,
				types.H: types.Q,
				types.I: types.V,
				types.J: types.Z,
				types.K: types.N,
				types.L: types.T,
				types.M: types.O,
				types.N: types.W,
				types.O: types.Y,
				types.P: types.H,
				types.Q: types.X,
				types.R: types.U,
				types.S: types.S,
				types.T: types.P,
				types.U: types.A,
				types.V: types.I,
				types.W: types.B,
				types.X: types.R,
				types.Y: types.C,
				types.Z: types.J,
			},
		},
	}
}
