package utils

import (
	"fmt"
	"testing"
)

func TestQuotes(t *testing.T) {
	o := ParseLine(`blablabla "   yolo\"" `, ' ')
	wanted := []string{"blablabla", "   yolo\""}
	if !strSliceEquals(o, wanted) {
		t.Fatalf("wanted %v got %v\n", wanted, o)
	}
}

func ExampleParseLineDoubleQuote() {
	var o []string
	o = ParseLine(`"a b c"`, ' ')

	for i := range o {
		fmt.Println(o[i])
	}

	//Output:
	//a b c
}

func ExampleParseLineSingleQuote() {
	var o []string
	o = ParseLine(`'a b c'`, ' ')

	for i := range o {
		fmt.Println(o[i])
	}

	//Output:
	//a b c
}

func ExampleParseLineSpace() {
	o := ParseLine(`a b c`, ' ')

	for i := range o {
		fmt.Println(o[i])
	}

	//Output:
	//a
	//b
	//c
	//
}

func ExampleParseLineEscape() {
	o := ParseLine(`\''a b c\''`, ' ')

	for i := range o {
		fmt.Println(o[i])
	}

	//Output:
	//'a b c'
}

func strSliceEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
