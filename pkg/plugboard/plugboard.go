package plugboard

import (
	"log/slog"
	"strings"

	"github.com/craftedbyjosh/go-engigma/pkg/types"
)

type Plugboard struct {
	keymap map[types.Position]types.Position
}

// NewPlugboard creates a new Plugboard instance with the provided keymap.
// The keymap should be a string of pairs where each pair represents a connection
// e.g. "AB CD" means A<-->B and C<-->D
func NewPlugboard(keymap string) Plugboard {
	if keymap == "" {
		return Plugboard{
			keymap: make(map[types.Position]types.Position),
		}
	}

	pairs := strings.Split(keymap, " ")
	plugboardMapping := make(map[types.Position]types.Position, len(pairs) * 2)

	for _, pair := range pairs {
		if len(pair) != 2 {
			panic("Plugboard keymap pairs must be exactly two characters long")
		}
	}

	if len(pairs) > 10 {
		panic("Plugboard keymap cannot contain more than 10 pairs")
	}

	for _, pair := range pairs {
		positions := types.StringToPositionSlice(pair)
		plugboardMapping[types.Position(positions[0])] = types.Position(positions[1])
		plugboardMapping[types.Position(positions[1])] = types.Position(positions[0])
	}

	slog.Info("Plugboard created with mapping", "mapping", plugboardMapping)

	return Plugboard{
		keymap: plugboardMapping,
	}
}

func (p *Plugboard) TranslatePos(char types.Position) types.Position {
	if translated, ok := p.keymap[char]; ok {
		return translated
	}
	return char
}

func (p *Plugboard) NumOfPairs() int {
	return len(p.keymap) / 2
}
