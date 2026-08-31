package Data_Structure

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 时间复杂度 O(log n)
func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 求区间和 a[l] + ... + a[r]
// 时间复杂度 O(log n)
// 0 <= l <= r < n
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}


// ======= 值域树状数组 ======= //

// type pair struct{ cnt, sum int }
// type fenwick struct {
// 	t       []pair
// 	sorted  []int
// 	highBit int
// }

// func newFenwickTree(sorted []int) fenwick {
// 	n := len(sorted)
// 	return fenwick{
// 		t:       make([]pair, n+1),
// 		sorted:  sorted,
// 		highBit: 1 << (bits.Len(uint(n)) - 1),
// 	}
// }

// // 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
// // 如果 num < 0，表示减少 -num 个 val
// // 注意 val = sorted[i-1]，无需手动传入
// func (f *fenwick) update(i, num int) {
// 	val := f.sorted[i-1]
// 	for ; i < len(f.t); i += i & -i {
// 		f.t[i].cnt += num
// 		f.t[i].sum += num * val
// 	}
// }

// // 返回第 k 小的数（k 从 1 开始）
// func (f *fenwick) kth(k int) int {
// 	i := 0
// 	for b := f.highBit; b > 0; b >>= 1 {
// 		if nxt := i | b; nxt < len(f.t) && f.t[nxt].cnt < k {
// 			k -= f.t[nxt].cnt
// 			i = nxt
// 		}
// 	}
// 	return f.sorted[i]
// }

// // 返回前 k 小的数之和（k 从 1 开始）
// func (f *fenwick) preSum(k int) (s int) {
// 	i := 0
// 	for b := f.highBit; b > 0; b >>= 1 {
// 		if nxt := i | b; nxt < len(f.t) && f.t[nxt].cnt < k {
// 			k -= f.t[nxt].cnt
// 			s += f.t[nxt].sum
// 			i = nxt
// 		}
// 	}
// 	// 加上等于第 k 小的数
// 	s += f.sorted[i] * k
// 	return
// }
