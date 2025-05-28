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

func (r *Rotor) EncodeCharForward(char types.Position, incPosition bool) (types.Position, bool) {
    if incPosition {
        r.position = types.WrapPosition(r.position + 1)
    }

    enter := types.WrapPosition(char + r.position)

    wired := r.keymap[enter]

    exited := types.WrapPosition(wired - r.position)

    trigger := (r.position == r.turnoverNotch)
    return exited, trigger
}

func (r *Rotor) EncodeCharBackward(char types.Position) types.Position {
    entered := types.WrapPosition(char + r.position)

    var wired types.Position
    for in, out := range r.keymap {
        if out == entered {
            wired = in
            break
        }
    }

    return types.WrapPosition(wired - r.position)
}
