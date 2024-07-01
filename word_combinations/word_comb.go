package word_combinations

import "strings"

func FindLongestCombinedWord(line string) string {
	words := strings.Fields(line)

	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}

	longestWord := ""
	subsets := AllSubsets(words)
	for _, subset := range subsets {
		orders := AllOrders(subset)
		for _, orderedWords := range orders {
			if len(orderedWords) <= 1 {
				continue
			}

			combined := strings.Join(orderedWords, "")
			if len(combined) > len(longestWord) && wordSet[combined] {
				longestWord = combined
			}
		}
	}
	return longestWord
}

func AllSubsets(words []string) [][]string {
	switch {
	case len(words) == 0:
		return nil
	case len(words) <= 1:
		return [][]string{words}
	default:
		var result [][]string
		for i := range words {
			rest := withoutWord(words, i)
			result = append(result, AllSubsets(rest)...)
		}
		result = append(result, words)
		return result
	}
}

func AllOrders(words []string) [][]string {
	switch {
	case len(words) == 0:
		return nil
	case len(words) == 1:
		return [][]string{words}
	default:
		var result [][]string
		for i := range words {
			rest := withoutWord(words, i)
			for _, restOrder := range AllOrders(rest) {
				var order []string
				order = append(order, restOrder...)
				order = append(order, words[i])
				result = append(result, order)
			}
		}
		return result
	}
}

func withoutWord(words []string, idx int) []string {
	switch {
	case idx == 0:
		return words[1:]
	case idx == len(words)-1:
		return words[:idx]
	default:
		var result []string
		result = append(result, words[:idx]...)
		result = append(result, words[idx+1:]...)
		return result
	}
}
