package main

import (
	"log/slog"

	"github.com/craftedbyjosh/go-engigma/internal/rotor"
)

func main() {
	slog.Info("Running go-enigma")

	rotor.NewRotor(rotor.RotorI)
}
