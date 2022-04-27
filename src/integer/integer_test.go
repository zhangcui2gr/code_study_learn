package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	a := Add(1, 2)
	b := 3
	if a != b {
		t.Errorf("a: %d, b: %d", a, b)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
