package Data_Structure

import (
	"math/bits"
)

// ------- Lazy线段树 ------- //

const todoInit = 0
// type info = int

type lazySeg []struct {
	l, r int
	sum  info
	todo int
}

func (lazySeg) mergeInfo(a, b info) info {
	return a + b
}

func (lazySeg) mergeTodo(a, b int) int {
	return a + b
}

func (t lazySeg) apply(o int, f int) {
	cur := &t[o]
	cur.sum += f * (cur.r - cur.l + 1)
	cur.todo = t.mergeTodo(f, cur.todo)
}

func (t lazySeg) maintain(o int) {
	t[o].sum = t.mergeInfo(t[o<<1].sum, t[o<<1|1].sum)
}

func (t lazySeg) spread(o int) {
	f := t[o].todo
	if f == todoInit {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit
}

func (t lazySeg) build(a []info, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].sum = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t lazySeg) update(o, l, r int, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t lazySeg) query(o, l, r int) info {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
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

func newLazySegmentTree(a []info) lazySeg {
	n := len(a)
	if n == 0 {
		panic("slice can't be empty")
	}
	t := make(lazySeg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func newLazySegmentTreeBySize(n int, defaultVal info) lazySeg {
	a := make([]info, n)
	for i := range a {
		a[i] = defaultVal
	}
	return newLazySegmentTree(a)
}


// ------- 动态开点 Lazy线段树 ------- //

type lazyInfo = int
type lazyTodo = int

const lazyDefaultVal lazyInfo = 0
const lazyDefaultTodo lazyTodo = 0

// 动态开点 Lazy 线段树，带哨兵机制节省空间
var emptyLazyNode = &lazyNode{sum: lazyDefaultVal, todo: lazyDefaultTodo}

func init() {
	emptyLazyNode.lo = emptyLazyNode
	emptyLazyNode.ro = emptyLazyNode
}

type lazyNode struct {
	lo, ro *lazyNode
	l, r   int
	sum    lazyInfo
	todo   lazyTodo
}

func (lazyNode) mergeInfo(a, b lazyInfo) lazyInfo {
	return a + b
}

func (lazyNode) mergeTodo(a, b lazyTodo) lazyTodo {
	return a + b
}

func (o *lazyNode) apply(f lazyTodo) {
	o.sum += lazyInfo(o.r-o.l+1) * f
	o.todo = o.mergeTodo(f, o.todo)
}

func (o *lazyNode) maintain() {
	o.sum = o.mergeInfo(o.lo.sum, o.ro.sum)
}

func (o *lazyNode) spread() {
	m := (o.l + o.r) >> 1
	if o.lo == emptyLazyNode {
		o.lo = &lazyNode{lo: emptyLazyNode, ro: emptyLazyNode, l: o.l, r: m, sum: lazyDefaultVal, todo: lazyDefaultTodo}
	}
	if o.ro == emptyLazyNode {
		o.ro = &lazyNode{lo: emptyLazyNode, ro: emptyLazyNode, l: m + 1, r: o.r, sum: lazyDefaultVal, todo: lazyDefaultTodo}
	}
	if f := o.todo; f != lazyDefaultTodo {
		o.lo.apply(f)
		o.ro.apply(f)
		o.todo = lazyDefaultTodo
	}
}

func (o *lazyNode) update(l, r int, add lazyTodo) {
	if l <= o.l && o.r <= r {
		o.apply(add)
		return
	}
	o.spread()
	m := (o.l + o.r) >> 1
	if l <= m {
		o.lo.update(l, r, add)
	}
	if m < r {
		o.ro.update(l, r, add)
	}
	o.maintain()
}

func (o *lazyNode) query(l, r int) lazyInfo {
	if o == emptyLazyNode || l > o.r || r < o.l {
		return lazyDefaultVal
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	o.spread()
	return o.mergeInfo(o.lo.query(l, r), o.ro.query(l, r))
}

// 线段树合并：将 b 合并到 o 上，返回合并后的根。
// 合并时直接累加 sum 和 todo，保持 lazy 语义，不调用 maintain。
func (o *lazyNode) merge(b *lazyNode) *lazyNode {
	if o == emptyLazyNode {
		return b
	}
	if b == emptyLazyNode {
		return o
	}
	if o.l == o.r {
		o.sum += b.sum
		o.todo = o.mergeTodo(o.todo, b.todo)
		return o
	}
	o.lo = o.lo.merge(b.lo)
	o.ro = o.ro.merge(b.ro)
	o.sum += b.sum
	o.todo = o.mergeTodo(o.todo, b.todo)
	return o
}

// 线段树分裂：将区间 [l,r] 从 o 中分离到 b 上，返回 (o, b)。
// 通常要求 b 为空树（emptyLazyNode）。
func (o *lazyNode) split(b *lazyNode, l, r int) (*lazyNode, *lazyNode) {
	if o == emptyLazyNode || l > o.r || r < o.l {
		return o, emptyLazyNode
	}
	if l <= o.l && o.r <= r {
		return emptyLazyNode, o
	}

	o.spread()

	if b == emptyLazyNode {
		b = &lazyNode{
			lo:   emptyLazyNode,
			ro:   emptyLazyNode,
			l:    o.l,
			r:    o.r,
			sum:  lazyDefaultVal,
			todo: lazyDefaultTodo,
		}
	}

	o.lo, b.lo = o.lo.split(b.lo, l, r)
	o.ro, b.ro = o.ro.split(b.ro, l, r)

	o.maintain()
	b.maintain()
	return o, b
}

func newLazyDynRoot(l, r int) *lazyNode {
	return &lazyNode{lo: emptyLazyNode, ro: emptyLazyNode, l: l, r: r, sum: lazyDefaultVal, todo: lazyDefaultTodo}
}