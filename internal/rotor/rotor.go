package rotor

import (
	"fmt"
	"log/slog"
	"os"
	// "github.com/BurntSushi/toml"
)

type Rotor struct {
	Name     string `toml:"rotor_name"`
	Notch    string `toml:"notch"`
	Turnover string `toml:"turnover"`

	Keymap struct {
		A string `toml:"a"`
		B string `toml:"b"`
		C string `toml:"c"`
		D string `toml:"d"`
		E string `toml:"e"`
		F string `toml:"f"`
		G string `toml:"g"`
		H string `toml:"h"`
		I string `toml:"i"`
		J string `toml:"j"`
		K string `toml:"k"`
		L string `toml:"l"`
		M string `toml:"m"`
		N string `toml:"n"`
		O string `toml:"o"`
		P string `toml:"p"`
		Q string `toml:"q"`
		R string `toml:"r"`
		S string `toml:"s"`
		T string `toml:"t"`
		U string `toml:"u"`
		V string `toml:"v"`
		W string `toml:"w"`
		X string `toml:"x"`
		Y string `toml:"y"`
		Z string `toml:"z"`
	}
}

func NewRotor(rotor RotorType) *Rotor {
	path, err := os.Executable()
	if err != nil {
		slog.Error("unable to get executable path")
	}

	rotorFile := fmt.Sprintf("../../configs/rotors/%s.go", rotor)
	rotorBytes, err := os.ReadFile(rotorFile)
	if err != nil {
		slog.Error("unable to read config for rotor", "rotor.file", rotorFile, "error", err)
	}

	slog.Info("printing loaded rotor file", "rotor.file.contents", string(rotorBytes))

	return nil
}
