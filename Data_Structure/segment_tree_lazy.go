package Data_Structure

import (
	"math/bits"
)

// ------- Lazy线段树 ------- //
type LazyT = int
type LazyF = int

const TODO_INIT LazyF = 0 // **根据题目修改** 懒标记初始值

type LazyNode struct {
	val  LazyT
	todo LazyF
}

type LazySegmentTree struct {
	n    int
	tree []LazyNode
}

// 合并两个 val
func (st *LazySegmentTree) mergeVal(a, b LazyT) LazyT {
	return a + b // **根据题目修改**
}

// 合并两个懒标记
func (st *LazySegmentTree) mergeTodo(a, b LazyF) LazyF {
	return a + b // **根据题目修改**
}

// 把懒标记作用到 node 子树（本例为区间加）
func (st *LazySegmentTree) apply(node, l, r int, todo LazyF) {
	// 计算 tree[node] 区间的整体变化
	// 如果是 max or min 写成 st.tree[node].val += todo;
	st.tree[node].val += todo * LazyT(r-l+1) // **根据题目修改**
	st.tree[node].todo = st.mergeTodo(todo, st.tree[node].todo)
}

// 把当前节点的懒标记下传给左右儿子
func (st *LazySegmentTree) spread(node, l, r int) {
	todo := st.tree[node].todo
	if todo == TODO_INIT { // 没有需要下传的信息
		return
	}
	m := (l + r) >> 1
	st.apply(node<<1, l, m, todo)
	st.apply(node<<1|1, m+1, r, todo)
	st.tree[node].todo = TODO_INIT // 下传完毕
}

// 合并左右儿子的 val 到当前节点的 val
func (st *LazySegmentTree) maintain(node int) {
	st.tree[node].val = st.mergeVal(st.tree[node<<1].val, st.tree[node<<1|1].val)
}

// 用 a 初始化线段树
// 时间复杂度 O(n)
func (st *LazySegmentTree) build(a []LazyT, node, l, r int) {
	st.tree[node].todo = TODO_INIT
	if l == r { // 叶子
		st.tree[node].val = a[l] // 初始化叶节点的值
		return
	}
	m := (l + r) >> 1
	st.build(a, node<<1, l, m)     // 初始化左子树
	st.build(a, node<<1|1, m+1, r) // 初始化右子树
	st.maintain(node)
}

func (st *LazySegmentTree) updateNode(node, l, r, ql, qr int, f LazyF) {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		st.apply(node, l, r, f)
		return
	}
	st.spread(node, l, r)
	m := (l + r) >> 1
	if ql <= m { // 更新左子树
		st.updateNode(node<<1, l, m, ql, qr, f)
	}
	if qr > m { // 更新右子树
		st.updateNode(node<<1|1, m+1, r, ql, qr, f)
	}
	st.maintain(node)
}

func (st *LazySegmentTree) queryNode(node, l, r, ql, qr int) LazyT {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return st.tree[node].val
	}
	st.spread(node, l, r)
	m := (l + r) >> 1
	if qr <= m { // [ql, qr] 在左子树
		return st.queryNode(node<<1, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 在右子树
		return st.queryNode(node<<1|1, m+1, r, ql, qr)
	}
	lRes := st.queryNode(node<<1, l, m, ql, qr)
	rRes := st.queryNode(node<<1|1, m+1, r, ql, qr)
	return st.mergeVal(lRes, rRes)
}

// 私有递归：查找第一个满足条件的下标
func (st *LazySegmentTree) findFirstNode(node, l, r, ql, qr int, f func(LazyT) bool) int {
	if r < ql || l > qr {
		return -1
	}
	if !f(st.tree[node].val) {
		return -1
	}
	if l == r {
		return l
	}
	st.spread(node, l, r) // 下传懒标记
	m := (l + r) >> 1
	leftRes := st.findFirstNode(node<<1, l, m, ql, qr, f)
	if leftRes != -1 {
		return leftRes
	}
	return st.findFirstNode(node<<1|1, m+1, r, ql, qr, f)
}

// 私有递归：查找最后一个满足条件的下标
func (st *LazySegmentTree) findLastNode(node, l, r, ql, qr int, f func(LazyT) bool) int {
	if r < ql || l > qr {
		return -1
	}
	if !f(st.tree[node].val) {
		return -1
	}
	if l == r {
		return l
	}
	st.spread(node, l, r)
	m := (l + r) >> 1
	rightRes := st.findLastNode(node<<1|1, m+1, r, ql, qr, f)
	if rightRes != -1 {
		return rightRes
	}
	return st.findLastNode(node<<1, l, m, ql, qr, f)
}

func NewLazySegmentTreeBySize(n int, initVal LazyT) *LazySegmentTree {
	a := make([]LazyT, n)
	for i := range a {
		a[i] = initVal
	}
	return NewLazySegmentTree(a)
}

func NewLazySegmentTree(a []LazyT) *LazySegmentTree {
	n := len(a)
	treeSize := 2 << bits.Len(uint(n-1))
	st := &LazySegmentTree{
		n:    n,
		tree: make([]LazyNode, treeSize),
	}
	if n > 0 {
		st.build(a, 1, 0, n-1)
	}
	return st
}

// ----- 外部接口 ----- //
func (st *LazySegmentTree) update(ql, qr int, f LazyF) {
	st.updateNode(1, 0, st.n-1, ql, qr, f)
}

func (st *LazySegmentTree) query(ql, qr int) LazyT {
	return st.queryNode(1, 0, st.n-1, ql, qr)
}

func (st *LazySegmentTree) findfirst(ql, qr int, f func(LazyT) bool) int {
	return st.findFirstNode(1, 0, st.n-1, ql, qr, f)
}

func (st *LazySegmentTree) findlast(ql, qr int, f func(LazyT) bool) int {
	return st.findLastNode(1, 0, st.n-1, ql, qr, f)
}

// ------- 动态开点 Lazy线段树 ------- //
type DynLazyT = int
type DynLazyF = int

type DynLazySegNode struct {
	left, right *DynLazySegNode
	l, r        int
	val         DynLazyT
	todo        DynLazyF
}

type DynamicLazySegmentTree struct {
	root       *DynLazySegNode
	empty      *DynLazySegNode // 同上，安全修正为实例绑定
	minX, maxX int             // 值域范围 [minX, maxX]
	valDefault DynLazyT        // 值默认值（如区间和时为0）
	todoInit   DynLazyF        // 懒标记初始值（如无标记时为0）
}

func NewDynamicLazySegmentTree(minX, maxX int, valDef DynLazyT, todoDef DynLazyF) *DynamicLazySegmentTree {
	empty := &DynLazySegNode{l: 0, r: 0, val: valDef, todo: todoDef}
	empty.left = empty
	empty.right = empty

	root := &DynLazySegNode{l: minX, r: maxX, val: valDef, todo: todoDef, left: empty, right: empty}

	return &DynamicLazySegmentTree{
		root:       root,
		empty:      empty,
		minX:       minX,
		maxX:       maxX,
		valDefault: valDef,
		todoInit:   todoDef,
	}
}

func (st *DynamicLazySegmentTree) mergeVal(a, b DynLazyT) DynLazyT {
	return a + b // 区间求和；可改为 max, min 等
}

func (st *DynamicLazySegmentTree) mergeTodo(a, b DynLazyF) DynLazyF {
	return a + b // 懒标记累加；可改为 max, min 等
}

func (st *DynamicLazySegmentTree) apply(o *DynLazySegNode, f DynLazyF) {
	o.val += f * DynLazyT(o.r-o.l+1) // ** 根据题目修改 ** //
	o.todo = st.mergeTodo(o.todo, f)
}

func (st *DynamicLazySegmentTree) maintain(o *DynLazySegNode) {
	o.val = st.mergeVal(o.left.val, o.right.val)
}

func (st *DynamicLazySegmentTree) spread(o *DynLazySegNode) {
	if o.todo == st.todoInit {
		return
	}
	m := (o.l + o.r) >> 1
	// 动态创建左子节点
	if o.left == st.empty {
		o.left = &DynLazySegNode{l: o.l, r: m, val: st.valDefault, todo: st.todoInit, left: st.empty, right: st.empty}
	}
	// 动态创建右子节点
	if o.right == st.empty {
		o.right = &DynLazySegNode{l: m + 1, r: o.r, val: st.valDefault, todo: st.todoInit, left: st.empty, right: st.empty}
	}
	f := o.todo
	st.apply(o.left, f)
	st.apply(o.right, f)
	o.todo = st.todoInit
}

func (st *DynamicLazySegmentTree) updateNode(o *DynLazySegNode, ql, qr int, f DynLazyF) {
	if ql <= o.l && o.r <= qr {
		st.apply(o, f)
		return
	}
	st.spread(o)
	m := (o.l + o.r) >> 1
	if ql <= m {
		st.updateNode(o.left, ql, qr, f)
	}
	if qr > m {
		st.updateNode(o.right, ql, qr, f)
	}
	st.maintain(o)
}

func (st *DynamicLazySegmentTree) queryNode(o *DynLazySegNode, ql, qr int) DynLazyT {
	if o == st.empty || ql > o.r || qr < o.l {
		return st.valDefault
	}
	if ql <= o.l && o.r <= qr {
		return o.val
	}
	st.spread(o)
	return st.mergeVal(st.queryNode(o.left, ql, qr), st.queryNode(o.right, ql, qr))
}

func (st *DynamicLazySegmentTree) findFirstNode(o *DynLazySegNode, ql, qr int, f func(DynLazyT) bool) int {
	if o == st.empty || o.r < ql || o.l > qr || !f(o.val) {
		return -1
	}
	if o.l == o.r {
		return o.l
	}
	st.spread(o)
	leftRes := st.findFirstNode(o.left, ql, qr, f)
	if leftRes != -1 {
		return leftRes
	}
	return st.findFirstNode(o.right, ql, qr, f)
}

func (st *DynamicLazySegmentTree) findLastNode(o *DynLazySegNode, ql, qr int, f func(DynLazyT) bool) int {
	if o == st.empty || o.r < ql || o.l > qr || !f(o.val) {
		return -1
	}
	if o.l == o.r {
		return o.l
	}
	st.spread(o)
	rightRes := st.findLastNode(o.right, ql, qr, f)
	if rightRes != -1 {
		return rightRes
	}
	return st.findLastNode(o.left, ql, qr, f)
}

func (st *DynamicLazySegmentTree) mergeNodes(a, b *DynLazySegNode) *DynLazySegNode {
	if a == st.empty {
		return b
	}
	if b == st.empty {
		return a
	}
	st.spread(a)
	st.spread(b)
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
func (st *DynamicLazySegmentTree) update(ql, qr int, f DynLazyF) {
	if ql > qr || ql > st.maxX || qr < st.minX {
		return
	}
	ql = max(ql, st.minX)
	qr = min(qr, st.maxX)
	st.updateNode(st.root, ql, qr, f)
}

func (st *DynamicLazySegmentTree) query(ql, qr int) DynLazyT {
	if ql > qr || ql > st.maxX || qr < st.minX {
		return st.valDefault
	}
	ql = max(ql, st.minX)
	qr = min(qr, st.maxX)
	return st.queryNode(st.root, ql, qr)
}

func (st *DynamicLazySegmentTree) get(i int) DynLazyT {
	return st.query(i, i)
}

// FindFirst 二分查找第一个满足 f 的位置
func (st *DynamicLazySegmentTree) findfirst(ql, qr int, f func(DynLazyT) bool) int {
	if ql > qr || ql > st.maxX || qr < st.minX {
		return -1
	}
	ql = max(ql, st.minX)
	qr = min(qr, st.maxX)
	return st.findFirstNode(st.root, ql, qr, f)
}

// FindLast 二分查找最后一个满足 f 的位置
func (st *DynamicLazySegmentTree) findlast(ql, qr int, f func(DynLazyT) bool) int {
	if ql > qr || ql > st.maxX || qr < st.minX {
		return -1
	}
	ql = max(ql, st.minX)
	qr = min(qr, st.maxX)
	return st.findLastNode(st.root, ql, qr, f)
}

// MergeTree 线段树合并
func (st *DynamicLazySegmentTree) merge(other *DynamicLazySegmentTree) {
	if st.minX != other.minX || st.maxX != other.maxX {
		panic("值域范围不同，无法合并")
	}
	st.root = st.mergeNodes(st.root, other.root)
}
