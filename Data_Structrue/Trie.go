package Data_Structrue

const width = 30

// ------- Trie ------- //
// 初始化：

type node struct {
	son [26]*node
	end bool
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{&node{}}
}

func (t trie) insert(word string) {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil { // 无路可走？
			cur.son[c] = &node{} // 那就造路！
		}
		cur = cur.son[c]
	}
	cur.end = true
}

func (t trie) find(word string) int {
	cur := t.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil { // 道不同，不相为谋
			return 0
		}
		cur = cur.son[c]
	}
	// 走过同样的路（2=完全匹配，1=前缀匹配）
	if cur.end {
		return 2
	}
	return 1
}

func (t trie) Search(word string) bool {
	return t.find(word) == 2
}

func (t trie) StartsWith(prefix string) bool {
	return t.find(prefix) != 0
}

// ------- Trie01 ------- //
// 初始化：t := newTrie()

type node01 struct {
	son  [2]*node01
	leaf int // 子树叶子个数
}

type trie01 struct {
	root *node01
}

func newTrie01() *trie01 {
	return &trie01{&node01{}}
}

func (t *trie01) put(val int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		bit := val >> i & 1
		if cur.son[bit] == nil {
			cur.son[bit] = &node01{}
		}
		cur = cur.son[bit]
		cur.leaf++
	}
}

func (t *trie01) del(val int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		cur = cur.son[val>>i&1]
		cur.leaf-- // 如果减成 0 了，说明子树是空的，可以理解成 cur == nil
	}
}

func (t *trie01) maxXor(val int) (ans int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		bit := val >> i & 1
		if cur.son[bit^1] != nil && cur.son[bit^1].leaf > 0 {
			ans |= 1 << i
			bit ^= 1
		}
		cur = cur.son[bit]
	}
	return
}

// ------- PersistentTrie01 ------- //
// 初始化与使用示范：
// t := newPTrie01()
//
// // 插入元素，并将返回的新根节点索引存入历史版本数组中
// // 假设我们要插入 val
// lastRoot := t.roots[len(t.roots)-1]
// newRoot := t.put(lastRoot, val)
// t.roots = append(t.roots, newRoot)
//
// // 区间查询最大异或值
// // 查询区间 [L, R] 内与 val 异或的最大值 (L, R 为基于 1 的插入次序)
// // 传入区间左端点的前一个版本 roots[L-1] 和右端点版本 roots[R]
// ans := t.maxXor(t.roots[L-1], t.roots[R], val)

type pNode01 struct {
	son  [2]int // 存储子节点在 nodes 切片中的索引，0 表示空（无子节点）
	leaf int    // 子树叶子个数 (对应 python 代码中的 cnt)
}

type pTrie01 struct {
	nodes []pNode01
	roots []int // 保存每次插入操作后生成的新根节点的索引
}

func newPTrie01() *pTrie01 {
	return &pTrie01{
		// 索引 0 作为哨兵空节点，所有未初始化的分支默认指向 0，完美避免 nil 指针检查
		nodes: make([]pNode01, 1),
		roots: []int{0},
	}
}

// 基于旧版本的根节点 prevRoot 插入新值，返回新版本根节点的索引
func (t *pTrie01) put(prevRoot int, val int) int {
	newRoot := len(t.nodes)
	// 复制旧根节点的状态，并在新根节点上增加叶子计数
	t.nodes = append(t.nodes, t.nodes[prevRoot])
	t.nodes[newRoot].leaf++

	curNew := newRoot
	curOld := prevRoot

	for i := width - 1; i >= 0; i-- {
		bit := val >> i & 1

		// 分配新的子节点索引
		newNode := len(t.nodes)
		// 复制旧字典树中对应分支节点的状态
		t.nodes = append(t.nodes, t.nodes[t.nodes[curOld].son[bit]])
		t.nodes[newNode].leaf++

		// 将新分支挂载到当前的 curNew 节点下
		t.nodes[curNew].son[bit] = newNode

		// 指针下移
		curNew = newNode
		curOld = t.nodes[curOld].son[bit]
	}

	return newRoot
}

// 查询由 root1(不包含) 到 root2(包含) 这个历史版本区间内，与 val 异或的最大值
func (t *pTrie01) maxXor(root1 int, root2 int, val int) (ans int) {
	node1 := root1
	node2 := root2

	for i := width - 1; i >= 0; i-- {
		bit := val >> i & 1

		// 我们期望走与当前 bit 相反的方向（0 走 1，1 走 0）以获取最大的异或值
		targetBit := bit ^ 1

		// 计算目标方向区间内实际新增的叶子节点数量
		cnt1 := t.nodes[t.nodes[node1].son[targetBit]].leaf
		cnt2 := t.nodes[t.nodes[node2].son[targetBit]].leaf

		// 如果 cnt2 - cnt1 > 0，说明这个区间内存在这条路径，可以走
		if cnt2-cnt1 > 0 {
			ans |= 1 << i
			node1 = t.nodes[node1].son[targetBit]
			node2 = t.nodes[node2].son[targetBit]
		} else {
			// 否则只能顺着原方向继续往下找
			node1 = t.nodes[node1].son[bit]
			node2 = t.nodes[node2].son[bit]
		}
	}
	return
}
