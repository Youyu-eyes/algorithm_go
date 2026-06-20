package Data_Structure

import (
	"math"
	"slices"
	"cmp"
)

// 普通莫队（奇偶排序优化）
// 母题：https://codeforces.com/contest/86/problem/D

func Mo(a []int, queries [][]int) {
	n := len(a)
	q := len(queries)

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := []query{}

	
	ans := make([]int, q)
	for Q := 0; Q < q; Q++ {
		l, r := queries[Q][0] - 1, queries[Q][1] + 1
		qs = append(qs, query{l / B, l, r, Q})
	}

	// 奇偶排序优化
	slices.SortFunc(qs, func(a, b query) int {
		if a.id != b.id {
			return cmp.Compare(a.id, b.id)
		}
		// 奇数块 r 降序，偶数块 r 升序
		if a.id & 1 == 1 {
			return cmp.Compare(b.r, a.r)
		}
		return cmp.Compare(a.r, b.r)
	})

	// 初始化左右端点
	l, r, res := -1, 0, 0

	for _, b := range qs {		
		// 右端点右移
		for ; r < b.r; r++ {

		}

		// 左端点左移
		for ; l > b.l; l-- {

		}

		// 右端点左移
		// 开区间，先左移再删除
		for r > b.r {
			r--

		}

		// 左端点右移
		// 开区间，先右移再删除
		for l < b.l {
			l++

		}

		ans[b.qIdx] = res
	}
}


// 回滚莫队

func RollbackMo(a []int, queries [][]int) {
	n := len(a)
	q := len(queries)

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := []query{}

	var res int
	ans := make([]int, q)
	for Q := 0; Q < q; Q++ {
		l, r := queries[Q][0], queries[Q][1] + 1

		// 大区间离线
		if r - l > B {
			qs = append(qs, query{l / B, l - 1, r, Q})
			continue
		}

		// 小区间暴力
		for i := l; i < r; i++ {

		}
	}

	slices.SortFunc(qs, func(a, b query) int {
		return cmp.Or(
			cmp.Compare(a.id, b.id),
			cmp.Compare(a.r, b.r),
		)
	})

	var l, r int
	for i, b := range qs {
		start := (b.id + 1) * B
		if i == 0 || b.id > qs[i - 1].id {
			l = start - 1
			r = start
			res = 0
		}
		
		// 右端点右移
		for ; r < b.r; r++ {

		}

		// 保留状态
		tmp := res

		// 左端点左移
		for ; l > b.l; l-- {

		}

		ans[b.qIdx] = res

		// 回滚
		res = tmp
		l = start - 1
		for j := b.l + 1; j <= l; j++ {
			
		}
	}
}