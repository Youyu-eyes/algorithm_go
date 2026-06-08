package dp

// 本模板 解决的是 至多 选择 limit 次的最优答案
// 并且是 f(x) 是上凸壳

func WQS(limit int) int {
	dp := func(k int) pair {
		res := pair{0, 0}
		_ = k
		return res
	}

	res := dp(0)
	if res.x <= limit {  // 至多 limit
		return res.b
	}

	ans := 0
	left, right := 0, 1_000_005
	for left + 1 < right {
		mid := left + (right - left) >> 1
		res = dp(mid)
		if res.b <= limit {
			ans = mid * limit + res.b
			right = mid
		} else {
			left = mid
		}
	}

	return ans
}

type pair struct{ b, x int }

// p 比 q 更劣
func less(p, q pair) bool {
	return p.b < q.b || p.b == q.b && p.x > q.x
}