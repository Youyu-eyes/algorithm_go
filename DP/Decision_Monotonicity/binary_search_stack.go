// 二分栈优化 DP
// 用双端队列维护最优决策点，以及决策点的作用域
// 这是一维情况，如果涉及 2 维，需要在结构体中额外定义一个快照
// 二维情况见 CF 321E

package dp

func BinaryStack(n int) int {
	f := make([]int, n + 1)

	q := []opt{{0, 1, n}}

	w := func (j, i int) int {
		return j + i  // **根据题目修改**
	}

	better := func(i, j, x int) bool {
		resI := f[i] + w(i, x)
		resJ := f[j] + w(j, x)
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
		for q[0].r < i {
			q = q[1:]
		}

		best := q[0].j
		f[i] = f[best] + w(best, i)
		
		for len(q) > 0 {
			tail := q[len(q)-1]
			
			// 1. i 比 j 在 [l, r] 区间内都更优，则出队
			if better(i, tail.j, tail.l) {
				q = q[:len(q)-1]
				continue
			}

			// 2. j 比 i 在 [l, r] 区间更优，i 没用
			if !better(i, tail.j, tail.r) {
				break
			}
			
			// 3. i 在 [l, r] 之间有一部分比 j 更优
			//    由决策单调性，一定是右侧，因此二分
			pos := firstFind(i, tail)
			q = q[:len(q)-1]
			q = append(q, opt{tail.j, tail.l, pos - 1})
			break
		}

		start := 0
		if len(q) == 0 {
			start = i + 1
		} else {
			start = q[len(q)-1].r + 1
		}

		// 如果是 2 情况退出，可能会遇到不需要将 i 入堆的情况
		if start <= n {
			q = append(q, opt{i, start, n})
		}
	}
	return f[n]
}

type opt struct {
	j int
	l int
	r int
}
