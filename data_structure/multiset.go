package dataStructure

import (
	"cmp"
	"time"
)

type multiNode[K comparable] struct {
	son      [2]*multiNode[K]
	priority uint
	key      K
	keyCnt   int
	subSize  int
}

func (o *multiNode[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *multiNode[K]) maintain() {
	o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size()
}

func (o *multiNode[K]) rotate(d int) *multiNode[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type multiset[K comparable] struct {
	rd         uint
	root       *multiNode[K]
	comparator func(a, b K) int
}

func (t *multiset[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *multiset[K]) size() int   { return t.root.size() }
func (t *multiset[K]) empty() bool { return t.size() == 0 }

func (t *multiset[K]) _put(o *multiNode[K], key K) *multiNode[K] {
	if o == nil {
		o = &multiNode[K]{priority: t.fastRand(), key: key, keyCnt: 1}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
			o.keyCnt++
		} else {
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *multiset[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *multiset[K]) _delete(o *multiNode[K], key K) *multiNode[K] {
	if o == nil {
		return nil
	}
	if c := t.comparator(key, o.key); c != 0 {
		d := 0
		if c > 0 {
			d = 1
		}
		o.son[d] = t._delete(o.son[d], key)
	} else {
		if o.keyCnt > 1 {
			o.keyCnt--
		} else {
			if o.son[1] == nil {
				return o.son[0]
			}
			if o.son[0] == nil {
				return o.son[1]
			}
			d := 0
			if o.son[0].priority > o.son[1].priority {
				d = 1
			}
			o = o.rotate(d)
			o.son[d] = t._delete(o.son[d], key)
		}
	}
	o.maintain()
	return o
}

func (t *multiset[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *multiset[K]) min() *multiNode[K] { return t.kth(0) }
func (t *multiset[K]) max() *multiNode[K] { return t.kth(t.size() - 1) }

func (t *multiset[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *multiset[K]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size() + o.keyCnt
			break
		}
	}
	return
}

func (t *multiset[K]) kth(k int) (o *multiNode[K]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + o.keyCnt
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func (t *multiset[K]) prev(key K) *multiNode[K] { return t.kth(t.lowerBoundIndex(key) - 1) }

func (t *multiset[K]) next(key K) *multiNode[K] { return t.kth(t.upperBoundIndex(key)) }

func (t *multiset[K]) find(key K) *multiNode[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newMultiset[K cmp.Ordered]() *multiset[K] {
	return &multiset[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[K],
	}
}

func neMultisetWith[K comparable](comp func(a, b K) int) *multiset[K] {
	return &multiset[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: comp,
	}
}
