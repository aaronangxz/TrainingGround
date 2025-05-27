package main

func differenceOfSums(n int, m int) int {
	totalSum := n * (n + 1) / 2
	divisibleCount := n / m
	divisibleSum := m * divisibleCount * (divisibleCount + 1) / 2
	return totalSum - 2*divisibleSum
}
