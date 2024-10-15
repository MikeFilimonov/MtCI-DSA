package dynamicprogramming

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
*/

func Rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	result := []int{}

	result = append(result, nums[0])
	result = append(result, maxNum(nums[0], nums[1]))

	for i := 2; i < len(nums); i++ {
		result = append(result, maxNum(result[i-2]+nums[i], result[i-1]))
	}

	return result[len(nums)-1]

}

func maxNum(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/
var cache = make(map[int]int)

func climbStairs(n int) int {

	if val, ok := cache[n]; ok {
		return val
	}

	if n < 3 {
		return n
	}

	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]

}

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/
func TimeToBuyStonks(prices []int) int {

	profit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {

		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			profit = prices[i] - minPrice
		}

	}

	return profit

}

func BottomTopFibo(n int) int {

	result := []int{0, 1}
	for i := 2; i <= n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}

	return result[n]

}

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
