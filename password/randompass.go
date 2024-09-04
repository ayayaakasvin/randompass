package randompass

import (
	"fmt"
	"math/rand"

	"github.com/wagslane/go-password-validator"
)

const (
	latinUpperTitle = "UpperLatin"
	latinLowerTitle = "LowerLatin"
	digitsTitle     = "Digits"
	specialTitle    = "Special"
	LatinLetterSliceSize = 26
	DigitsSliceSize = 10
	SpecialSetSliceSize = 13
)

// PasswordStruct represents a generated password and its metadata.
type PasswordStruct struct {
	PasswordItself   []byte           // The generated password in byte slice format.
	Length           int              // The desired length of the generated password.
	currentMissing   int              // The number of characters left to generate.
	UsedSymbolTypes  map[string]bool // Map indicating which symbol types are used.
	EntropyOfPassword float64        // Entropy of the generated password, calculated for strength estimation.
}

// SymbolType represents a type of symbols that can be used in password generation.
type SymbolType struct {
	Title          string  // The name or title of the symbol type.
	Collection     []byte  // The collection of symbols available for this type.
	TypeSliceLength int     // The length of the symbol collection.
}

// SymbolSet holds different symbol types available for password generation.
type SymbolSet struct {
	UpperLatin SymbolType // Uppercase Latin letters (A-Z).
	LowerLatin SymbolType // Lowercase Latin letters (a-z).
	Digits     SymbolType // Digits (0-9).
	Special    SymbolType // Special characters like '@', '#', '$', etc.
}

// TypeAndCollection is a map from symbol type titles to SymbolType.
type TypeAndCollection map[string]SymbolType

var SymbolsToUse SymbolSet // Global variable holding the symbol sets.
var KeyToCollectionMap TypeAndCollection // Map from symbol type titles to their collections.

func init() {
	SymbolsToUse = SymbolSet{
		UpperLatin: SymbolType{
			Title: latinUpperTitle,
			Collection: func() []byte {
				var result []byte
				for i := 'A'; i <= 'Z'; i++ {
					result = append(result, byte(i))
				}
				return result
			}(),
			TypeSliceLength: LatinLetterSliceSize,
		},

		LowerLatin: SymbolType{
			Title: latinLowerTitle,
			Collection: func() []byte {
				var result []byte
				for i := 'a'; i <= 'z'; i++ {
					result = append(result, byte(i))
				}
				return result
			}(),
			TypeSliceLength: LatinLetterSliceSize,
		},

		Digits: SymbolType{
			Title: digitsTitle,
			Collection: func() []byte {
				var result []byte
				for i := '0'; i <= '9'; i++ {
					result = append(result, byte(i))
				}
				return result
			}(),
			TypeSliceLength: DigitsSliceSize,
		},

		Special: SymbolType{
			Title: specialTitle,
			Collection: []byte{'@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '!', '~'},
			TypeSliceLength: SpecialSetSliceSize,
		},
	}

	KeyToCollectionMap = map[string]SymbolType{
		SymbolsToUse.UpperLatin.Title: SymbolsToUse.UpperLatin,
		SymbolsToUse.LowerLatin.Title: SymbolsToUse.LowerLatin,
		SymbolsToUse.Digits.Title:     SymbolsToUse.Digits,
		SymbolsToUse.Special.Title:    SymbolsToUse.Special,
	}
}

// CheckIfAnyTypeSelected checks if at least one symbol type is selected for password generation.
// Returns true if at least one type is selected; otherwise, returns false.
func (ps *PasswordStruct) CheckIfAnyTypeSelected() bool {
	for _, isSelected := range ps.UsedSymbolTypes {
		if isSelected {
			return true
		}
	}
	return false
}

// RandomCharacter returns a random character from the symbol collection.
func (SymbolStruct SymbolType) RandomCharacter() byte {
	return SymbolStruct.Collection[rand.Intn(SymbolStruct.TypeSliceLength)]
}

// CreateRandomPassword generates a random password based on the selected types and length.
// Returns a PasswordStruct with the generated password. If no symbol types are selected, returns an empty PasswordStruct.
func CreateRandomPassword(LatinUpperUse, LatinLowerUse, DigitsUse, SpecialSetUse bool, Length int) PasswordStruct {
	var ResultPassword PasswordStruct = PasswordStruct{
		PasswordItself:   []byte{},
		Length:           Length,
		currentMissing:   Length,
		UsedSymbolTypes: map[string]bool{
			latinUpperTitle: LatinUpperUse,
			latinLowerTitle: LatinLowerUse,
			digitsTitle:     DigitsUse,
			specialTitle:    SpecialSetUse,
		},
	}

	if !ResultPassword.CheckIfAnyTypeSelected() {
		// Return an empty struct if no types are selected
		return PasswordStruct{}
	}

	ResultPassword.CreateUnShuffledPassword()
	ResultPassword.ShuffleThePassword()
	ResultPassword.EntropyOfPassword = passwordvalidator.GetEntropy(string(ResultPassword.PasswordItself))

	return ResultPassword
}

// ShuffleThePassword randomly shuffles the characters in the password.
func (Password *PasswordStruct) ShuffleThePassword() {
	rand.Shuffle(Password.Length, func(i, j int) {
		Password.PasswordItself[i], Password.PasswordItself[j] = Password.PasswordItself[j], Password.PasswordItself[i]
	})
}

// CreateUnShuffledPassword generates the password without shuffling. It ensures that the desired length
// is met by selecting characters from the available symbol types.
func (Password *PasswordStruct) CreateUnShuffledPassword() {
	for ; Password.currentMissing > 0; {
		for TypeName, isUsed := range Password.UsedSymbolTypes {
			if Password.currentMissing <= 0 {
				break
			}

			if isUsed {
				SymbolToAppend := KeyToCollectionMap[TypeName]
				var Char byte = SymbolToAppend.RandomCharacter()
				Password.PasswordItself = append(Password.PasswordItself, Char)
				Password.currentMissing--
			}
		}
	}
}

func (Password *PasswordStruct) DisplayPassword() {
	if Password == nil || Password.PasswordItself == nil {
		fmt.Println("Password is empty")
	}
	fmt.Println(string(Password.PasswordItself))
}