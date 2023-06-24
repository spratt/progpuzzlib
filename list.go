package progpuzzlib

// ListNode represents a singly-linked list in which each node
// contains a single generic value
type ListNode struct {
  Value interface{}
  Next *ListNode
}

// NewListFromSlice takes a []interface{} and returns a linked
// list containing the same values
func NewListFromSlice(input []interface{}) *ListNode {
	var head *ListNode
	prevNode := head
	for i := 0; i < len(input); i++ {
		newNode := make(ListNode)
		newNode.Value = input[i]
		if head == nil {
			head = newNode
		} else {
			prevNode.Next = newNode
			prevNode = prevNode.Next
		}			
	}
	return head
}

// Traverse simply follows the nodes from one end of the list
// to the other
func Traverse(head *ListNode, fn func(v, acc interface{}) (interface{}, bool)) *ListNode {
	var acc interface{}
	var prevNode *ListNode
	node := head
	for {
		if node == nil {
			return
		}
		var del bool
		acc, del = fn(node.Value, acc)
		if del {
			if node == head {
				head = head.Next
			} else {
				prevNode.Next = node.Next
			}
		} else {
			prevNode = node
		}
		node = node.Next
	}
	return head
}
