package Bitwise

import (
	"fmt"
)

// 对于每个右端点 i，计算所有子数组的或值，打印这些或值的分布范围（子数组左端点范围）
// 时间复杂度 O(nlogU)，其中 U = max(nums)
func logTrick(nums []int) {
	type pair struct{ or, left int } // 子数组或值，最小左端点
	orLeft := []pair{}
	for i, x := range nums {
		// 计算以 i 为右端点的子数组或值
		for j := range orLeft {
			orLeft[j].or |= x // **根据题目修改**
		}
		// x 单独一个数作为子数组
		orLeft = append(orLeft, pair{x, i})

		// 原地去重（相同或值只保留最左边的）
		idx := 1
		for j := 1; j < len(orLeft); j++ {
			if orLeft[j].or != orLeft[j-1].or {
				orLeft[idx] = orLeft[j]
				idx++
			}
		}
		orLeft = orLeft[:idx]

		for k, p := range orLeft {
			orVal := p.or
			left := p.left
			right := i
			if k < len(orLeft)-1 {
				right = orLeft[k+1].left - 1
			}
			// 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都是 orVal
			fmt.Println(left, right, orVal)
		}
	}
}
