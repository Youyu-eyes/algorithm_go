package Data_Structure

// 返回 left，其中 left[i] 是 nums[i] 左侧最近的严格大于 nums[i] 的数的下标，若不存在则为 -1
func leftGreater(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	for i := range left { left[i] = -1 }
	st := []int{-1} // 哨兵
	for i, x := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] <= x { // 如果求严格小于，改成 >=
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	return left
}

// 返回 right，其中 right[i] 是 nums[i] 右侧最近的严格大于 nums[i] 的数的下标，若不存在则为 len(nums)
func rightGreater(nums []int) []int {
	n := len(nums)
	right := make([]int, n)
	for i := range right { right[i] = n }
	st := []int{n} // 哨兵
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		for len(st) > 1 && nums[st[len(st)-1]] <= x { // 如果求严格小于，改成 >=
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}
	return right
}
