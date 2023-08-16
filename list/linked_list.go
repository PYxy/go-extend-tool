package list

var (
	_ List[any] = &LinkedList[any]{}
)

type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

type LinkedList[T any] struct {
	head      *node[T]
	tail      *node[T]
	baseValue T
	length    int
}

// NewLinkedList 创建一个空的双向循环链表
func NewLinkedList[T any]() *LinkedList[T] {
	//head tail  是哨兵 不动的
	head := &node[T]{}
	tail := &node[T]{
		prev: head,
		next: head,
	}
	head.next, head.prev = tail, tail
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

func NewLinkedBySlice[T any](srcSlice []T) (*LinkedList[T], error) {
	tmpLink := NewLinkedList[T]()

	err := tmpLink.Append(srcSlice...)
	if err != nil {
		return nil, err
	}
	tmpLink.length = len(srcSlice)
	return tmpLink, nil
}

func (l *LinkedList[T]) findNode(index int) *node[T] {
	var cur *node[T]
	if index <= l.Len()/2 {
		cur = l.head.next
		for i := 0; i < index; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail.prev
		for i := l.Len() - 1; i > index; i-- {
			cur = cur.prev
		}
	}

	return cur
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	//TODO implement me
	if !l.checkIndex(index) {

		return l.baseValue, IndexOutOfRange
		//errors.New(fmt.Sprintf(IndexOutOfRange, l.length, index))
	}
	tmpNode := l.findNode(index)
	return tmpNode.data, nil
}

func (l *LinkedList[T]) Append(vals ...T) error {
	if l.head == nil || l.tail == nil {
		return QueueIsNotIni
	}
	for _, val := range vals {
		newNode := &node[T]{data: val}
		tail := l.tail
		//先改变自己
		newNode.next, newNode.prev = tail, tail.prev //  newNode.prev = tail.prev
		//在改变外围环境
		//会相互循环
		tail.prev, tail.prev.next = newNode, newNode //tail.prev = newNode

		l.length++
	}

	return nil
}

func (l *LinkedList[T]) Add(index int, val T) error {
	//TODO implement me
	if !l.checkIndex(index) {
		return IndexOutOfRange
	}

	cur := l.findNode(index)
	newNode := &node[T]{data: val}
	newNode.prev, newNode.next = cur.prev, cur
	cur.prev, cur.prev.next = newNode, newNode
	l.length++

	return nil
}

func (l *LinkedList[T]) Set(index int, val T) error {
	//TODO implement me
	if !l.checkIndex(index) {
		return IndexOutOfRange
	}
	cur := l.findNode(index)
	cur.data = val
	return nil
}

func (l *LinkedList[T]) Delete(index int) (T, error) {
	if !l.checkIndex(index) {
		return l.baseValue, IndexOutOfRange
	}
	cur := l.findNode(index)
	cur.prev.next, cur.next.prev = cur.next, cur.prev
	cur.prev, cur.next = nil, nil
	l.length--
	return cur.data, nil
}

func (l *LinkedList[T]) Len() int {
	//TODO implement me
	return l.length
}

func (l *LinkedList[T]) Cap() int {
	//TODO implement me
	return l.Len()
}

func (l *LinkedList[T]) Range(fn func(index int, t T) error) error {
	//TODO implement me
	for cur, i := l.head.next, 0; i < l.length; i++ {
		err := fn(i, cur.data)
		if err != nil {
			return err
		}
		cur = cur.next
	}
	return nil
}

func (l *LinkedList[T]) AsSliceAsc() []T {
	var tmpSlice []T
	tmpSlice = make([]T, 0, l.length)
	//TODO implement me
	cur := l.head.next
	for i := 0; i < l.length; i++ {
		tmpSlice = append(tmpSlice, cur.data)
		cur = cur.next
	}
	return tmpSlice
}

func (l *LinkedList[T]) AsSliceRev() []T {
	var tmpSlice []T
	tmpSlice = make([]T, 0, l.length)
	//TODO implement me
	cur := l.tail.prev
	for i := 0; i < l.length; i++ {
		tmpSlice = append(tmpSlice, cur.data)
		cur = cur.prev
	}
	return tmpSlice
}

func (l *LinkedList[T]) checkIndex(index int) bool {
	return 0 <= index && index < l.Len()
}
