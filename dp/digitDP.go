package dp

import (
	"math/bits"
	"strconv"
)

// 数位DP 计算满足条件的数字的数量
func digitDP(low, high int, target int) int {
	lowS := strconv.Itoa(low)
	highS := strconv.Itoa(high)
	// 二进制
	// lowS := strconv.FormatInt(int64(low), 2)
	// highS := strconv.FormatInt(int64(high), 2)

	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int
	dfs = func(i, cnt0 int, limitLow, limitHigh bool) (res int) {
		// 不合法
		if cnt0 > target {
			return 0
		}
		if i == n {
			// 不合法
			if cnt0 < target {
				return 0
			}
			// 合法
			return 1
		}
		if !limitLow && !limitHigh {
			p := &memo[i][cnt0]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9 // 二进制填 1
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		// 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
		// 如果前导零不影响答案，去掉这个 if block
		if limitLow && i < diffLH {
			// 不填数字，上界不受约束
			res = dfs(i+1, 0, true, false)
			d = 1 // 下面填数字，从 1 开始填
		}

		for ; d <= hi; d++ {
			c0 := cnt0
			if d == 0 {
				c0++ // 统计 0 的个数
			}
			res += dfs(i+1,
				c0,
				limitLow && d == lo,
				limitHigh && d == hi)
			// res %= mod
		}
		return
	}

	return dfs(0, 0, true, true)
}

// 计算满足条件的数的价值
// 计算在 [low, high] 中的整数 x 的数位和，满足 x 中的不同数字个数不超过 k
func digitDPContribution(low, high int, k int) int {
	lowS := strconv.Itoa(low)
	highS := strconv.Itoa(high)
	// 二进制
	// lowS := strconv.FormatInt(int64(low), 2)
	// highs := strconv.FormatInt(int64(high), 2)

	n := len(highS)
	diffLH := n - len(lowS)
	type pair struct{ cnt, sum int }
	memo := make([][1 << 10]pair, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j].cnt = -1
		}
	}

	// dfs 返回两个数：子树合法数字个数，子树数位总和
	var dfs func(int, int, bool, bool) pair
	dfs = func(i, mask int, limitLow, limitHigh bool) (res pair) {
		if i == n {
			// 如果没有特殊约束，那么能递归到终点的都是合法数字
			return pair{1, 0}
		}
		if !limitLow && !limitHigh {
			p := &memo[i][mask]
			if p.cnt >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		// 如果前导零不影响答案，可以去掉这个 if block
		if limitLow && i < diffLH {
			// 不填数字，上界不受约束
			res = dfs(i+1, 0, true, false)
			d = 1
		}

		for ; d <= hi; d++ {
			newMask := mask | 1<<d
			if bits.OnesCount(uint(newMask)) > k { // 不满足要求
				continue
			}
			sub := dfs(i+1,
				newMask,
				limitLow && d == lo,
				limitHigh && d == hi)
			res.cnt += sub.cnt     // 累加子树的合法数字个数
			res.sum += sub.sum     // 累加子树的数位总和
			res.sum += d * sub.cnt // d 会出现在 sub.cnt 个数中（贡献法）
			// res.cnt %= mod; res.sum %= mod
		}
		return
	}

	return dfs(0, 0, true, true).sum
}
