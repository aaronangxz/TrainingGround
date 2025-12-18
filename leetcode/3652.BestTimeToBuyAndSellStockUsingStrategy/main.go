package main

func maxProfit(prices []int, strategy []int, k int) int64 {
	n, h := len(prices), k>>1
	var base, prev, nxt, best int64

	for i := 0; i < n; i++ {
		base += int64(strategy[i] * prices[i])
	}

	for i := 0; i < k; i++ {
		prev += int64(strategy[i] * prices[i])
		if i >= h {
			nxt += int64(prices[i])
		}
	}

	if nxt-prev > best {
		best = nxt - prev
	}

	for r := k; r < n; r++ {
		l := r - k + 1
		prev += int64(strategy[r]*prices[r] -
			strategy[l-1]*prices[l-1])
		nxt += int64(prices[r] - prices[l-1+h])
		if nxt-prev > best {
			best = nxt - prev
		}
	}

	return base + best
}
