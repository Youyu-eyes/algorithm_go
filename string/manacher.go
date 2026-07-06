package string

type IntegerManacher interface {
	~int | ~int64 | ~byte | ~rune
}

// Manacher 接受一个泛型切片 s，返回一个长度为 2n-1 的切片 ans。
// ans[i] 的含义如下（以下的 i 全部指代 ti）：
// - 当 i 为偶数时，表示以 s[i/2] 为中心的最长奇回文子串（或子数组）的长度
// - 当 i 为奇数时，表示以 s[i/2] 和 s[i/2+1] 之间为中心的最长偶回文子串（或子数组）的长度
// 左端点：start = (i - L + 1) / 2
// 右端点：end   = (i + L - 1) / 2

func Manacher[T IntegerManacher](s []T) []int {
	n := len(s)
	if n == 0 {
		return nil
	}

	// 虚拟数组的长度 m，逻辑上等同于包含了首尾防越界符以及中间的 '#'
	// 对应关系：
	// 0: 首边界, 1: #, 2: s[0], 3: #, 4: s[1] ... 2n: s[n-1], 2n+1: #, 2n+2: 尾边界
	m := 2 * n + 3
	halfLen := make([]int, m)
	boxM, boxR := 0, 0

	// 遍历所有原数组范围内可能的回文中心
	// 从虚拟下标 2 (即 s[0]) 遍历到 2n (即 s[n-1])
	// 这正好包含了 n 个元素中心和 n-1 个间隙中心，共 2n-1 个中心
	for i := 2; i <= 2 * n; i++ {
		hl := 1
		if i < boxR {
			hl = min(halfLen[boxM * 2 - i], boxR - i)
		}

		// 暴力向两边扩展
		for {
			left, right := i - hl, i + hl

			// 1. 如果触碰到逻辑上的首尾边界，停止扩展
			if left == 0 || right == m - 1 {
				break
			}
			// 2. 如果扩展到了虚拟的 '#' 位置，必然相等，直接半径 +1 继续
			if left % 2 == 1 {
				hl++
				continue
			}
			// 3. 如果扩展到了原数组的元素，映射回原切片进行比较
			// 虚拟下标 x 对应的原数组下标为 x/2 - 1
			if s[left / 2 - 1] == s[right / 2 - 1] {
				hl++
			} else {
				break
			}
		}

		// 更新最右回文边界及其中心
		if i + hl > boxR {
			boxM, boxR = i, i + hl
		}
		halfLen[i] = hl
	}

	ans := make([]int, 2 * n - 1)
	for i := 0; i < 2 * n - 1; i++ {
		ans[i] = halfLen[i + 2] - 1
	}

	return ans
}
