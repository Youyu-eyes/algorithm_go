package Data_Structure

import (
	"math/bits"
)

// ------- 线段树 ------- //
type SegT = int

type SegmentTree struct {
	n    int
	tree []SegT
}

func (st *SegmentTree) mergeVal(a, b SegT) SegT {
	return max(a, b) // **根据题目修改** 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
}

func (st *SegmentTree) maintain(node int) {
	st.tree[node] = st.mergeVal(st.tree[node<<1], st.tree[node<<1|1])
}

func (st *SegmentTree) build(a []SegT, node, l, r int) {
	if l == r {
		st.tree[node] = a[l]
		return
	}
	m := (l + r) >> 1
	st.build(a, node<<1, l, m)
	st.build(a, node<<1|1, m+1, r)
	st.maintain(node)
}

func (st *SegmentTree) updateNode(node, l, r, i int, val SegT) {
	if l == r {
		st.tree[node] = st.mergeVal(st.tree[node], val) // 如果想直接覆盖就改成 st.tree[node] = val;
		return
	}
	m := (l + r) >> 1
	if i <= m {
		st.updateNode(node<<1, l, m, i, val)
	} else {
		st.updateNode(node<<1|1, m+1, r, i, val)
	}
	st.maintain(node)
}

func (st *SegmentTree) queryNode(node, l, r, ql, qr int) SegT {
	if ql <= l && r <= qr {
		return st.tree[node]
	}
	m := (l + r) >> 1
	if qr <= m {
		return st.queryNode(node<<1, l, m, ql, qr)
	}
	if ql > m {
		return st.queryNode(node<<1|1, m+1, r, ql, qr)
	}
	lRes := st.queryNode(node<<1, l, m, ql, qr)
	rRes := st.queryNode(node<<1|1, m+1, r, ql, qr)
	return st.mergeVal(lRes, rRes)
}

func (st *SegmentTree) findFirstNode(node, l, r, ql, qr int, f func(SegT) bool) int {
	if r < ql || l > qr {
		return -1
	}
	if !f(st.tree[node]) {
		return -1
	}
	if l == r {
		return l
	}
	m := (l + r) >> 1
	leftRes := st.findFirstNode(node<<1, l, m, ql, qr, f)
	if leftRes != -1 {
		return leftRes
	}
	return st.findFirstNode(node<<1|1, m+1, r, ql, qr, f)
}

func (st *SegmentTree) findLastNode(node, l, r, ql, qr int, f func(SegT) bool) int {
	if r < ql || l > qr {
		return -1
	}
	if !f(st.tree[node]) {
		return -1
	}
	if l == r {
		return l
	}
	m := (l + r) >> 1
	rightRes := st.findLastNode(node<<1|1, m+1, r, ql, qr, f)
	if rightRes != -1 {
		return rightRes
	}
	return st.findLastNode(node<<1, l, m, ql, qr, f)
}

// ----- 外部接口 ----- //

// NewSegmentTreeBySize 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 init_val
func NewSegmentTreeBySize(n int, initVal SegT) *SegmentTree {
	a := make([]SegT, n)
	for i := range a {
		a[i] = initVal
	}
	return NewSegmentTree(a)
}

// NewSegmentTree 线段树维护数组 a
func NewSegmentTree(a []SegT) *SegmentTree {
	n := len(a)
	// 2 << bits.Len(uint(len(a)-1)) 等同于 C++ 中的 2 << bit_width(a.size() - 1)
	treeSize := 2 << bits.Len(uint(n-1))
	st := &SegmentTree{
		n:    n,
		tree: make([]SegT, treeSize),
	}
	if n > 0 {
		st.build(a, 1, 0, n-1)
	}
	return st
}

// Update 单点更新 将 tree[node] 改成 merge_val(tree[node], val)，需要直接覆盖要修改私有函数
func (st *SegmentTree) update(i int, val SegT) {
	st.updateNode(1, 0, st.n-1, i, val)
}

// [ql, qr] 双闭
func (st *SegmentTree) query(ql, qr int) SegT {
	return st.queryNode(1, 0, st.n-1, ql, qr)
}

func (st *SegmentTree) get(i int) SegT {
	return st.queryNode(1, 0, st.n-1, i, i)
}

// 线段树二分：返回 [l,r] 内最后一个满足 f 的下标，如果不存在，返回 -1
// 例如查询 [l,r] 内最后一个小于等于 target 的元素下标，需要线段树维护区间最小值
// t.findLast(l, r, func(nodeMin int) bool { return nodeMin <= target })

func (st *SegmentTree) findfirst(ql, qr int, f func(SegT) bool) int {
	return st.findFirstNode(1, 0, st.n-1, ql, qr, f)
}

func (st *SegmentTree) findlast(ql, qr int, f func(SegT) bool) int {
	return st.findLastNode(1, 0, st.n-1, ql, qr, f)
}

// ------- 动态开点线段树 ------- //
type DynT = int

// 线段树节点
type DynSegNode struct {
	left, right *DynSegNode
	l, r        int
	val         DynT
}

type DynamicSegmentTree struct {
	root       *DynSegNode
	empty      *DynSegNode
	minX, maxX int
	defaultVal DynT
}

func NewDynamicSegmentTree(minX, maxX int, def DynT) *DynamicSegmentTree {
	empty := &DynSegNode{l: 0, r: 0, val: def}
	empty.left = empty
	empty.right = empty

	root := &DynSegNode{l: minX, r: maxX, val: def, left: empty, right: empty}

	return &DynamicSegmentTree{
		root:       root,
		empty:      empty,
		minX:       minX,
		maxX:       maxX,
		defaultVal: def,
	}
}

func (st *DynamicSegmentTree) mergeVal(a, b DynT) DynT {
	return max(a, b) // 可改为 min, +, gcd 等
}

func (st *DynamicSegmentTree) maintain(o *DynSegNode) {
	o.val = st.mergeVal(o.left.val, o.right.val)
}

func (st *DynamicSegmentTree) updateNode(o *DynSegNode, i int, val DynT) {
	if o.l == o.r {
		o.val = st.mergeVal(o.val, val)
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.left == st.empty {
			o.left = &DynSegNode{l: o.l, r: m, val: st.defaultVal, left: st.empty, right: st.empty}
		}
		st.updateNode(o.left, i, val)
	} else {
		if o.right == st.empty {
			o.right = &DynSegNode{l: m + 1, r: o.r, val: st.defaultVal, left: st.empty, right: st.empty}
		}
		st.updateNode(o.right, i, val)
	}
	st.maintain(o)
}

func (st *DynamicSegmentTree) queryNode(o *DynSegNode, ql, qr int) DynT {
	if o == st.empty || ql > o.r || qr < o.l {
		return st.defaultVal
	}
	if ql <= o.l && o.r <= qr {
		return o.val
	}
	return st.mergeVal(st.queryNode(o.left, ql, qr), st.queryNode(o.right, ql, qr))
}

func (st *DynamicSegmentTree) findFirstNode(o *DynSegNode, ql, qr int, f func(DynT) bool) int {
	if o == st.empty || o.r < ql || o.l > qr || !f(o.val) {
		return -1
	}
	if o.l == o.r {
		return o.l
	}
	leftRes := st.findFirstNode(o.left, ql, qr, f)
	if leftRes != -1 {
		return leftRes
	}
	return st.findFirstNode(o.right, ql, qr, f)
}

func (st *DynamicSegmentTree) findLastNode(o *DynSegNode, ql, qr int, f func(DynT) bool) int {
	if o == st.empty || o.r < ql || o.l > qr || !f(o.val) {
		return -1
	}
	if o.l == o.r {
		return o.l
	}
	rightRes := st.findLastNode(o.right, ql, qr, f)
	if rightRes != -1 {
		return rightRes
	}
	return st.findLastNode(o.left, ql, qr, f)
}

func (st *DynamicSegmentTree) mergeNodes(a, b *DynSegNode) *DynSegNode {
	if a == st.empty {
		return b
	}
	if b == st.empty {
		return a
	}
	if a.l == a.r {
		a.val = st.mergeVal(a.val, b.val)
		return a
	}
	a.left = st.mergeNodes(a.left, b.left)
	a.right = st.mergeNodes(a.right, b.right)
	st.maintain(a)
	return a
}

// ----- 外部接口 ----- //
func (st *DynamicSegmentTree) update(i int, val DynT) {
	if i < st.minX || i > st.maxX {
		return
	}
	st.updateNode(st.root, i, val)
}

// [ql, qr] 双闭
func (st *DynamicSegmentTree) query(ql, qr int) DynT {
	if ql > qr || ql > st.maxX || qr < st.minX {
		return st.defaultVal
	}
	ql = max(ql, st.minX)
	qr = min(qr, st.maxX)
	return st.queryNode(st.root, ql, qr)
}

func (st *DynamicSegmentTree) get(i int) DynT {
	return st.query(i, i)
}

func (st *DynamicSegmentTree) findfirst(ql, qr int, f func(DynT) bool) int {
	if ql > qr || ql > st.maxX || qr < st.minX {
		return -1
	}
	ql = max(ql, st.minX)
	qr = min(qr, st.maxX)
	return st.findFirstNode(st.root, ql, qr, f)
}

func (st *DynamicSegmentTree) findlast(ql, qr int, f func(DynT) bool) int {
	if ql > qr || ql > st.maxX || qr < st.minX {
		return -1
	}
	ql = max(ql, st.minX)
	qr = min(qr, st.maxX)
	return st.findLastNode(st.root, ql, qr, f)
}

func (st *DynamicSegmentTree) merge(other *DynamicSegmentTree) {
	if st.minX != other.minX || st.maxX != other.maxX {
		panic("值域范围不同，无法合并")
	}
	st.root = st.mergeNodes(st.root, other.root)
}
