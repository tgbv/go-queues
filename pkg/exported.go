package pkg

/*
	Go queue implemented from scratch (LIFO)
	Uses doubly linked lists to maintain good time complexity

	Next/Prev are in the order of queue processing aka:
		Next - head to tail
		Prev - tail to head

*/

/*
*	a node from stack
 */
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

/*
*	the queue
 */
type Queue struct {
	Length uint
	Head   *Node
	Tail   *Node
}

/*
*	Add node to queue
	O(1)
*/
func (q *Queue) Add(v interface{}) *Queue {
	// make node
	n := &Node{Value: v}

	// check length
	if q.Length == 0 {
		q.Head = n
		q.Tail = n
	} else

	// assign prev of tail to node
	// assign node next to current tail
	// assign tail to node
	{
		q.Tail.Next = n
		n.Prev = q.Tail
		q.Tail = n
	}

	q.Length++

	return q
}

/*
*	Traverse queue with a callback
	Traverses from tail to head

	O(n)
*/
func (q *Queue) Traverse(f func(*Node)) *Queue {
	n := q.Tail

	for n != nil {
		f(n)
		n = n.Prev
	}

	return q
}

/*
*	Gets the first element (head) and deletes it from queue
	O(1)
*/
func (q *Queue) Pop() *Node {
	// check queue length
	if q.Length == 0 {
		return nil
	}

	n := q.Head
	q.Head = q.Head.Next

	if q.Length == 1 {
		q.Tail = nil
	}

	q.Length--

	return n
}

/*
*	finds a node from queue based on callback
	traverses from head to tail

	O(n)
*/
func (q *Queue) Search(f func(*Node) bool) *Node {
	n := q.Head

	for n != nil {
		if f(n) == true {
			return n
		}

		n = n.Next
	}

	return nil
}

/*
*	alias for Search()
 */
func (q *Queue) Find(f func(*Node) bool) *Node {
	return q.Search(f)
}
