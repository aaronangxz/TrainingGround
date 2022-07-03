package common

type TreeNodeWithNext struct {
	Val   int
	Left  *TreeNodeWithNext
	Right *TreeNodeWithNext
	Next  *TreeNodeWithNext
}

func NewTreeNodeWithNext(val int) *TreeNodeWithNext {
	return &TreeNodeWithNext{
		Val:   val,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
}
