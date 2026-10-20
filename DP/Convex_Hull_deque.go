package dp

// --- 下凸包（求最小值） --- //
// lh := LowerHull

// type LowerHull struct {
// 	hull deque[vec]
// }

// func (lh *LowerHull) add(p vec) {
// 	for lh.hull.size() > 1 {
// 		back := lh.hull.back()
// 		prev := lh.hull.get(lh.hull.size() - 2)
// 		if back.sub(prev).detCmp(p.sub(back)) <= 0 {
// 			lh.hull.popBack()
// 		} else {
// 			break
// 		}
// 	}
// 	lh.hull.pushBack(p)
// }

// // 保证 v0.x 单调
// // 如果 v0.x 单调递增，则 LowerHull.queryMonotonic(v0,  1)
// // 如果 v0.x 单调递减，则 LowerHull.queryMonotonic(v0, -1)
// // 复杂度 O(n)
// func (lh *LowerHull) queryMonotonic(p vec, dir int) int {
// 	if dir < 0 {
// 		for lh.hull.size() > 1 && p.dot(lh.hull.get(0)) >= p.dot(lh.hull.get(1)) {
// 			lh.hull.popFront()
// 		}
// 		return p.dot(lh.hull.front())
// 	} else {
// 		for lh.hull.size() > 1 && p.dot(lh.hull.back()) >= p.dot(lh.hull.get(lh.hull.size()-2)) {
// 			lh.hull.popBack()
// 		}
// 		return p.dot(lh.hull.back())
// 	}
// }

// // 二分查询最小值，复杂度 O(nlogn)
// func (lh *LowerHull) queryBinary(p vec) int {
// 	l, r := 0, lh.hull.size()-1
// 	for l < r {
// 		mid := (l + r) >> 1
// 		if p.dot(lh.hull.get(mid)) >= p.dot(lh.hull.get(mid+1)) {
// 			l = mid + 1
// 		} else {
// 			r = mid
// 		}
// 	}
// 	return p.dot(lh.hull.get(l))
// }

// func (lh *LowerHull) empty() bool {
// 	return lh.hull.empty()
// }

// func (lh *LowerHull) clear() {
// 	lh.hull = deque[vec]{}
// }

// // addFront 向下凸包头部插入元素（用于启发式合并）
// func (lh *LowerHull) addFront(p vec) {
// 	for lh.hull.size() > 1 {
// 		front := lh.hull.front()
// 		next := lh.hull.get(1)

// 		if front.sub(p).detCmp(next.sub(front)) <= 0 {
// 			lh.hull.popFront()
// 		} else {
// 			break
// 		}
// 	}
// 	lh.hull.pushFront(p)
// }

// // merge 启发式合并两个下凸包
// // 要求：lh1 中的所有点的 x 坐标 严格小于 lh2 中所有点的 x 坐标
// func merge(lh1, lh2 LowerHull) LowerHull {
// 	if lh1.hull.size() <= lh2.hull.size() {
// 		for !lh1.empty() {
// 			lh2.addFront(lh1.hull.popBack())
// 		}
// 		return lh2
// 	} else {
// 		for !lh2.empty() {
// 			lh1.add(lh2.hull.popFront())
// 		}
// 		return lh1
// 	}
// }
