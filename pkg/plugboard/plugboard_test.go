package plugboard_test

import (
	"testing"

	"github.com/craftedbyjosh/go-engigma/pkg/plugboard"
	"github.com/craftedbyjosh/go-engigma/pkg/types"
)

func TestNewPlugboard(t *testing.T) {
	tests := []struct {
		name      string
		keymap    string
		expectedNumPairs int
		shouldPanic bool
	}{
		{
			name:      "EmptyKeymap",
			keymap:    "",
			expectedNumPairs: 0,
		},
		{
			name:   "ValidPairs",
			keymap: "AB CD",
			expectedNumPairs: 2,
		},
		{
			name:        "InvalidPairLength",
			keymap:      "ABC DE",
			shouldPanic: true,
		},
		{
			name:        "TooManyPairs",
			keymap:      "AB CD EF GH IJ KL MN OP QR ST UV",
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic for %s", tt.name)
					}
				}()
				plugboard.NewPlugboard(tt.keymap)
			} else {
				pb := plugboard.NewPlugboard(tt.keymap)
				if pb.NumOfPairs() != tt.expectedNumPairs {
					t.Errorf("expected %d pairs, got %d", tt.expectedNumPairs, pb.NumOfPairs())
				}
			}
		})
	}
}

func TestPlugboard_TranslatePos(t *testing.T) {
	pb := plugboard.NewPlugboard("AB CD")

	tests := []struct {
		input    types.Position
		expected types.Position
	}{
		{types.A, types.B},
		{types.B, types.A},
		{types.C, types.D},
		{types.D, types.C},
		{types.E, types.E},
		{types.F, types.F},
	}

	for _, tc := range tests {
		got := pb.TranslatePos(tc.input)
		if got != tc.expected {
			t.Errorf("expected %v to translate to %v, got %v", tc.input, tc.expected, got)
		}
	}
}