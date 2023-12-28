package linkedlist

type LinkedListNode[T any] struct {
	Val  T
	next *LinkedListNode[T]
}

func Create[T any](val T) *LinkedListNode[T] {
	root := LinkedListNode[T]{
		Val:  val,
		next: nil,
	}
	return &root
}

func (l *LinkedListNode[T]) GetLastNode() *LinkedListNode[T] {
	if l.next == nil {
		return l
	}
	return l.next.GetLastNode()
}

func (l *LinkedListNode[T]) Push(val T) {
	lastNode := l.GetLastNode()
	lastNode.next = &LinkedListNode[T]{
		Val:  val,
		next: nil,
	}
}

func (l *LinkedListNode[T]) Pop() T {
	if l.next == nil {
		aux := l.Val
		l = nil
		return aux
	}
	return l.next.Pop()
}
