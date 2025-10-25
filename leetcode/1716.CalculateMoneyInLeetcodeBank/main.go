package main

func totalMoney(n int) int {
	w := n / 7
	r := n % 7
	fullWeeksSum := w*28 + 7*(w*(w-1)/2)
	remSum := r*(1+w) + (r * (r - 1) / 2)
	return fullWeeksSum + remSum
}
