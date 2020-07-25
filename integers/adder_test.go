package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Print(sum)
	// Output: 7
}

func TestSubtract(t *testing.T) {
	diff := Subtract(2, 2)
	expected := 0

	if diff != expected {
		t.Errorf("expected %d got %d", expected, diff)
	}
}

func ExampleSubtract() {
	diff := Subtract(2, 2)
	fmt.Print(diff)
	//Output: 0
}
