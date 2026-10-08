package Data_Structure

import (
	"math/bits"
)

const (
	inf = 1 << 60
	eps = 1e-9
)

type line struct {
	k, b float64
	id   int
}

func newLine() line {
	return line{k: 0, b: inf, id: 0}
	// 最大值：return line{k: 0, b: -inf, id: 0}
}

func (l line) calc(x int) float64 {
	return l.k*float64(x) + l.b
}

// 李超线段树 (浮点数版)
type lcSeg []struct {
	l, r int
	line line
}

func (lcSeg) better(a, b line, x int) line {
	if a.id == 0 {
		return b
	}
	if b.id == 0 {
		return a
	}
	va, vb := a.calc(x), b.calc(x)
	if va-vb < -eps { // 最大值：if va-vb > eps
		return a
	}
	if vb-va < -eps { // 最大值：if vb-va > eps
		return b
	}
	if a.id < b.id {
		return a
	}
	return b
}

func (t lcSeg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].line = newLine()
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t lcSeg) update(o int, nl line) {
	l, r := t[o].l, t[o].r
	m := (l + r) >> 1
	old := t[o].line

	if nl.calc(m) < old.calc(m)-eps { // 最大值：nl.calc(m) > old.calc(m) + eps
		t[o].line, nl = nl, old
	}
	if l == r {
		return
	}
	if nl.calc(l) < t[o].line.calc(l)-eps { // 最大值：nl.calc(l) > t[o].line.calc(l)+eps
		t.update(o<<1, nl)
	} else if nl.calc(r) < t[o].line.calc(r)-eps { // 最大值：nl.calc(r) > t[o].line.calc(r)+eps
		t.update(o<<1|1, nl)
	}
}

func (t lcSeg) insert(o, l, r int, nl line) {
	if t[o].l > r || t[o].r < l {
		return
	}
	if l <= t[o].l && t[o].r <= r {
		t.update(o, nl)
		return
	}
	t.insert(o<<1, l, r, nl)
	t.insert(o<<1|1, l, r, nl)
}

func (t lcSeg) queryNode(o, x int) line {
	if t[o].l == t[o].r {
		return t[o].line
	}
	m := (t[o].l + t[o].r) >> 1
	var child line
	if x <= m {
		child = t.queryNode(o<<1, x)
	} else {
		child = t.queryNode(o<<1|1, x)
	}
	return t.better(t[o].line, child, x)
}

// 查询 x 坐标上的最值
func (t lcSeg) query(o, x int) float64 {
	ln := t.queryNode(o, x)
	if ln.id == 0 {
		return inf // 最大值：return -inf
	}
	return ln.calc(x)
}

// 建立范围 [0, n] 的李超线段树
func newLcSeg(n int) lcSeg {
	t := make(lcSeg, 2<<bits.Len(uint(n)))
	t.build(1, 0, n)
	return t
}


// ------- 动态开点 李超线段树 (浮点数版) ------- //

var emptyLcNode = &lcNode{line: newLine()}

func init() {
	emptyLcNode.lo = emptyLcNode
	emptyLcNode.ro = emptyLcNode
}

type lcNode struct {
	lo, ro *lcNode
	l, r   int
	line   line
}

func (o *lcNode) update(nl line) {
	l, r := o.l, o.r
	m := (l + r) >> 1
	if nl.calc(m) < o.line.calc(m)-eps { // nl.calc(m) > o.line.calc(m)+eps
		o.line, nl = nl, o.line
	}
	if l == r {
		return
	}
	if nl.calc(l) < o.line.calc(l)-eps { // 最大值：nl.calc(l) > o.line.calc(l)+eps
		if o.lo == emptyLcNode {
			o.lo = &lcNode{lo: emptyLcNode, ro: emptyLcNode, l: l, r: m, line: newLine()}
		}
		o.lo.update(nl)
	} else if nl.calc(r) < o.line.calc(r)-eps { // 最大值：nl.calc(r) > o.line.calc(r)+eps
		if o.ro == emptyLcNode {
			o.ro = &lcNode{lo: emptyLcNode, ro: emptyLcNode, l: m + 1, r: r, line: newLine()}
		}
		o.ro.update(nl)
	}
}

func (o *lcNode) insert(ql, qr int, nl line) {
	if o == emptyLcNode || ql > o.r || qr < o.l {
		return
	}
	if ql <= o.l && o.r <= qr {
		o.update(nl)
		return
	}
	m := (o.l + o.r) >> 1
	if ql <= m {
		if o.lo == emptyLcNode {
			o.lo = &lcNode{lo: emptyLcNode, ro: emptyLcNode, l: o.l, r: m, line: newLine()}
		}
		o.lo.insert(ql, qr, nl)
	}
	if m < qr {
		if o.ro == emptyLcNode {
			o.ro = &lcNode{lo: emptyLcNode, ro: emptyLcNode, l: m + 1, r: o.r, line: newLine()}
		}
		o.ro.insert(ql, qr, nl)
	}
}

func (o *lcNode) query(x int) float64 {
	if o == emptyLcNode {
		return float64(inf) // 最大值：return float64(-inf)
	}
	res := float64(inf) // 最大值：res := float64(-inf)
	if o.line.id != 0 {
		res = o.line.calc(x)
	}
	if o.l == o.r {
		return res
	}
	m := (o.l + o.r) >> 1
	var child float64
	if x <= m {
		child = o.lo.query(x)
	} else {
		child = o.ro.query(x)
	}
	if child < res { // 最大值：if child > res
		return child
	}
	return res
}

// 线段树合并：将 b 合并到 o 上，返回合并后的根。
func (o *lcNode) merge(b *lcNode) *lcNode {
	if o == emptyLcNode {
		return b
	}
	if b == emptyLcNode {
		return o
	}
	if o.l == o.r {
		if b.line.calc(o.l) < o.line.calc(o.l)-eps { // 最大值：b.line.calc(o.l) > o.line.calc(o.l)+eps
			o.line = b.line
		}
		return o
	}
	o.lo = o.lo.merge(b.lo)
	o.ro = o.ro.merge(b.ro)
	if b.line.id != 0 {
		o.update(b.line)
	}
	return o
}

func newLcRoot(l, r int) *lcNode {
	return &lcNode{lo: emptyLcNode, ro: emptyLcNode, l: l, r: r, line: newLine()}
}


// ------- 可持久化李超线段树 (浮点数版) ------- //

type pLcNode struct {
	lo, ro *pLcNode
	l, r   int
	line   line
}

func buildPlc(l, r int) *pLcNode {
	o := &pLcNode{l: l, r: r, line: newLine()}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPlc(l, m)
	o.ro = buildPlc(m+1, r)
	return o
}

func (o pLcNode) updateLine(nl line) *pLcNode {
	l, r := o.l, o.r
	m := (l + r) >> 1

	if nl.calc(m) < o.line.calc(m)-eps { // 最大值：if nl.calc(m) > o.line.calc(m)+eps
		o.line, nl = nl, o.line
	}
	if l == r {
		return &o
	}

	if nl.calc(l) < o.line.calc(l)-eps { // 最大值：if nl.calc(l) > o.line.calc(l)+eps
		o.lo = o.lo.updateLine(nl)
	} else if nl.calc(r) < o.line.calc(r)-eps { // 最大值：if nl.calc(r) > o.line.calc(r)+eps
		o.ro = o.ro.updateLine(nl)
	}
	return &o
}

// 插入线段 [ql, qr] 上的直线 nl
func (o pLcNode) insert(nl line, ql, qr int) *pLcNode {
	if ql > o.r || qr < o.l {
		return &o
	}
	if ql <= o.l && o.r <= qr {
		return o.updateLine(nl)
	}
	m := (o.l + o.r) >> 1
	if ql <= m {
		o.lo = o.lo.insert(nl, ql, qr)
	}
	if m < qr {
		o.ro = o.ro.insert(nl, ql, qr)
	}
	return &o
}

// 查询 x 处的最值
func (o *pLcNode) query(x int) (res float64) {
	if o == nil {
		return inf // 最大值：return -inf
	}
	res = inf // 最大值：res := -inf
	
	if o.line.id != 0 {
		res = o.line.calc(x)
	}
	if o.l == o.r {
		return res
	}
	m := (o.l + o.r) >> 1
	if x <= m {
		return min(res, o.lo.query(x)) // 最大值：return max(res, o.lo.query(x))
	}
	return min(res, o.ro.query(x)) // 最大值：return max(res, o.ro.query(x))
}

func newPlc(l, r int) []*pLcNode {
	return []*pLcNode{buildPlc(l, r)}
}