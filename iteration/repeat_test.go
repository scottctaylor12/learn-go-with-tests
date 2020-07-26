package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expect := "aaaaa"
	if repeated != expect {
		t.Errorf("got %q but got %q", repeated, expect)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	//Output: aaaaa
}
