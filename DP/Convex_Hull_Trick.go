package dp

import (
	"math/big"
)

// 在使用模板时要注意：
// 首先，在维护凸包时要保证查询向量 v0 的 y 值 v0.y > 0，这样我们才能用 max or min 简单地判断用 上凸包 or 下凸包
// 其次，在维护凸包时要保证待维护向量 v1 的 x 值 v1.x 单调递增

// --- 向量模板 --- //
// 初始化：v := vec{x, y}

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

// a.det(b) > 0 => a 到 b 逆时针
// a.det(b) < 0 => a 到 b 顺时针
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x } // 如果乘法会溢出，用 detCmp
func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

// --- 泛型双端队列 --- //

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

// --- 上凸包（求最大值） --- //
// uh := UpperHull{}

type UpperHull struct {
	hull deque[vec]
}

func (uh *UpperHull) add(p vec) {
	for uh.hull.size() > 1 {
		back := uh.hull.back()
		prev := uh.hull.get(uh.hull.size() - 2)
		if back.sub(prev).detCmp(p.sub(back)) >= 0 {
			uh.hull.popBack()
		} else {
			break
		}
	}
	uh.hull.pushBack(p)
}

// 保证 v0.x 单调
// 如果 v0.x 单调递增，则 UpperHull.queryMonotonic(v0,  1)
// 如果 v0.x 单调递减，则 UpperHull.queryMonotonic(v0, -1)
// 复杂度 O(n)
func (uh *UpperHull) queryMonotonic(p vec, dir int) int {
	if dir > 0 {
		for uh.hull.size() > 1 && p.dot(uh.hull.get(0)) <= p.dot(uh.hull.get(1)) {
			uh.hull.popFront()
		}
		return p.dot(uh.hull.front())
	} else {
		for uh.hull.size() > 1 && p.dot(uh.hull.back()) <= p.dot(uh.hull.get(uh.hull.size()-2)) {
			uh.hull.popBack()
		}
		return p.dot(uh.hull.back())
	}
}

// 二分查询最大值，复杂度 O(nlogn)
func (uh *UpperHull) queryBinary(p vec) int {
	l, r := 0, uh.hull.size()-1
	for l < r {
		mid := (l + r) >> 1
		if p.dot(uh.hull.get(mid)) <= p.dot(uh.hull.get(mid+1)) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return p.dot(uh.hull.get(l))
}

// addFront 向下凸包头部插入元素（用于启发式合并）
func (uh *UpperHull) addFront(p vec) {
	for uh.hull.size() > 1 {
		front := uh.hull.front()
		next := uh.hull.get(1)
		
		if front.sub(p).detCmp(next.sub(front)) >= 0 {
			uh.hull.popFront()
		} else {
			break
		}
	}
	uh.hull.pushFront(p)
}

// mergeUpper 启发式合并两个上凸包
// 要求：uh1 中的所有点的 x 坐标 严格小于 uh2 中所有点的 x 坐标
func mergeUpper(uh1, uh2 UpperHull) UpperHull {
	if uh1.hull.size() <= uh2.hull.size() {
		for !uh1.empty() {
			uh2.addFront(uh1.hull.popBack())
		}
		return uh2
	} else {
		for !uh2.empty() {
			uh1.add(uh2.hull.popFront())
		}
		return uh1
	}
}

func (uh *UpperHull) empty() bool {
	return uh.hull.empty()
}

func (uh *UpperHull) clear() {
	uh.hull = deque[vec]{}
}

// --- 下凸包（求最小值） --- //
// lh := LowerHull

type LowerHull struct {
	hull deque[vec]
}

func (lh *LowerHull) add(p vec) {
	for lh.hull.size() > 1 {
		back := lh.hull.back()
		prev := lh.hull.get(lh.hull.size() - 2)
		if back.sub(prev).detCmp(p.sub(back)) <= 0 {
			lh.hull.popBack()
		} else {
			break
		}
	}
	lh.hull.pushBack(p)
}

// 保证 v0.x 单调
// 如果 v0.x 单调递增，则 LowerHull.queryMonotonic(v0,  1)
// 如果 v0.x 单调递减，则 LowerHull.queryMonotonic(v0, -1)
// 复杂度 O(n)
func (lh *LowerHull) queryMonotonic(p vec, dir int) int {
	if dir < 0 {
		for lh.hull.size() > 1 && p.dot(lh.hull.get(0)) >= p.dot(lh.hull.get(1)) {
			lh.hull.popFront()
		}
		return p.dot(lh.hull.front())
	} else {
		for lh.hull.size() > 1 && p.dot(lh.hull.back()) >= p.dot(lh.hull.get(lh.hull.size()-2)) {
			lh.hull.popBack()
		}
		return p.dot(lh.hull.back())
	}
}

// 二分查询最小值，复杂度 O(nlogn)
func (lh *LowerHull) queryBinary(p vec) int {
	l, r := 0, lh.hull.size()-1
	for l < r {
		mid := (l + r) >> 1
		if p.dot(lh.hull.get(mid)) >= p.dot(lh.hull.get(mid+1)) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return p.dot(lh.hull.get(l))
}

// addFront 向下凸包头部插入元素（用于启发式合并）
func (lh *LowerHull) addFront(p vec) {
	for lh.hull.size() > 1 {
		front := lh.hull.front()
		next := lh.hull.get(1)

		if front.sub(p).detCmp(next.sub(front)) <= 0 {
			lh.hull.popFront()
		} else {
			break
		}
	}
	lh.hull.pushFront(p)
}

// merge 启发式合并两个下凸包
// 要求：lh1 中的所有点的 x 坐标 严格小于 lh2 中所有点的 x 坐标
func mergeLower(lh1, lh2 LowerHull) LowerHull {
	if lh1.hull.size() <= lh2.hull.size() {
		for !lh1.empty() {
			lh2.addFront(lh1.hull.popBack())
		}
		return lh2
	} else {
		for !lh2.empty() {
			lh1.add(lh2.hull.popFront())
		}
		return lh1
	}
}

func (lh *LowerHull) empty() bool {
	return lh.hull.empty()
}

func (lh *LowerHull) clear() {
	lh.hull = deque[vec]{}
}
