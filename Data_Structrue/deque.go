package Data_Structrue

// 泛型双端队列，T 可以是任意类型
type deque[T any] struct{ l, r []T }

func (q deque[T]) empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q deque[T]) size() int {
	return len(q.l) + len(q.r)
}

func (q *deque[T]) pushFront(v T) {
	q.l = append(q.l, v)
}

func (q *deque[T]) pushBack(v T) {
	q.r = append(q.r, v)
}

func (q *deque[T]) popFront() (v T) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func (q *deque[T]) popBack() (v T) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q deque[T]) front() T {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}

func (q deque[T]) back() T {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}

// 0 <= i < q.size()
func (q deque[T]) get(i int) T {
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}

// insert 在逻辑索引 idx 处插入 val (0 <= idx <= q.size())
func (q *deque[T]) insert(val T, idx int) {
	if idx == 0 {
		q.pushFront(val)
		return
	}
	if idx == q.size() {
		q.pushBack(val)
		return
	}

	if idx <= len(q.l) {
		pos := len(q.l) - idx
		var zero T
		q.l = append(q.l, zero)
		copy(q.l[pos+1:], q.l[pos:len(q.l)-1])
		q.l[pos] = val
	} else {
		ridx := idx - len(q.l)
		if ridx == 0 {
			q.r = append([]T{val}, q.r...)
		} else {
			var zero T
			q.r = append(q.r, zero)
			copy(q.r[ridx+1:], q.r[ridx:len(q.r)-1])
			q.r[ridx] = val
		}
	}
}

// erase 删除逻辑索引 idx 处的元素 (0 <= idx < q.size())
func (q *deque[T]) erase(idx int) {
	if idx < len(q.l) {
		pos := len(q.l) - 1 - idx
		if pos == len(q.l)-1 {
			q.l = q.l[:len(q.l)-1]
		} else {
			copy(q.l[pos:], q.l[pos+1:])
			q.l = q.l[:len(q.l)-1]
		}
	} else {
		ridx := idx - len(q.l)
		if ridx == 0 {
			q.r = q.r[1:]
		} else {
			copy(q.r[ridx:], q.r[ridx+1:])
			q.r = q.r[:len(q.r)-1]
		}
	}
}

// 初始化
// intQueue := deque[int]{}
