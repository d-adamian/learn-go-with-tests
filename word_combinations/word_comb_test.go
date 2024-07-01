package word_combinations

import (
	"maps"
	"strings"
	"testing"
)

func TestLongestCombinedWord(t *testing.T) {
	line := "the dog cat mouse thedogcat mouseandcat"
	longestWord := FindLongestCombinedWord(line)
	want := "thedogcat"
	if longestWord != want {
		t.Errorf("Expected %q got %q", longestWord, want)
	}
}

func TestAllSubsets(t *testing.T) {
	t.Run("Zero elements", func(t *testing.T) {
		checkEmptyList(t, AllSubsets([]string{}))
	})

	t.Run("One element", func(t *testing.T) {
		got := AllSubsets([]string{"a"})
		want := [][]string{{"a"}}
		checkArrayEqualsIgnoreOrder(t, got, want)
	})

	t.Run("Two elements", func(t *testing.T) {
		got := AllSubsets([]string{"a", "b"})
		want := [][]string{{"a"}, {"b"}, {"a", "b"}}
		checkArrayEqualsIgnoreOrder(t, got, want)
	})

	t.Run("Three elements", func(t *testing.T) {
		got := AllSubsets([]string{"a", "b", "c"})
		want := [][]string{{"a"}, {"b"}, {"c"}, {"a", "b"}, {"a", "c"}, {"b", "c"}, {"a", "b", "c"}}
		checkArrayEqualsIgnoreOrder(t, got, want)
	})
}

func TestAllOrders(t *testing.T) {
	t.Run("Zero elements", func(t *testing.T) {
		got := AllOrders([]string{})
		checkEmptyList(t, got)
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

func checkEmptyList(t testing.TB, got [][]string) {
	t.Helper()

	if len(got) > 0 {
		t.Errorf("Expected empty, got %v", got)
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
