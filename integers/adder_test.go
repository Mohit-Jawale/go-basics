package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(4, 2)
	excepted := 6

	if sum != excepted {
		t.Errorf("expected '%d' but got '%d'", sum, excepted)
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
