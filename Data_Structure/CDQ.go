package Data_Structure

import (
	"slices"
	"sort"
	"cmp"
)

type tuple struct {
	a, b, c, cnt, ans int
}

func CDQ(N int, A, B, C []int) {
	ra, _ := discretize(A)
	rb, _ := discretize(B)
	rc, _ := discretize(C)

	Nums := make([]tuple, N)
	for i := range Nums {
		Nums[i] = tuple{ra[i], rb[i], rc[i], 1, 0}
	}

	slices.SortFunc(Nums, func(x, y tuple) int {
		if x.a != y.a {
			return cmp.Compare(x.a, y.a)
		}
		if x.b != y.b {
			return cmp.Compare(x.b, y.b)
		}
		return cmp.Compare(x.c, y.c)
	})

	// nums[i].ans 表示下标在 i 之前的满足条件的 (j, i) 组的个数
	nums := []tuple{Nums[0]}
	for i := 1; i < N; i++ {
		last := nums[len(nums)-1]
		if last.a == Nums[i].a && last.b == Nums[i].b && last.c == Nums[i].c {
			nums[len(nums)-1].cnt++
		} else {
			nums = append(nums, Nums[i])
		}
	}

	for i := range nums {
		nums[i].ans = nums[i].cnt - 1
	}

	bit := newFenwickTree(N)
	n := len(nums)
	tmp := make([]tuple, n)
	var cdq func(int, int)
	cdq = func(l, r int) {
		if l == r {
			return
		}
		mid := l + (r - l) >> 1
		cdq(l, mid)
		cdq(mid + 1, r)

		j, k := l, l
		for i := mid + 1; i <= r; i++ {
			for ; j <= mid && nums[j].b <= nums[i].b; j++ {
				bit.update(nums[j].c, nums[j].cnt)
				tmp[k] = nums[j]
				k++
			}
			nums[i].ans += bit.query(0, nums[i].c)
			tmp[k] = nums[i]
			k++
		}

		for p := l; p < j; p++ {
			bit.update(nums[p].c, -nums[p].cnt)
		}

		for ; j <= mid; j++ {
			tmp[k] = nums[j]
			k++
		}

		copy(nums[l:r+1], tmp[l:r+1])
	}

	cdq(0, n-1)

	// ans[i] 表示有 i 个逆序对的下标个数
	ans := make([]int, N)
	for _, x := range nums {
		ans[x.ans] += x.cnt
	}
}

// 返回离散化后的数组以及去重后的元素个数 m
func discretize(arr []int) ([]int, int) {
    unique := slices.Clone(arr)
	slices.Sort(unique)
	unique = slices.Compact(unique)

    m := len(unique)
    rank := make([]int, len(arr))
	for i, x := range arr {
		rank[i] = sort.SearchInts(unique, x)
	}

	return rank, m
}
