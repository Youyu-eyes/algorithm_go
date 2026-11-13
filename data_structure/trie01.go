package dataStructure

type trie01Node struct {
	son [2]*trie01Node
	sum int // 子树叶子数（元素个数）
	min int // 子树最小值
}

type trie01 struct{ root *trie01Node }

func newTrie01() *trie01 { return &trie01{&trie01Node{min: -inf}} }

const trieBitLen = 31 // 30 for 1e9, 63 for int64, or bits.Len(MAX_VAL)

func (t *trie01) put(v int) *trie01Node {
	o := t.root
	o.min = min(o.min, v)
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &trie01Node{}
		}
		o = o.son[b]
		o.sum++
		o.min = min(o.min, v)
	}
	return o
}

func (t *trie01) del(v int) *trie01Node {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		o = o.son[v>>i&1]
		o.sum--
	}
	return o
}

func (t *trie01) minXor(v int) (ans int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 完全图，边权为 a[u]^a[v]，求 MST
// Boruvka 算法，分治连边
func xorMST(a []int) (ans int) {
	var f func([]int, int)
	f = func(a []int, p int) {
		if a == nil || p < 0 {
			return
		}
		b := [2][]int{}
		for _, x := range a {
			k := x >> p & 1
			b[k] = append(b[k], x)
		}
		if b[0] != nil && b[1] != nil {
			// b[0] 和 b[1] 至少连一条边，贪心地连最小边
			t := newTrie01()
			for _, x := range b[0] {
				t.put(x)
			}
			minXor := inf
			for _, x := range b[1] {
				minXor = min(minXor, t.minXor(x))
			}
			ans += minXor
		}
		f(b[0], p-1)
		f(b[1], p-1)
	}
	f(a, trieBitLen-1)
	return
}

func (t *trie01) maxXor(v int) (ans int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 第 k 大异或值（0-base）
func (t *trie01) maxXorKth(v, k int) (ans int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil {
			if k < o.son[b^1].sum {
				ans |= 1 << i
				b ^= 1
			} else {
				k -= o.son[b^1].sum
			}
		}
		o = o.son[b]
	}
	return
}

// trie 中元素 <= limit 和 v 的最大异或和
func (t *trie01) maxXorLimitVal(v, limit int) (ans int) {
	o := t.root
	if o.min > limit {
		return -1
	}
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil && o.son[b^1].min <= limit {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// trie 中与 v 异或和 <= limit 的元素个数
func (t *trie01) countLimitXor(v, limit int) (ans int) {
	limit++

	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if limit >> i & 1 > 0 {
			if o.son[b] != nil {
				ans += o.son[b].sum
			}
			b ^= 1
		}
		if o.son[b] == nil {
			break
		}
		o = o.son[b]
	}
	return
}

// trie 中与 v 异或和 <= limit 的最大异或和
func (t *trie01) maxLimitXor(v, limit int) (ans int) {
	limit++

	o := t.root
	var lastO *trie01Node
	lastI, lastAns := -2, 0
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if limit >> i & 1 > 0 {
			if o.son[b] != nil {
				lastO, lastI, lastAns = o.son[b], i-1, ans
			}
			if o.son[b^1] != nil {
				ans |= 1 << i
			}
			b ^= 1
		}
		if o.son[b] == nil {
			break
		}
		o = o.son[b]
	}

	if lastI < -1 {
		return -1
	}

	o = lastO
	ans = lastAns
	for i := lastI; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 计算 trie 中所有数和 v 的异或和的 mex
func (t *trie01) mex(v int) (mex int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0 && o != nil; i-- {
		b := v >> i & 1
		if o.son[b] != nil && o.son[b].sum == 1<<i {
			mex |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 可持久化 01 字典树
// 调用别的函数时，将方法对象改成 (o *trie01Node)，删除 o := t.root
// roots := make([]*trie01Node, n+1)
// roots[0] = trie01Node{}.put(0, trieBitLen-1)
// roots[i+1] = roots[i].put(v, trieBitLen-1)
func (o trie01Node) put(v, k int) *trie01Node {
	if k < 0 {
		return &o
	}
	b := v >> k & 1
	if o.son[b] == nil {
		o.son[b] = &trie01Node{}
	}
	o.son[b] = o.son[b].put(v, k-1)
	//o.maintain()
	return &o
}