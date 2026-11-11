package string

// s 和 t 的下标转换关系：
// (si+1)*2 = ti
// ti/2-1 = si
func manacher(s string) {
	n := len(s)
	t := append(make([]byte, 0, n*2+3), '^')
	for _, c := range s {
		t = append(t, '#', byte(c))
	}
	t = append(t, '#', '$')

	// []int 切片
	// t := append(make([]int, 0, len(s)*2+3), -2)
	// for _, c := range s {
	// 	t = append(t, -1, c)
	// }
	// t = append(t, -1, -3)

	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1

	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ {
		hl := 1
		if i < boxR {
			hl = min(boxR-i, halfLen[boxM*2-i])
		}

		for t[i-hl] == t[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}

		halfLen[i] = hl
	}

	// 判断子串 s[l:r]（左闭右开）是否为回文串
	isPalindrome := func(l, r int) bool {
		return halfLen[l+r+1] > r-l
	}

	// isOdd=true:  返回以 s[i] 为回文中心的最长奇回文子串长度
	// isOdd=false: 返回以 s[i] 和 s[i+1] 为回文中心的最长偶回文子串长度（若不存在，则为 0）
	midLen := func(i int, isOdd bool) int {
		if isOdd {
			return halfLen[i*2+2] - 1
		}
		return halfLen[i*2+3] - 1
	}
	
	isPalindrome(0, 0)
	midLen(0, true)
}
