package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/craftedbyjosh/go-engigma/internal/reflector"
	"github.com/craftedbyjosh/go-engigma/internal/rotor"
)

func main() {
	slog.Info("Running go-enigma", "A.byte", byte('A'))

	rotor1 := rotor.NewRotorI(rotor.A)
	rotor2 := rotor.NewRotorI(rotor.A)
	rotor3 := rotor.NewRotorI(rotor.A)
	reflector := reflector.NewReflectorUKWB()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, " \n")

	encryptedChars := make([]string, 0)

	for i := range len(text) {
		pos := text[i] - 97

		// TODO: Nest these, have each object aware of its neighbours so we can just call the one and have it do the legwork
		// 		 That would be much cleaner than this
		char, turnNext := rotor1.EncodeCharForward(rotor.Position(pos), true) // Encrypt with rotor 1
		char, turnNext = rotor2.EncodeCharForward(char, turnNext)
		char, _ = rotor3.EncodeCharForward(char, turnNext)
		char = reflector.Reflect(char)
		char = rotor3.EncodeCharBackward(char)
		char = rotor2.EncodeCharBackward(char)
		char = rotor1.EncodeCharBackward(char)
		encryptedChars = append(encryptedChars, char.ToString())
	}

	fmt.Println(strings.Join(encryptedChars, ""))
}
