# go-queues
Go queues created from scratch
Uses singly linked lists principle to maintain efficient time complexity

To create a stack: 
```
import s "github.com/tgbv/go-queues/pkg"

v := s.Queue{}

// or

var v s.Queue = s.Queue{}
```

Available methods:
```
func (s *Queue) Add(v interface{}) *Queue

func (s *Queue) Pop() *Node

func (s *Queue) Traverse(f func(*Node)) *Queue

func (s *Queue) Find(f func(*Node) bool) *Node
```
