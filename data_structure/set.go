package dataStructure

import (
	"cmp"
	"time"
)

type setNode[K comparable] struct {
	son      [2]*setNode[K]
	priority uint
	key      K
	subSize  int
}

func (o *setNode[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *setNode[K]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *setNode[K]) rotate(d int) *setNode[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type set[K comparable] struct {
	rd         uint
	root       *setNode[K]
	comparator func(a, b K) int
}

func (t *set[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *set[K]) size() int   { return t.root.size() }
func (t *set[K]) empty() bool { return t.size() == 0 }

func (t *set[K]) _put(o *setNode[K], key K) *setNode[K] {
	if o == nil {
		o = &setNode[K]{priority: t.fastRand(), key: key}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
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

func (t *set[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *set[K]) _delete(o *setNode[K], key K) *setNode[K] {
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
	o.maintain()
	return o
}

func (t *set[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *set[K]) min() *setNode[K] { return t.kth(0) }
func (t *set[K]) max() *setNode[K] { return t.kth(t.size() - 1) }

func (t *set[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *set[K]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size() + 1
			break
		}
	}
	return
}

func (t *set[K]) kth(k int) (o *setNode[K]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + 1
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func (t *set[K]) prev(key K) *setNode[K] { return t.kth(t.lowerBoundIndex(key) - 1) }

func (t *set[K]) next(key K) *setNode[K] { return t.kth(t.upperBoundIndex(key)) }

func (t *set[K]) find(key K) *setNode[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newTreap[K cmp.Ordered]() *set[K] {
	return &set[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[K],
	}
}

func newTreapWith[K comparable](comp func(a, b K) int) *set[K] {
	return &set[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: comp,
	}
}