package Data_Structure

// ------- 李超线段树 (浮点数版) ------- //
// 初始化：tree := newLiChaoSegmentTree(n)

const inf = 1e18
const eps = 1e-9

type Line struct {
	k, b float64
	id   int
}

func newLineDefault() Line {
	return Line{k: 0, b: inf, id: 0}
	// 最大值：return Line{k: 0, b: -inf, id: 0}
}

func (l Line) calc(x int) float64 {
	return l.k*float64(x) + l.b
}

type LiChaoNode struct {
	line Line
}

type LiChaoSegmentTree struct {
	n    int
	tree []LiChaoNode
}

func newLiChaoSegmentTree(rangeSize int) *LiChaoSegmentTree {
	st := &LiChaoSegmentTree{
		n:    rangeSize,
		tree: make([]LiChaoNode, 4*(rangeSize+1)),
	}
	for i := range st.tree {
		st.tree[i].line = newLineDefault()
	}
	return st
}

func (st *LiChaoSegmentTree) better(a, b Line, x int) Line {
	if a.id == 0 {
		return b
	}
	if b.id == 0 {
		return a
	}
	va, vb := a.calc(x), b.calc(x)
	if va-vb < -eps {
		return a
	}
	// 最大值：if va - vb > eps { return a }

	if vb-va < -eps {
		return b
	}
	// 最大值：if vb - va > eps { return b }

	if a.id < b.id {
		return a
	}
	return b
}

func (st *LiChaoSegmentTree) updateNode(node, l, r int, newLine Line) {
	if l > r {
		return
	}
	oldLine := st.tree[node].line
	mid := (l + r) >> 1
	betterAtMid := newLine.calc(mid) < oldLine.calc(mid)-eps
	// 最大值：betterAtMid := newLine.calc(mid) > oldLine.calc(mid)+eps

	if betterAtMid {
		st.tree[node].line, newLine = newLine, st.tree[node].line
	}
	if l == r {
		return
	}
	betterAtLeft := newLine.calc(l) < oldLine.calc(l)-eps
	// 最大值：betterAtLeft := newLine.calc(l) > oldLine.calc(l)+eps
	betterAtRight := newLine.calc(r) < oldLine.calc(r)-eps
	// 最大值：betterAtRight := newLine.calc(r) > oldLine.calc(r)+eps

	if betterAtLeft != betterAtMid {
		st.updateNode(node<<1, l, mid, newLine)
	} else if betterAtRight != betterAtMid {
		st.updateNode(node<<1|1, mid+1, r, newLine)
	}
}

func (st *LiChaoSegmentTree) insertLine(node, l, r, ql, qr int, line Line) {
	if l > qr || r < ql {
		return
	}
	if ql <= l && r <= qr {
		st.updateNode(node, l, r, line)
		return
	}
	mid := (l + r) >> 1
	st.insertLine(node<<1, l, mid, ql, qr, line)
	st.insertLine(node<<1|1, mid+1, r, ql, qr, line)
}

func (st *LiChaoSegmentTree) queryNode(node, l, r, x int) Line {
	if l == r {
		return st.tree[node].line
	}
	mid := (l + r) >> 1
	res := st.tree[node].line
	var childRes Line
	if x <= mid {
		childRes = st.queryNode(node<<1, l, mid, x)
	} else {
		childRes = st.queryNode(node<<1|1, mid+1, r, x)
	}
	return st.better(res, childRes, x)
}

// ----- 外接接口 ----- //
func (st *LiChaoSegmentTree) insert(l, r int, line Line) {
	st.insertLine(1, 0, st.n, l, r, line)
}

func (st *LiChaoSegmentTree) query(x int) float64 {
	line := st.queryNode(1, 0, st.n, x)
	if line.id == 0 {
		return inf
	}
	// 最大值：if line.id == 0 { return -inf }
	return line.calc(x)
}

// ------- 动态开点 李超线段树 (浮点数版) ------- //
type DynLiChaoNode struct {
	lc, rc *DynLiChaoNode
	line   Line
}

type DynamicLiChaoTree struct {
	minX, maxX int
	root       *DynLiChaoNode
}

func NewDynamicLiChaoTree(minX, maxX int) *DynamicLiChaoTree {
	return &DynamicLiChaoTree{
		minX: minX,
		maxX: maxX,
		root: nil,
	}
}

func (st *DynamicLiChaoTree) better(a, b Line, x int) bool {
	if a.id == 0 {
		return false
	}
	if b.id == 0 {
		return true
	}
	va, vb := a.calc(x), b.calc(x)
	if va-vb < -eps {
		return true
	}
	// 最大值：if va - vb > eps { return true }
	if vb-va < -eps {
		return false
	}
	// 最大值：if vb - va > eps { return false }
	return a.id < b.id
}

func (st *DynamicLiChaoTree) updateNode(p **DynLiChaoNode, l, r int, newLine Line) {
	if *p == nil {
		*p = &DynLiChaoNode{line: newLine}
		return
	}
	mid := (l + r) >> 1
	leftBetter := st.better(newLine, (*p).line, l)
	midBetter := st.better(newLine, (*p).line, mid)

	if midBetter {
		(*p).line, newLine = newLine, (*p).line // swap
		midBetter = st.better(newLine, (*p).line, mid)
		leftBetter = st.better(newLine, (*p).line, l)
	}

	if l == r {
		return
	}
	if leftBetter != midBetter {
		st.updateNode(&(*p).lc, l, mid, newLine)
	} else {
		st.updateNode(&(*p).rc, mid+1, r, newLine)
	}
}

func (st *DynamicLiChaoTree) insertSegment(p **DynLiChaoNode, l, r, ql, qr int, line Line) {
	if *p == nil {
		*p = &DynLiChaoNode{line: newLineDefault()}
	}
	if ql <= l && r <= qr {
		st.updateNode(p, l, r, line)
		return
	}
	mid := (l + r) >> 1
	if ql <= mid {
		st.insertSegment(&(*p).lc, l, mid, ql, qr, line)
	}
	if qr > mid {
		st.insertSegment(&(*p).rc, mid+1, r, ql, qr, line)
	}
}

func (st *DynamicLiChaoTree) queryNode(p *DynLiChaoNode, l, r, x int) float64 {
	if p == nil {
		return inf
		// 最大值：return -inf
	}
	res := float64(inf)
	// 最大值：res := float64(-inf)
	if p.line.id != 0 {
		res = p.line.calc(x)
	}
	if l == r {
		return res
	}
	mid := (l + r) >> 1
	if x <= mid {
		return min(res, st.queryNode(p.lc, l, mid, x))
	}
	return min(res, st.queryNode(p.rc, mid+1, r, x))
	// 最大值：return max(res, ...)
}

func (st *DynamicLiChaoTree) mergeNodes(p, q *DynLiChaoNode, l, r int) *DynLiChaoNode {
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	if l == r {
		if st.better(q.line, p.line, l) {
			p.line = q.line
		}
		return p
	}
	mid := (l + r) >> 1
	p.lc = st.mergeNodes(p.lc, q.lc, l, mid)
	p.rc = st.mergeNodes(p.rc, q.rc, mid+1, r)
	if q.line.id != 0 {
		st.updateNode(&p, l, r, q.line)
	}
	return p
}

// ----- 外部接口 ----- //
func (st *DynamicLiChaoTree) insert(l, r int, line Line) {
	if l > r {
		return
	}
	l = max(l, st.minX)
	r = min(r, st.maxX)
	if l <= r {
		st.insertSegment(&st.root, st.minX, st.maxX, l, r, line)
	}
}

func (st *DynamicLiChaoTree) query(x int) float64 {
	return st.queryNode(st.root, st.minX, st.maxX, x)
}

func (st *DynamicLiChaoTree) merge(other *DynamicLiChaoTree) {
	if other.root == nil {
		return
	}
	st.root = st.mergeNodes(st.root, other.root, st.minX, st.maxX)
	other.root = nil
}


// ------- 可持久化李超线段树 (浮点数版) ------- //
// 初始化：tree := NewPersistentLiChaoTree(minX, maxX)

type PersistLiChaoNode struct {
	lc, rc *PersistLiChaoNode
	line   Line
}

type PersistentLiChaoTree struct {
	minX, maxX int
	history    []*PersistLiChaoNode // 维护历史版本根节点
}

func NewPersistentLiChaoTree(minX, maxX int) *PersistentLiChaoTree {
	return &PersistentLiChaoTree{
		minX:    minX,
		maxX:    maxX,
		history: []*PersistLiChaoNode{nil}, // 初始第 0 个版本为空树 (nil)
	}
}

func (st *PersistentLiChaoTree) better(a, b Line, x int) bool {
	if a.id == 0 {
		return false
	}
	if b.id == 0 {
		return true
	}
	va, vb := a.calc(x), b.calc(x)
	if va-vb < -eps {
		return true
	}
	// 最大值：if va - vb > eps { return true }
	if vb-va < -eps {
		return false
	}
	// 最大值：if vb - va > eps { return false }
	return a.id < b.id
}

func (st *PersistentLiChaoTree) insertNode(p *PersistLiChaoNode, l, r int, newLine Line) *PersistLiChaoNode {
	// 可持久化核心：不修改原节点，而是创建一个新节点（拷贝旧节点的状态）
	curr := &PersistLiChaoNode{}
	if p != nil {
		*curr = *p // 浅拷贝左右儿子指针和当前线段
	} else {
		curr.line = newLineDefault()
	}

	mid := (l + r) >> 1
	leftBetter := st.better(newLine, curr.line, l)
	midBetter := st.better(newLine, curr.line, mid)

	if midBetter {
		curr.line, newLine = newLine, curr.line // swap
		midBetter = st.better(newLine, curr.line, mid)
		leftBetter = st.better(newLine, curr.line, l)
	}

	if l == r {
		return curr
	}

	// 下放处于劣势的线段到对应的子树，并更新新节点的子节点指针
	if leftBetter != midBetter {
		curr.lc = st.insertNode(curr.lc, l, mid, newLine)
	} else {
		curr.rc = st.insertNode(curr.rc, mid+1, r, newLine)
	}

	return curr
}

func (st *PersistentLiChaoTree) queryNode(p *PersistLiChaoNode, l, r, x int) float64 {
	if p == nil {
		return inf
		// 最大值：return -inf
	}
	res := float64(inf)
	// 最大值：res := float64(-inf)
	if p.line.id != 0 {
		res = p.line.calc(x)
	}
	if l == r {
		return res
	}
	mid := (l + r) >> 1
	if x <= mid {
		return min(res, st.queryNode(p.lc, l, mid, x))
	}
	return min(res, st.queryNode(p.rc, mid+1, r, x))
	// 最大值：return max(res, ...)
}

// ----- 外部接口 ----- //

// insert 插入一条全局直线，生成一个新版本
// 时空复杂度：O(log C)
func (st *PersistentLiChaoTree) insert(line Line) {
	root := st.history[len(st.history)-1]
	newRoot := st.insertNode(root, st.minX, st.maxX, line)
	st.history = append(st.history, newRoot)
}

// query 查询点 x 的最值（在最新版本上查询）
// 时间复杂度：O(log C)
func (st *PersistentLiChaoTree) query(x int) float64 {
	root := st.history[len(st.history)-1]
	return st.queryNode(root, st.minX, st.maxX, x)
}

// rollBack 返回上一个版本（撤销最后一次插入的直线）
// 时间复杂度：O(1)
func (st *PersistentLiChaoTree) rollBack() {
	if len(st.history) > 1 {
		st.history = st.history[:len(st.history)-1]
	}
}

func (st *PersistentLiChaoTree) clear() {
	st.history = []*PersistLiChaoNode{nil}
}
