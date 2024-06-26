package word_combinations

import (
	"maps"
	"strings"
	"testing"
)

func TestAllOrders(t *testing.T) {
	t.Run("Zero elements", func(t *testing.T) {
		got := AllOrders([]string{})
		if len(got) > 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("One element", func(t *testing.T) {
		got := AllOrders([]string{"a"})
		want := [][]string{{"a"}}
		checkArrayEqualsIgnoreOrder(t, got, want)
	})

	t.Run("Two elements", func(t *testing.T) {
		got := AllOrders([]string{"a", "b"})
		want := [][]string{{"a", "b"}, {"b", "a"}}
		checkArrayEqualsIgnoreOrder(t, got, want)
	})

	t.Run("Three elements", func(t *testing.T) {
		got := AllOrders([]string{"a", "b", "c"})
		want := [][]string{
			{"a", "b", "c"}, {"a", "c", "b"}, {"b", "a", "c"}, {"b", "c", "a"}, {"c", "a", "b"}, {"c", "b", "a"},
		}
		checkArrayEqualsIgnoreOrder(t, got, want)
	})
}

func checkArrayEqualsIgnoreOrder(t testing.TB, got, want [][]string) {
	t.Helper()

	if !maps.Equal(asSet(got), asSet(want)) {
		t.Errorf("got %q want %q", got, want)
	}
}

func asSet(arrays [][]string) map[string]bool {
	result := make(map[string]bool)
	for _, words := range arrays {
		s := strings.Join(words, "____$$$____")
		result[s] = true
	}
	return result
}
