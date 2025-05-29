package machine

import (
	"strings"

	"github.com/craftedbyjosh/go-engigma/pkg/plugboard"
	"github.com/craftedbyjosh/go-engigma/pkg/reflector"
	"github.com/craftedbyjosh/go-engigma/pkg/rotor"
	"github.com/craftedbyjosh/go-engigma/pkg/types"
)

type Machine struct {
	rotorPos1  rotor.IRotor
	rotorPos2 rotor.IRotor
	rotorPos3   rotor.IRotor
	reflector   reflector.IReflector
	plugboard plugboard.Plugboard
}

type MachineRotorArgs struct {
	TypeName string
	InitialPos types.Position
}

func NewMachine(rotor1Args, rotor2Args, rotor3Args MachineRotorArgs, reflectorType, plugboardSettings string) *Machine {
	return &Machine{
		rotorPos1: rotor.NewRotor(rotor1Args.TypeName, rotor1Args.InitialPos),
		rotorPos2: rotor.NewRotor(rotor2Args.TypeName, rotor2Args.InitialPos),
		rotorPos3: rotor.NewRotor(rotor3Args.TypeName, rotor3Args.InitialPos),
		reflector: reflector.NewReflectorUKWB(),
		plugboard: plugboard.NewPlugboard(plugboardSettings),
	}
}

func (m *Machine) Run(input string) string {
	inputSlice := types.StringToPositionSlice(input)
	encryptedChars := make([]string, 0)

	for _, pos := range inputSlice {
		// TODO: Nest these, have each object aware of its neighbours so we can just call the one and have it do the legwork
		// 		 That would be much cleaner than this
		char := m.plugboard.TranslatePos(pos)
		char, turnNext := m.rotorPos1.EncodeCharForward(types.Position(char), true)
		char, turnNext = m.rotorPos2.EncodeCharForward(char, turnNext)
		char, _ = m.rotorPos3.EncodeCharForward(char, turnNext)
		char = m.reflector.Reflect(char)
		char = m.rotorPos3.EncodeCharBackward(char)
		char = m.rotorPos2.EncodeCharBackward(char)
		char = m.rotorPos1.EncodeCharBackward(char)
		char = m.plugboard.TranslatePos(char)
		encryptedChars = append(encryptedChars, char.ToString())
	}

	return string(strings.Join(encryptedChars, ""))
}