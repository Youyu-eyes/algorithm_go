package Data_Structure

import (
	"math"
)

func _(a []int) {
	n := len(a)
	B := int(math.Sqrt(float64(n)))
	//B := int(math.Sqrt(float64(n * bits.Len(uint(n)))))

	type block struct {
		l, r int // [l, r)
		sum  int
		todo int
	}
	blocks := make([]block, (n-1)/B+1)
	calcBlock := func(l, r int) (res int) {
		for j := l; j < r; j++ {
			v := a[j]
			_ = v
		}
		return
	}
	for i := 0; i < n; i += B {
		r := min(i+B, n)
		sum := calcBlock(i, r)
		blocks[i/B] = block{i, r, sum, 0}
	}

	// [l, r), 从 0 开始
	sqrtOp := func(l, r int, v int) {
		for i := range blocks {
			b := &blocks[i]
			if b.l >= r {
				break
			}
			if b.r <= l {
				continue
			}
			if l <= b.l && b.r <= r {
				// 完整块
			} else {
				// 部分块
				bl := max(b.l, l)
				br := min(b.r, r)
				for j := bl; j < br; j++ {

				}
			}
		}
	}

	_ = sqrtOp
}