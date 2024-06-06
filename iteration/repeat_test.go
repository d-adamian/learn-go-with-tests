package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaa" + "aaaaa"

	if repeated != expected {
		t.Errorf("Expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 20)
	}
}

func ExampleRepeat() {
	repeated := Repeat("1", 5)
	fmt.Println(repeated)
	// Output: 11111
}
