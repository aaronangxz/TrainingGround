package main

func countHillValley(nums []int) int {
	temp := []int{}
	for _, x := range nums {
		if len(temp) > 0 && temp[len(temp)-1] == x {
			continue
		}
		temp = append(temp, x)
	}
	cnt := 0
	for i := 1; i < len(temp)-1; i++ {
		if (temp[i] > temp[i-1] && temp[i] > temp[i+1]) ||
			(temp[i] < temp[i-1] && temp[i] < temp[i+1]) {
			cnt++
		}
	}
	return cnt
}
