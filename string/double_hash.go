package string

type integerHash interface {
	~int | ~int64 | ~byte | ~rune
}

const (
	base1 = 1_000_000_021
	mod1  = 1_000_000_007
	base2 = 1_000_000_033
	mod2  = 1_000_000_009
)

type doubleHash struct {
	pref1, pref2 []int
	pow1, pow2   []int
}

type hashVal struct {
	h1 int
	h2 int
}

// 字符串双哈希模板
// 支持 int，byte，rune 类型
// O(1) 时间返回任意子数组的双哈希值 hashVal

// 可以用于判断两个不同字符串的子串是否相等

func newDoubleHash[T integerHash](arr []T) *doubleHash {
	n := len(arr)
	h := &doubleHash{
		pref1: make([]int, n+1),
		pref2: make([]int, n+1),
		pow1:  make([]int, n+1),
		pow2:  make([]int, n+1),
	}
	h.pow1[0], h.pow2[0] = 1, 1

	for i, val := range arr {
		v := int(val)

		h.pow1[i+1] = h.pow1[i] * base1 % mod1
		h.pow2[i+1] = h.pow2[i] * base2 % mod2

		h.pref1[i+1] = (h.pref1[i]*base1 + v) % mod1
		h.pref2[i+1] = (h.pref2[i]*base2 + v) % mod2
	}
	return h
}

func (h *doubleHash) getHash(l, r int) hashVal {
	if l > r || l < 0 || r >= len(h.pref1)-1 {
		return hashVal{-1, -1}
	}
	length := r - l + 1
	
	hash1 := (h.pref1[r+1] - h.pref1[l]*h.pow1[length]) % mod1
	if hash1 < 0 {
		hash1 += mod1
	}

	hash2 := (h.pref2[r+1] - h.pref2[l]*h.pow2[length]) % mod2
	if hash2 < 0 {
		hash2 += mod2
	}
	
	return hashVal{hash1, hash2}
}

// compareSubstrings 比较 arr[l1: r1+1] 和 arr[l2: r2+1] 的字典序
// 返回 -1 表示 子串1 < 子串2
// 返回  0 表示 子串1 == 子串2
// 返回  1 表示 子串1 > 子串2
func compareSubstrings[T integerHash](h *doubleHash, arr []T, l1, r1, l2, r2 int) int {
	len1 := r1 - l1 + 1
	len2 := r2 - l2 + 1
	minLen := min(len1, len2)

	// 双开区间二分 (l, r)
	left := 0          // l 维护“匹配成功”的边界，初始 0 必定成功
	right := minLen + 1 // r 维护“匹配失败”的边界，初始 minLen+1 必定失败
	
	for left + 1 < right {
		mid := left + (right - left) >> 1
		
		// 检查长度为 mid 的前缀哈希是否相等
		if h.getHash(l1, l1 + mid - 1) == h.getHash(l2, l2 + mid - 1) {
			left = mid
		} else {
			right = mid
		}
	}
	
	lcp := left

	if lcp == len1 && lcp == len2 {
		return 0
	}

	if lcp == len1 {
		return -1
	}

	if lcp == len2 {
		return 1
	}

	if arr[l1 + lcp] < arr[l2 + lcp] {
		return -1
	}
	return 1
}
