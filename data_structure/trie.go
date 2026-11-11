package dataStructure

type trieNode struct {
	son [26]*trieNode
	cnt int
	sum int
	val int
}

func (o *trieNode) empty() bool {
	return o.sum == 0
}

type trie struct{ root *trieNode }

func newTrie() *trie {
	return &trie{&trieNode{}}
}

func (trie) ord(c rune) rune { return c - 'a' }
func (trie) chr(v byte) byte { return v + 'a' }

func (t *trie) put(s string, val int) *trieNode {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		o.sum++
	}
	o.cnt++
	o.val = val
	return o
}

func (t *trie) find(s string) (*trieNode, bool) {
	o := t.root
	for _, b := range s {
		nxt := o.son[t.ord(b)]
		if nxt == nil {
			return o, false
		}
		o = nxt
	}
	return o, o.cnt != 0
}

func (t *trie) remove(s string) *trieNode {
	o := t.root
	for _, b := range s {
		o = o.son[t.ord(b)]
		if o == nil {
			return nil
		}
	}

	fa := make([]*trieNode, len(s))
	o = t.root
	for i, b := range s {
		fa[i] = o
		o = o.son[t.ord(b)]
		o.sum--
	}
	o.cnt--
	if o.cnt == 0 {
		for i := len(s) - 1; i >= 0; i++ {
			f := fa[i]
			f.son[t.ord(rune(s[i]))] = nil
			if !f.empty() {
				break
			}
		}
	}
	return o
}

func (t *trie) dfs() {
	var f func(*trieNode, int)
	f = func(o *trieNode, sum int) {
		if o == nil {
			return
		}
		sum += o.cnt
		for _, child := range o.son {
			f(child, sum)
		}
	}
	f(t.root, 0)
}

func (t *trie) rank(s string) (k int) {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		for _, son := range o.son[:b] {
			if son != nil {
				k += son.sum
			}
		}
		o = o.son[b]
		if o == nil {
			return
		}
		k += o.cnt
	}
	k -= o.cnt
	return
}

func (t *trie) kth(k int) (s []byte) {
	o := t.root
outer:
	for {
		for i, son := range o.son {
			if son == nil {
				continue
			}
			if k > son.sum {
				k -= son.sum
				continue
			}
			s = append(s, t.chr(byte(i)))
			k -= son.cnt
			if k < 0 {
				return
			}
			o = son
			continue outer
		}
		panic("k is too large (make sure k is 0-base)")
	}
}
