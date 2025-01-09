package arrays_generics

func Sum(numbers []int) int {
	add := func(acc, x int) int {
		return acc + x
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[T, U any](collection []T, f func(U, T) U, initialValue U) U {
	var result = initialValue
	for i := range collection {
		result = f(result, collection[i])
	}
	return result
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
