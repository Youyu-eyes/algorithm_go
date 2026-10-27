package Data_Structure

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1)
}

func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

// 返回最小的 0-based 下标 i，满足 a[0] + ... + a[i] >= k
// 若总和 < k，返回 n
func (f fenwick) kth(k int) int {
	n := len(f) - 1
	bit := 1
	for bit<<1 <= n {
		bit <<= 1
	}
	idx := 0
	for ; bit > 0; bit >>= 1 {
		nxt := idx + bit
		if nxt <= n && f[nxt] < k {
			k -= f[nxt]
			idx = nxt
		}
	}
	return idx
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

// func (f *fenwick) update(i, num int) {
// 	val := f.sorted[i]
// 	for i++; i < len(f.t); i += i & -i {
// 		f.t[i].cnt += num
// 		f.t[i].sum += num * val
// 	}
// }

// // 第 k 小（0-base）
// func (f *fenwick) kth(k int) int {
// 	i := 0
// 	for b := f.highBit; b > 0; b >>= 1 {
// 		if nxt := i | b; nxt < len(f.t) && f.t[nxt].cnt <= k {
// 			k -= f.t[nxt].cnt
// 			i = nxt
// 		}
// 	}
// 	return f.sorted[i]
// }

// // 前 k 小的和（1-base）
// func (f *fenwick) preSum(k int) (s int) {
// 	i := 0
// 	for b := f.highBit; b > 0; b >>= 1 {
// 		if nxt := i | b; nxt < len(f.t) && f.t[nxt].cnt < k {
// 			k -= f.t[nxt].cnt
// 			s += f.t[nxt].sum
// 			i = nxt
// 		}
// 	}
// 	s += f.sorted[i] * k
// 	return
// }
