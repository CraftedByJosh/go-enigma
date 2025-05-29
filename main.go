package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/craftedbyjosh/go-engigma/pkg/machine"
	"github.com/craftedbyjosh/go-engigma/pkg/types"
)

func main() {
	machine := machine.NewMachine(
		machine.MachineRotorArgs{TypeName: "I", InitialPos: types.A},
		machine.MachineRotorArgs{TypeName: "I", InitialPos: types.A},
		machine.MachineRotorArgs{TypeName: "I", InitialPos: types.A},
		"UKWB",
		"AB CD EF",
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, " \n")

	encryptedText := machine.Run(text)
	fmt.Println("Encrypted text:", encryptedText)
}
