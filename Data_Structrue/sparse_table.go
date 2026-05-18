package Data_Structrue

import (
	"math/bits"
)

// ------- 数值 ST 表 ------- //
// 初始化：st := newSparseTable(nums, func(a, b int) int { return max(a, b) })

type SparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

// op 需满足结合律（如 min, max, gcd 等）
func newSparseTable[T any](nums []T, op func(T, T) T) *SparseTable[T] {
	n := len(nums)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], nums)
	for i := 1; i < w; i++ {
		step := 1 << (i - 1)
		for j := 0; j+step < n && j+(1<<i) <= n; j++ {
			st[i][j] = op(st[i-1][j], st[i-1][j+step])
		}
	}
	return &SparseTable[T]{st: st, op: op}
}

func (st *SparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return st.op(st.st[k][l], st.st[k][r-1<<k])
}


// ------- 下标版本 ST 表 ------- //
// 初始化：st := newSparseTableIndex(nums, func(a, b int) int { return max(a, b) })

type SparseTableIndex[T comparable] struct {
	nums []T
	st   [][]int // 存储最值对应的下标
	op   func(T, T) T
}

// 当区间最值有多个时，返回最左侧的下标
func newSparseTableIndex[T comparable](nums []T, op func(T, T) T) *SparseTableIndex[T] {
	n := len(nums)
	w := bits.Len(uint(n))
	st := make([][]int, w)
	for i := range st {
		st[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		st[0][i] = i
	}
	for i := 1; i < w; i++ {
		step := 1 << (i - 1)
		for j := 0; j+step < n && j+(1<<i) <= n; j++ {
			leftIdx := st[i-1][j]
			rightIdx := st[i-1][j+step]
			if op(nums[leftIdx], nums[rightIdx]) == nums[leftIdx] {
				st[i][j] = leftIdx
			} else {
				st[i][j] = rightIdx
			}
		}
	}
	return &SparseTableIndex[T]{nums: nums, st: st, op: op}
}

func (st *SparseTableIndex[T]) query(l, r int) int {
	length := r - l
	k := bits.Len(uint(length)) - 1
	leftIdx := st.st[k][l]
	rightIdx := st.st[k][r-1<<k]
	if st.op(st.nums[leftIdx], st.nums[rightIdx]) == st.nums[leftIdx] {
		return leftIdx
	}
	return rightIdx
}


// ------- 二维 ST 表 ------- //
// 初始化：stMax := newSparseTable2D(matrix, func(a, b int) int { return max(a, b) })

type SparseTable2D[T any] struct {
	st [][][][]T
	op func(T, T) T
}

func newSparseTable2D[T any](matrix [][]T, op func(T, T) T) *SparseTable2D[T] {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return &SparseTable2D[T]{st: nil, op: op}
	}

	m, n := len(matrix), len(matrix[0])
	logM := bits.Len(uint(m))
	logN := bits.Len(uint(n))

	st := make([][][][]T, logM)
	for k := range st {
		st[k] = make([][][]T, logN)
		for l := range st[k] {
			st[k][l] = make([][]T, m)
			for i := range st[k][l] {
				st[k][l][i] = make([]T, n)
			}
		}
	}

	for i := 0; i < m; i++ {
		copy(st[0][0][i], matrix[i])
	}

	for l := 1; l < logN; l++ {
		stepC := 1 << (l - 1)
		for i := 0; i < m; i++ {
			for j := 0; j+(1<<l) <= n; j++ {
				st[0][l][i][j] = op(st[0][l-1][i][j], st[0][l-1][i][j+stepC])
			}
		}
	}

	for k := 1; k < logM; k++ {
		stepR := 1 << (k - 1)
		for l := 0; l < logN; l++ {
			for i := 0; i+(1<<k) <= m; i++ {
				for j := 0; j+(1<<l) <= n; j++ {
					st[k][l][i][j] = op(st[k-1][l][i][j], st[k-1][l][i+stepR][j])
				}
			}
		}
	}

	return &SparseTable2D[T]{st: st, op: op}
}

func (st *SparseTable2D[T]) query(r1, r2, c1, c2 int) T {
	k := bits.Len(uint(r2-r1)) - 1
	l := bits.Len(uint(c2-c1)) - 1

	lenR := 1 << k
	lenC := 1 << l

	op1 := st.st[k][l][r1][c1]
	op2 := st.st[k][l][r2-lenR][c1]
	op3 := st.st[k][l][r1][c2-lenC]
	op4 := st.st[k][l][r2-lenR][c2-lenC]

	return st.op(st.op(op1, op2), st.op(op3, op4))
}