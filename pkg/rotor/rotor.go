package rotor

import "github.com/craftedbyjosh/go-engigma/pkg/types"

type IRotor interface {
	EncodeCharForward(types.Position, bool) (types.Position, bool)
	EncodeCharBackward(types.Position) types.Position
}

type Rotor struct {
  name          string
  position      types.Position
  turnoverNotch types.Position
  keymap        map[types.Position]types.Position
}

// wrap normalizes a given Position to stay within 0–25 (inclusive)
// It ensures correct wrapping behavior for both positive and negative input values.
//
// Examples:
//   wrap(27)  => 1   (27 mod 26 = 1)
//   wrap(-1)  => 25  (-1 mod 26 = 25, i.e. Z)
//   wrap(0)   => 0   (A)
//   wrap(26)  => 0   (wraps to A again)
func wrap(p types.Position) types.Position {
    return (p%26 + 26) % 26
}


func (r *Rotor) EncodeCharForward(char types.Position, incPosition bool) (types.Position, bool) {
    if incPosition {
        r.position = wrap(r.position + 1)
    }

    enter := wrap(char + r.position)

    wired := r.keymap[enter]

    exited := wrap(wired - r.position)

    trigger := (r.position == r.turnoverNotch)
    return exited, trigger
}

func (r *Rotor) EncodeCharBackward(char types.Position) types.Position {
    entered := wrap(char + r.position)

    var wired types.Position
    for in, out := range r.keymap {
        if out == entered {
            wired = in
            break
        }
    }

    return wrap(wired - r.position)
}
