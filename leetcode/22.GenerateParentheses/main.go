package main

import "fmt"

/*
suppose n=3

left = 3  > 0
(
left = 2 right = 3  right > left                                            left = 2 right = 3  left > 0
()                                                                          ((
left = 2 right = 2                                                          left = 1 right = 3  right > left                                               left = 1 right = 3  left > 0
()(                                                                         (()                                                                            (((
left = 1 right = 2  right > left      left = 1 right = 2  left > 0          left = 1 right = 2  right > left     left = 1 right = 2  left > 0              left = 0 right = 3  right > left
()()                                  ()((                                  (())                                 (()(                                      ((()
left = 0 right = 1                    left = 0 right = 2  right > left      left = 1 right = 1  left > 0         left = 0 right = 2  right > left          left = 0 right = 2  right > left
()()(                                 ()(()                                 (())(                                (()()                                     ((())
left = 0 right = 1  right > left      left = 0 right = 1  right > left      left = 0 right = 1  right > left     left = 0 right = 1  right > left          left = 0 right = 1  right > left
()()()                                ()(())                                (())()                               (()())                                    ((()))
left = 0 right = 0  return            left = 0 right = 0  return            left = 0 right = 0  return           left = 0 right = 0  return                left = 0 right = 0  return
"Case 1"                                 "Case 2"                            "Case 3"                                  "Case 4"                                  "Case 5"
*/
func gen(res *[]string, s string, left, right int) {
	//left / right = remaining left and right parentheses
	//whenever there is no remaining, push the complete set into res
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}
	if left > 0 {
		gen(res, s+"(", left-1, right)
	}
	//Only add right only if we have the same number of left
	if right > left {
		gen(res, s+")", left, right-1)
	}
}

func generateParenthesis(n int) []string {
	var res []string
	gen(&res, "", n, n)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
