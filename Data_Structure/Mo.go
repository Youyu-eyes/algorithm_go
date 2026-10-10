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
	m := len(queries)

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := make([]query, m)

	
	ans := make([]int, m)
	for qi := range m {
		l, r := queries[qi][0], queries[qi][1] + 1
		qs[qi] = query{l / B, l - 1, r, qi}
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

	res := 0

	add := func(x int) {

	}

	del := func(x int) {

	}

	// 初始化左右端点
	l, r := -1, 0

	for _, b := range qs {		
		// 右端点右移
		for ; r < b.r; r++ {
			add(a[r])
		}

		// 左端点左移
		for ; l > b.l; l-- {
			add(a[l])
		}

		// 右端点左移
		// 开区间，先左移再删除
		for r > b.r {
			r--
			del(a[r])
		}

		// 左端点右移
		// 开区间，先右移再删除
		for l < b.l {
			l++
			del(a[l])
		}

		ans[b.qIdx] = res
	}
}


// 回滚莫队

func RollbackMo(a []int, queries [][]int) {
	n := len(a)
	m := len(queries)

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := []query{}

	var res int
	ans := make([]int, m)
	for qi := range m {
		l, r := queries[qi][0], queries[qi][1] + 1

		// 大区间离线
		if r - l > B {
			qs = append(qs, query{l / B, l - 1, r, qi})
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


// 带修莫队

func ModifyMo(a []int, queries [][]int) {
	n, m := len(a), len(queries)

	B := int(max(1, math.Pow(float64(n), 2.0/3.0)))

	type query struct {
		lid, rid int
		l, r, t     int
		qIdx     int
	}
	type version struct {
		pos, color int
	}

	qs := []query{}
	vs := []version{}
	curQ := 0

	for qi := range m {
		op := queries[qi][0]
		if op == 1 {
			l, r := queries[qi][0], queries[qi][1]
			qs = append(qs, query{l/B, r/B, l - 1, r + 1, qi - curQ, curQ})
			curQ++
		} else {
			idx, val := queries[qi][0], queries[qi][1]
			vs = append(vs, version{idx, val})
		}
	}

	slices.SortFunc(qs, func(a, b query) int {
		if a.lid != b.lid {
			return cmp.Compare(a.lid, b.lid)
		}
		if a.rid != b.rid {
			return cmp.Compare(a.rid, b.rid)
		}
		return cmp.Compare(a.t, b.t)
	})

	res := 0
	add := func(x int) {

	}

	del := func(x int) {

	}

	upd := func(q query, t int) {
		pos, color := vs[t].pos, vs[t].color
		if q.l < pos && pos < q.r {
			del(a[pos])
			add(color)
		}
		a[pos], vs[t].color = vs[t].color, a[pos]
	}

	ans := make([]int, curQ)
	l, r, t := -1, 0, 0
	for _, b := range qs {
		for ; r < b.r; r++ {
			add(a[r])
		}
		for ; l > b.l; l-- {
			add(a[l])
		}
		for r > b.r {
			r--
			del(a[r])
		}
		for l < b.l {
			l++
			del(a[l])
		}
		for ; t < b.t; t++ {
			upd(b, t)
		}
		for t > b.t {
			t--
			upd(b, t)
		}
		ans[b.qIdx] = res
	}
}
