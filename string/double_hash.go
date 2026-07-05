package string

type IntegerHash interface {
	~int | ~int64 | ~byte | ~rune
}

const (
	Base1 = 1_000_000_021
	Mod1  = 1_000_000_007
	Base2 = 1_000_000_033
	Mod2  = 1_000_000_009
)

type DoubleHash struct {
	pref1, pref2 []int
	pow1, pow2   []int
}

type HashVal struct {
	h1 int
	h2 int
}

// 字符串双哈希模板
// 支持 int，byte，rune 类型
// O(1) 时间返回任意子数组的双哈希值 HashVal

func NewDoubleHash[T IntegerHash](arr []T) *DoubleHash {
	n := len(arr)
	h := &DoubleHash{
		pref1: make([]int, n+1),
		pref2: make([]int, n+1),
		pow1:  make([]int, n+1),
		pow2:  make([]int, n+1),
	}
	h.pow1[0], h.pow2[0] = 1, 1

	for i, val := range arr {
		v := int(val)

		h.pow1[i+1] = h.pow1[i] * Base1 % Mod1
		h.pow2[i+1] = h.pow2[i] * Base2 % Mod2

		h.pref1[i+1] = (h.pref1[i]*Base1 + v) % Mod1
		h.pref2[i+1] = (h.pref2[i]*Base2 + v) % Mod2
	}
	return h
}

func (h *DoubleHash) GetHash(l, r int) HashVal {
	if l > r || l < 0 || r >= len(h.pref1)-1 {
		return HashVal{-1, -1}
	}
	length := r - l + 1
	
	hash1 := (h.pref1[r+1] - h.pref1[l]*h.pow1[length]) % Mod1
	if hash1 < 0 {
		hash1 += Mod1
	}

	hash2 := (h.pref2[r+1] - h.pref2[l]*h.pow2[length]) % Mod2
	if hash2 < 0 {
		hash2 += Mod2
	}
	
	return HashVal{hash1, hash2}
}
