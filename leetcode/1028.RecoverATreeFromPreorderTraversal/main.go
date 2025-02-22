package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"strconv"
)

// Summary:
// Iterate through the string and perform operation whenever one value is met (count dashes along the way)
// The stack maintains the depth we are currently at, and if the dashes(expected depth) exceeds,
// we need to pop off the stack to get the correct depth (means the node that got popped off has no more nodes linked)
// Once we get to the correct depth, start connecting again
// Because in the string there will be no node that has 0 depth, so root is safely kept in the bottom of stack, we can directly return that
func recoverFromPreorder(traversal string) *common.TreeNode {
	dashes := 0
	st := common.Stack[*common.TreeNode]{}
	i := 0

	for i < len(traversal) {
		// Count dashes (depth)
		if string(traversal[i]) == "-" {
			dashes++
			i++
		} else {
			// From the current index, find where the number stops (before another '-')
			j := i
			for j < len(traversal) && string(traversal[j]) != "-" {
				j++
			}
			// Extract the string
			val := traversal[i:j]
			// Create the new node
			n := new(common.TreeNode)
			valInt, _ := strconv.Atoi(string(val))
			n.Val = valInt

			// As long as there are more nodes than the depth, pop it
			for st.Len() > dashes {
				st.Pop()
			}

			// If stack is not empty and the left node is still nil
			// (recall should fill left node first)
			if st.Len() != 0 && st.Top().Left == nil {
				// Add to left node
				st.Top().Left = n
			} else if st.Len() != 0 {
				// Unless left is occupied, add to right node
				st.Top().Right = n
			}
			// Add new node into stack
			st.Push(n)

			// Continue to the next number and restart dashes count
			i = j
			dashes = 0
		}
	}
	// The root
	return st.Get(0)
}
