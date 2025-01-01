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

func NewSliceToLinkedList(slice []int) *TreeNodeWithNext {
	// Handle the case of an empty slice
	if len(slice) == 0 {
		return nil
	}

	// Create the head of the linked list
	head := &TreeNodeWithNext{Val: slice[0]}
	current := head

	// Iterate through the slice to create and link nodes
	for _, val := range slice[1:] {
		newNode := &TreeNodeWithNext{Val: val}
		current.Next = newNode
		current = newNode
	}

	return head
}

func NewSliceToCyclicLinkedList(slice []int, pos int) *TreeNodeWithNext {
	if len(slice) == 0 {
		return nil
	}

	// Create the head node
	head := &TreeNodeWithNext{Val: slice[0]}
	current := head

	// To track the node at the `pos` index for cycle creation
	var cycleNode *TreeNodeWithNext
	if pos == 0 {
		cycleNode = head
	}

	// Iterate through the slice to create the linked list
	for i := 1; i < len(slice); i++ {
		newNode := &TreeNodeWithNext{Val: slice[i]}
		current.Next = newNode
		current = newNode

		// Set the cycle node when reaching `pos`
		if i == pos {
			cycleNode = newNode
		}
	}

	// Link the last node to the `cycleNode` to form the cycle
	if cycleNode != nil {
		current.Next = cycleNode
	}

	return head
}
