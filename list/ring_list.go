package list

type RingQueue[T any] struct {
	count   int
	head    int
	tail    int
	q       []T
	zeroVal T
	cap     int
}

func NewRingList[T any](cap int) *RingQueue[T] {
	return &RingQueue[T]{
		q:   make([]T, cap),
		cap: cap,
	}
}

func (r *RingQueue[T]) In(val T) error {

	if r.count == r.cap {
		return QueueFull
	}

	r.q[r.tail] = val
	r.tail++
	r.count++
	if r.count >= r.cap {
		r.count = r.cap
	}
	if r.tail == r.cap {
		r.tail = 0
	}

	return nil
}

func (r *RingQueue[T]) Out() (T, error) {
	if r.IsEmpty() {
		return r.zeroVal, QueueEmpty
	}
	front := r.q[r.head]
	r.q[r.head] = r.zeroVal

	r.head++
	r.count--
	if r.head == r.cap {
		r.head = 0
	}
	return front, nil
}

func (r *RingQueue[T]) IsEmpty() bool {
	return r.count == 0
}
