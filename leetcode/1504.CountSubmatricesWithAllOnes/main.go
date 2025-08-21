package main

func numSubmat(mat [][]int) int {
	r, c := len(mat), len(mat[0])
	h := make([]int, c)
	ans := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] == 0 {
				h[j] = 0
			} else {
				h[j]++
			}
		}
		sum := make([]int, c)
		st := []int{}
		for j := 0; j < c; j++ {
			for len(st) > 0 && h[st[len(st)-1]] >= h[j] {
				st = st[:len(st)-1]
			}
			if len(st) > 0 {
				p := st[len(st)-1]
				sum[j] = sum[p] + h[j]*(j-p)
			} else {
				sum[j] = h[j] * (j + 1)
			}
			st = append(st, j)
			ans += sum[j]
		}
	}
	return ans
}
