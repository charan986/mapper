package mapper

import "unicode"

// Interface - Generic Transform interface
type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

// SkipString - Holds the String Details
type SkipString struct {
	pos              int
	skipStr          string
	runeChars        []rune
	currentCharCount int
}

// TransformRune - Transforms the current position of alpha char
func (SkipString *SkipString) TransformRune(pos int) {
	ch := SkipString.runeChars[pos]
	if unicode.IsLetter(ch) || unicode.IsNumber(ch) { // Checking only alphanumeric characters, ie [a-z][A-Z][0-9]
		SkipString.currentCharCount++
		if !unicode.IsNumber(ch) {
			if (SkipString.currentCharCount)%SkipString.pos == 0 {
				SkipString.runeChars[pos] = unicode.ToUpper(ch)
			} else {
				SkipString.runeChars[pos] = unicode.ToLower(ch)
			}
		}
	}
}

// Println ...
func (SkipString SkipString) String() string {
	return string(SkipString.runeChars)
}

// GetValueAsRuneSlice ...
func (SkipString *SkipString) GetValueAsRuneSlice() []rune {
	SkipString.runeChars = []rune(SkipString.skipStr)
	return SkipString.runeChars
}

// NewSkipString - Initialize the SkipString structure
func NewSkipString(pos int, skipStr string) SkipString {
	return SkipString{pos: pos, skipStr: skipStr}
}

// MapString - Maps the string as per the give position
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}
