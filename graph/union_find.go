package graph

// ------- 并查集 ------- //
// 初始化：uf := newUnionFind(n)

type unionFind struct {
	fa []int // 代表元
	sz []int // 集合大小
	cc int   // 连通块个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return unionFind{fa, sz, n}
}

func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u unionFind) same(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}

	// 按秩合并
	if u.sz[x] > u.sz[y] {
		x, y = y, x
	}

	u.fa[x] = y
	u.sz[y] += u.sz[x]
	u.cc--
	return true
}

func (u unionFind) size(x int) int {
	return u.sz[u.find(x)]
}

// ------- 带权并查集 ------- //
// 初始化：uf := newUnionFindWithIndex(n)

type unionFindWithIndex struct {
	fa  []int
	dis []int
}

func newUnionFindWithIndex(n int) unionFindWithIndex {
	fa := make([]int, n)
	dis := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFindWithIndex{fa, dis}
}

func (u unionFindWithIndex) find(x int) int {
	if u.fa[x] != x {
		root := u.find(u.fa[x])
		u.dis[x] += u.dis[u.fa[x]]
		u.fa[x] = root
	}
	return u.fa[x]
}

func (u unionFindWithIndex) same(x, y int) bool {
	return u.find(x) == u.find(y)
}

// 计算从 from 到 to 的相对距离
// 调用时需保证 from 和 to 在同一个集合中，否则返回值无意义
func (u unionFindWithIndex) getRelativeDistance(from, to int) int {
	u.find(from)
	u.find(to)
	// to-from = (x-from) - (x-to) = dis[from] - dis[to]
	return u.dis[from] - u.dis[to]
}

// 合并 from 和 to，新增信息 to - from = value
// 其中 to 和 from 表示未知量，下文的 x 和 y 也表示未知量
// 如果 from 和 to 不在同一个集合，返回 true，否则返回是否与已知信息矛盾
func (u unionFindWithIndex) merge(from, to int, value int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		// to-from = (x-from) - (x-to) = dis[from] - dis[to] = value
		return u.dis[from]-u.dis[to] == value
	}
	//    x --------- y
	//   /           /
	// from ------- to
	// 已知 x-from = dis[from] 和 y-to = dis[to]，现在合并 from 和 to，新增信息 to-from = value
	// 由于 y-from = (y-x) + (x-from) = (y-to) + (to-from)
	// 所以 y-x = (to-from) + (y-to) - (x-from) = value + dis[to] - dis[from]
	u.dis[x] = value + u.dis[to] - u.dis[from]
	u.fa[x] = y
	return true
}
