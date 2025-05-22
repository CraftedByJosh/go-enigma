package reflector

import "github.com/craftedbyjosh/go-engigma/internal/rotor"

type IReflector interface {
	Reflect(rotor.Position) rotor.Position
}

type Reflector struct {
	name          string
	keymap        map[rotor.Position]rotor.Position
}

type RelectorUKWB struct {
	Reflector
}

func (r *Reflector) Reflect(char rotor.Position) rotor.Position {
	return r.keymap[char]
}

func NewReflectorUKWB() IReflector {
	return &RelectorUKWB{
		Reflector: Reflector{
			name:          "RotorI",
			keymap: map[rotor.Position]rotor.Position{
				rotor.A: rotor.Y,
				rotor.B: rotor.R,
				rotor.C: rotor.U,
				rotor.D: rotor.H,
				rotor.E: rotor.Q,
				rotor.F: rotor.S,
				rotor.G: rotor.L,
				rotor.H: rotor.D,
				rotor.I: rotor.P,
				rotor.J: rotor.X,
				rotor.K: rotor.N,
				rotor.L: rotor.G,
				rotor.M: rotor.O,
				rotor.N: rotor.K,
				rotor.O: rotor.M,
				rotor.P: rotor.I,
				rotor.Q: rotor.E,
				rotor.R: rotor.B,
				rotor.S: rotor.F,
				rotor.T: rotor.Z,
				rotor.U: rotor.C,
				rotor.V: rotor.W,
				rotor.W: rotor.V,
				rotor.X: rotor.J,
				rotor.Y: rotor.A,
				rotor.Z: rotor.T,
			},
		},
	}
}
