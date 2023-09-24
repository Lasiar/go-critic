package checkers

import (
	"fmt"
	"testing"
	"unicode"
)

func isASCIIRange(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func TestIsASCIIRange(t *testing.T) {
	fmt.Println(isASCIIRange("Привет"))
}
