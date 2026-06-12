package dp

func WQS(limit int) int {
	dp := func(k int) pair {
		res := pair{0, 0}
		_ = k
		return res
	}

	var res pair

	// 如果是 至多/至少 k 个，需要加上优化
	// res = dp(0)
	// if res.x <= limit {  // 至少改成 >=
	// 	return res.b
	// }

	ans := 0
	left, right := -1_000_005, 1_000_005
	for left + 1 < right {
		mid := left + (right - left) >> 1
		res = dp(mid)
		if res.x <= limit {
			ans = mid * limit + res.b
			right = mid  // left = mid  // 下凸壳
		} else {
			left = mid   // right = mid // 下凸壳
		}
	}

	return ans
}

type pair struct{ b, x int }

// p 比 q 更劣
func less(p, q pair) bool {
	return p.b < q.b || p.b == q.b && p.x > q.x
}