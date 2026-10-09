package Data_Structure

import (
	"math/bits"
	"sort"
	"slices"
)

// ------- 线段树 ------- //

type info = int

type seg []struct {
	l, r int
	val  info
}

// 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
func (seg) mergeInfo(a, b info) info {
	return max(a, b)
}

func (t seg) maintain(o int) {
	t[o].val = t.mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) build(a []info, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 调用时 o=1  0<=i<=n-1
func (t seg) update(o, i int, val info) {
	if t[o].l == t[o].r {
		t[o].val = t.mergeInfo(t[o].val, val) // 直接覆盖 t[o].val = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t seg) query(o, l, r int) info {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lRes := t.query(o<<1, l, r)
	rRes := t.query(o<<1|1, l, r)
	return t.mergeInfo(lRes, rRes)
}

// 线段树二分：返回 [l,r] 内第一个满足 f 的下标，如果不存在，返回 -1
// 例如查询 [l,r] 内第一个大于等于 target 的元素下标，需要线段树维护区间最大值
//     t.findFirst(1, l, r, func(nodeMax int) bool { return nodeMax >= target })
// 调用时 o=1
func (t seg) findFirst(o, l, r int, f func(int) bool) int {
	if t[o].l > r || t[o].r < l || !f(t[o].val) {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	idx := t.findFirst(o<<1, l, r, f)
	if idx < 0 {
		idx = t.findFirst(o<<1|1, l, r, f)
	}
	return idx
}

// 线段树二分：返回 [l,r] 内最后一个满足 f 的下标，如果不存在，返回 -1
// 例如查询 [l,r] 内最后一个小于等于 target 的元素下标，需要线段树维护区间最小值
//     t.findLast(1, l, r, func(nodeMin int) bool { return nodeMin <= target })
// 调用时 o=1
func (t seg) findLast(o, l, r int, f func(int) bool) int {
	if t[o].l > r || t[o].r < l || !f(t[o].val) {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	idx := t.findLast(o<<1|1, l, r, f)
	if idx < 0 {
		idx = t.findLast(o<<1, l, r, f)
	}
	return idx
}

func newSegmentTree(a []info) seg {
	n := len(a)
	if n == 0 {
		panic("slice can't be empty")
	}
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func newSegmentTreeBySize(n int, defaultVal info) seg {
	a := make([]info, n)
	for i := range a {
		a[i] = defaultVal
	}
	return newSegmentTree(a)
}


// ------- 动态开点线段树 ------- //

// type info = int

const defaultVal info = 0

// 动态开点线段树，默认哨兵节点避免判空
var emptyStNode = &stNode{val: defaultVal}

func init() {
	emptyStNode.lo = emptyStNode
	emptyStNode.ro = emptyStNode
}

type stNode struct {
	lo, ro *stNode
	l, r   int
	val    info
}

func (stNode) mergeInfo(a, b info) info {
	return max(a, b) // 按照所需合并
}

func (o *stNode) maintain() {
	o.val = o.mergeInfo(o.lo.val, o.ro.val)
}

func (o *stNode) update(i int, val info) {
	if o.l == o.r {
		o.val = o.mergeInfo(o.val, val)
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == emptyStNode {
			o.lo = &stNode{lo: emptyStNode, ro: emptyStNode, l: o.l, r: m, val: defaultVal}
		}
		o.lo.update(i, val)
	} else {
		if o.ro == emptyStNode {
			o.ro = &stNode{lo: emptyStNode, ro: emptyStNode, l: m + 1, r: o.r, val: defaultVal}
		}
		o.ro.update(i, val)
	}
	o.maintain()
}

func (o *stNode) query(l, r int) info {
	if o == emptyStNode || l > o.r || r < o.l {
		return defaultVal
	}
	if l <= o.l && o.r <= r {
		return o.val
	}
	return o.mergeInfo(o.lo.query(l, r), o.ro.query(l, r))
}

// 线段树合并
func (o *stNode) merge(b *stNode) *stNode {
	if o == emptyStNode || o == nil {
		return b
	}
	if b == emptyStNode || b == nil {
		return o
	}
	if o.l == o.r {
		o.val += b.val // 按需调整逻辑
		return o
	}
	o.lo = o.lo.merge(b.lo)
	o.ro = o.ro.merge(b.ro)
	o.maintain()
	return o
}

// 线段树分裂：将区间 [l,r] 从 o 中分离到 b 上
func (o *stNode) split(b *stNode, l, r int) (*stNode, *stNode) {
	if o == emptyStNode || l > o.r || r < o.l {
		return o, emptyStNode
	}
	if l <= o.l && o.r <= r {
		return emptyStNode, o
	}
	if b == emptyStNode || b == nil {
		b = &stNode{lo: emptyStNode, ro: emptyStNode, l: o.l, r: o.r, val: defaultVal}
	}
	o.lo, b.lo = o.lo.split(b.lo, l, r)
	o.ro, b.ro = o.ro.split(b.ro, l, r)
	o.maintain()
	b.maintain()
	return o, b
}

func newStRoot(l, r int) *stNode {
	return &stNode{lo: emptyStNode, ro: emptyStNode, l: l, r: r, val: defaultVal}
}


// ------- 可持久化线段树（主席树） ------- //

type pInfo struct{
    cnt int
    sum int
}

type pNode struct {
    lo, ro *pNode
	l, r   int
	pInfo
}

func (pNode) mergeInfo(l, r pInfo) pInfo {
	return pInfo{l.cnt + r.cnt, l.sum + r.sum}
}

func (pNode) diffInfo(l, r pInfo) pInfo {
	return pInfo{r.cnt - l.cnt, r.sum - l.sum}
}

func (o *pNode) maintain() {
	o.pInfo = o.mergeInfo(o.lo.pInfo, o.ro.pInfo)
}

func buildPst(l, r int) *pNode {
	o := &pNode{l: l, r: r}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPst(l, m)
	o.ro = buildPst(m+1, r)
	o.maintain()
	return o
}

func (o pNode) update(i int, val int) *pNode {
	if o.l == o.r {
		o.cnt++
		o.sum += val
		return &o
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		o.lo = o.lo.update(i, val)
	} else {
		o.ro = o.ro.update(i, val)
	}
	o.maintain()
	return &o
}

func (o *pNode) query(old *pNode, ql, qr int) pInfo {
	if ql <= o.l && o.r <= qr {
		return o.diffInfo(old.pInfo, o.pInfo)
	}
	m := (o.l + o.r) >> 1
	if qr <= m {
		return o.lo.query(old.lo, ql, qr)
	}
	if m < ql {
		return o.ro.query(old.ro, ql, qr)
	}
	lRes := o.lo.query(old.lo, ql, qr)
	rRes := o.ro.query(old.ro, ql, qr)
	return o.mergeInfo(lRes, rRes)
}

// k 从 0 开始
func (o *pNode) kth(old *pNode, k int) int {
	if o.l == o.r {
		return o.l
	}
	cntL := o.lo.cnt - old.lo.cnt
	if k < cntL {
		return o.lo.kth(old.lo, k)
	}
	return o.ro.kth(old.ro, k-cntL)
}

func newPst(a []int) ([]*pNode, []int) {
	unique := slices.Clone(a)
	slices.Sort(unique)
	unique = slices.Compact(unique)

	t := make([]*pNode, len(a)+1)
	t[0] = buildPst(0, len(unique)-1)
	for i, v := range a {
		j := sort.SearchInts(unique, v)
		t[i+1] = t[i].update(j, v)
	}
	return t, unique
}
