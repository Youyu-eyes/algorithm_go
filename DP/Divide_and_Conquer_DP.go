package dp

func Divide_and_Conquer(n int) {
	f := make([]int, n + 1)
    nf := make([]int, n + 1)

	w := func(j, i int) int {  // 表示从 j 转移到 i
		return j + i  // **根据题目修改**
	}

	var op func(int, int, int, int)
	op = func(l, r, optL, optR int) {
		if l > r {
			return
		}

		mid := l + (r - l) >> 1
		best := -1
		for j := optL; j <= min(optR, mid - 1); j++ {
			if best < 0 || f[j] + w(j, mid) < f[best] + w(best, mid) {  // **根据题目修改**
				best = j
			}
		}
		nf[mid] = f[best] + w(best, mid)
		op(l, mid - 1, optL, best)
		op(mid + 1, r, best, optR)
	}
}