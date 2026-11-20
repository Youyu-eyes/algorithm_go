package dataStructure

const acamNodeSize = 26

type acamNode struct {
	son [acamNodeSize]*acamNode

	fail *acamNode
	last *acamNode

	idx int
	cnt int

	nodeId int
}

type gInfo struct {
	l, r int
}

type acam struct {
	patterns []string

	root *acamNode
	nodeCnt int

	g [][]int
	gi []gInfo
	dfn int

	inDeg map[*acamNode]int
}

func newAcam(patterns []string) *acam {
	ac := &acam{
		patterns: patterns,
		root: &acamNode{},
		nodeCnt: 1,
		inDeg: map[*acamNode]int{},
	}

	for i, s := range patterns {
		ac.put(s, i+1)
	}

	ac.buildFail()
	return ac
}

func (ac *acam) ord(c rune) rune { return c - 'a' }

func (ac *acam) put(s string, idx int) {
	o := ac.root
	for _, b := range s {
		b = ac.ord(b)
		if o.son[b] == nil {
			newNode := &acamNode{nodeId: ac.nodeCnt}
			o.son[b] = newNode
			ac.inDeg[newNode] = 0
			ac.nodeCnt++
		}
		o = o.son[b]
	}
	o.cnt++
	o.idx = idx
}

func (ac *acam) buildFail() {
	ac.g = make([][]int, ac.nodeCnt)
	ac.root.fail = ac.root
	ac.root.last = ac.root
	q := make([]*acamNode, 0, ac.nodeCnt)
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root
			son.last = ac.root
			ac.g[0] = append(ac.g[0], son.nodeId)
			q = append(q, son)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		for i, son := range o.son[:] {
			if son == nil {
				o.son[i] = o.fail.son[i]
				continue
			}
			son.fail = o.fail.son[i]
			ac.g[son.fail.nodeId] = append(ac.g[son.fail.nodeId], son.nodeId)
			ac.inDeg[son.fail]++
			if son.fail.idx > 0 {
				son.last = son.fail
			} else {
				son.last = son.fail.last
			}
			q = append(q, son)
		}
	}
}

// 每个串只统计一次
func (ac *acam) acCountOnce(text string) (ans int) {
	o := ac.root
	for _, b := range text {
		o = o.son[ac.ord(b)]
		for match := o; match.cnt >= 0; match = match.last {
			ans += match.cnt
			match.cnt = -1
		}
	}
	return
}

func (ac *acam) acPos(text string) [][]int {
	pos := make([][]int, len(ac.patterns))
	o := ac.root
	for i, b := range text {
		o = o.son[ac.ord(b)]
		for match := o; match != ac.root; match = match.last {
			if match.idx == 0 {
				continue
			}
			pIdx := match.idx - 1
			pos[pIdx] = append(pos[pIdx], i - len(ac.patterns[pIdx]) + 1)
		}
	}
	return pos
}

func (ac *acam) acCnt(text string) []int {
	// 【注意】调用时去掉 put 函数里的 cnt++
	o := ac.root
	for _, b := range text {
		o = o.son[ac.ord(b)]
		o.cnt++
	}
	cnt := make([]int, len(ac.patterns))
	inDeg := ac.inDeg
	q := make([]*acamNode, 0, ac.nodeCnt)
	for o, d := range inDeg {
		if d == 0 {
			q = append(q, o)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v.idx > 0 {
			cnt[v.idx-1] = v.cnt
		}
		u := v.fail
		u.cnt += v.cnt
		if inDeg[u]--; inDeg[u] == 0 {
			q = append(q, u)
		}
	}
	return cnt
}

// todo 差分树状数组 + DFN 维护模式串删除
// https://codeforces.com/contest/163/problem/E

func (ac *acam) _buildDFN(u int) {
	ac.dfn++
	ac.gi[u].l = ac.dfn
	for _, v := range ac.g[u] {
		ac._buildDFN(v)
	}
	ac.gi[u].r = ac.dfn
}

func (ac *acam) buildDFN() {
	ac.gi = make([]gInfo, ac.nodeCnt)
	ac._buildDFN(ac.root.nodeId)

}