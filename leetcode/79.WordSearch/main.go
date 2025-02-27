package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				rs := dfs(board, i, j, 0, word)

				if rs {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j, pos int, word string) bool {
	// Out of bounds or already visited
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '*' {
		return false
	}

	// Not a match
	if word[pos] != board[i][j] {
		return false
	}

	// Found a full match
	if len(word)-1 == pos {
		return true
	}

	// Found current char, checking the next char of the word
	pos++
	tmp := board[i][j]
	// Mark as visited first
	board[i][j] = '*'

	// Check adjacent directions
	dfsResult :=
		dfs(board, i-1, j, pos, word) ||
			dfs(board, i+1, j, pos, word) ||
			dfs(board, i, j+1, pos, word) ||
			dfs(board, i, j-1, pos, word)
	// Revert value
	board[i][j] = tmp

	return dfsResult
}
