package dynamicprogramming

func RealDynamicFibbo(n int) int {

	if n < 2 {
		return n
	}

	cache := make([]int, n+1)

	cache[0] = 0
	cache[1] = 1

	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]

}

type Memoized struct {
	Function func(int) int
	Cache    map[int]int
}

func Memoize(g func(int) int) *Memoized {
	return &Memoized{

		Function: g,
		Cache:    make(map[int]int),
	}
}

func (m *Memoized) DynamicFibonacciNumbers(n int) int {

	if value, ok := m.Cache[n]; ok {
		return value
	}

	result := m.Function(n)
	m.Cache[n] = result
	return result

}

func Fibo(n int) int {

	if n < 2 {
		return n
	}

	return Fibo(n-1) + Fibo(n-2)

}

func theory() {

	/*
		1. Can be divided into subproblems
		2. Has recursive solution
		3. There are repetetive subproblems.
		4. Memoize subproblems
	*/

}
