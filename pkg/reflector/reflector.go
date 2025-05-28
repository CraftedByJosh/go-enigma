package reflector

import (
	"github.com/craftedbyjosh/go-engigma/pkg/types"
)

type IReflector interface {
	Reflect(types.Position) types.Position
}

type Reflector struct {
	name          string
	keymap        map[types.Position]types.Position
}

type RelectorUKWB struct {
	Reflector
}

func (r *Reflector) Reflect(char types.Position) types.Position {
	return r.keymap[char]
}

func NewReflectorUKWB() IReflector {
	return &RelectorUKWB{
		Reflector: Reflector{
			name:          "RotorI",
			keymap: map[types.Position]types.Position{
				types.A: types.Y,
				types.B: types.R,
				types.C: types.U,
				types.D: types.H,
				types.E: types.Q,
				types.F: types.S,
				types.G: types.L,
				types.H: types.D,
				types.I: types.P,
				types.J: types.X,
				types.K: types.N,
				types.L: types.G,
				types.M: types.O,
				types.N: types.K,
				types.O: types.M,
				types.P: types.I,
				types.Q: types.E,
				types.R: types.B,
				types.S: types.F,
				types.T: types.Z,
				types.U: types.C,
				types.V: types.W,
				types.W: types.V,
				types.X: types.J,
				types.Y: types.A,
				types.Z: types.T,
			},
		},
	}
}
