// 二分栈优化 DP
// 用双端队列维护最优决策点，以及决策点的作用域
// 这是一维情况，如果涉及 2 维，需要在结构体中额外定义一个快照
// 二维情况见 CF 321E

package dp

func BinaryStack(n int) int {
	f := make([]int, n + 1)

	q := deque[opt]{}
	q.pushBack(opt{0, 1, n})

	w := func (j, i int) int {
		return j + i  // **根据题目修改**
	}

	better := func(i, j, x int) bool {
		resI := f[j] + w(j, x)
		resJ := f[i] + w(i, x)
		return resI < resJ  // **根据题目修改**
	}

	firstFind := func(i int, tail opt) int {
		left, right := tail.l - 1, tail.r + 1
		for left + 1 < right {
			mid := left + (right - left) >> 1
			if better(i, tail.j, mid) {
				right = mid
			} else {
				left = mid
			}
		}
		return right
	}

	for i := 1; i <= n; i++ {
		for q.front().r < i {
			q.popFront()
		}

		best := q.front().j
		f[i] = f[best] + w(best, i)

		for !q.empty() {
			tail := q.back()
			
			// 1. i 比 j 在 [l, r] 区间内都更优，则出队
			if better(i, tail.j, tail.l) {
				q.popBack()
				continue
			}

			// 2. j 比 i 在 [l, r] 区间更优，i 没用
			if !better(i, tail.j, tail.r) {
				break
			}
			
			// 3. i 在 [l, r] 之间有一部分比 j 更优
			//    由决策单调性，一定是右侧，因此二分
			pos := firstFind(i, tail)
			q.popBack()
			q.pushBack(opt{tail.j, tail.l, pos - 1})
			break
		}

		start := 0
		if q.empty() {
			start = i + 1
		} else {
			start = q.back().r + 1
		}

		// 如果是 2 情况退出，可能会遇到不需要将 i 入堆的情况
		if start <= n {
			q.pushBack(opt{i, start, n})
		}
	}
	return f[n]
}

type opt struct {
	j int
	l int
	r int
}

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