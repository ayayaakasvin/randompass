package randompass

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateRandomPassword(t *testing.T) {
	tests := []struct {
		name         string
		latinUpperUse bool
		latinLowerUse bool
		digitsUse     bool
		specialSetUse bool
		length        int
		expectLength  int
	}{
		{
			name:         "All symbol types, length 12",
			latinUpperUse: true,
			latinLowerUse: true,
			digitsUse:     true,
			specialSetUse: true,
			length:        12,
			expectLength:  12,
		},
		{
			name:         "Only lowercase letters, length 8",
			latinUpperUse: false,
			latinLowerUse: true,
			digitsUse:     false,
			specialSetUse: false,
			length:        8,
			expectLength:  8,
		},
		{
			name:         "Only digits, length 10",
			latinUpperUse: false,
			latinLowerUse: false,
			digitsUse:     true,
			specialSetUse: false,
			length:        10,
			expectLength:  10,
		},
		{
			name:         "Only special symbols, length 6",
			latinUpperUse: false,
			latinLowerUse: false,
			digitsUse:     false,
			specialSetUse: true,
			length:        6,
			expectLength:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := CreateRandomPassword(tt.latinUpperUse, tt.latinLowerUse, tt.digitsUse, tt.specialSetUse, tt.length)
			assert.Equal(t, tt.expectLength, len(pw.PasswordItself), "Password length should match")
			assert.True(t, len(pw.PasswordItself) >= tt.length, "Password length should be at least the specified length")
		})
	}
}

func TestRandomCharacter(t *testing.T) {
	// Test if RandomCharacter returns a valid character from the collection
	tests := []struct {
		name     string
		symbol   SymbolType
	}{
		{
			name: "UpperLatin characters",
			symbol: SymbolType{
				Collection: []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'},
				TypeSliceLength: LatinLetterSliceSize,
			},
		},
		{
			name: "Digits",
			symbol: SymbolType{
				Collection: []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
				TypeSliceLength: DigitsSliceSize,
			},
		},
		{
			name: "Special characters",
			symbol: SymbolType{
				Collection: []byte{'@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '!', '~'},
				TypeSliceLength: SpecialSetSliceSize,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			char := tt.symbol.RandomCharacter()
			assert.Contains(t, tt.symbol.Collection, char, "Character should be in the collection")
		})
	}
}
