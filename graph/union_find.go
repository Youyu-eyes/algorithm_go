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
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己，大小为 1
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return unionFind{fa, sz, n}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	// 如果 fa[x] == x，则表示 x 是代表元
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x]) // fa 改成代表元
	}
	return u.fa[x]
}

// 判断 x 和 y 是否在同一个集合
func (u unionFind) same(x, y int) bool {
	// 如果 x 的代表元和 y 的代表元相同，那么 x 和 y 就在同一个集合
	// 这就是代表元的作用：用来快速判断两个元素是否在同一个集合
	return u.find(x) == u.find(y)
}

// 把 from 所在集合合并到 to 所在集合中
// 返回是否合并成功
func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		return false
	}

	// 按秩合并
	if u.sz[x] > u.sz[y] {
		x, y = y, x
	}

	u.fa[x] = y        // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
	u.sz[y] += u.sz[x] // 更新集合大小（注意集合大小保存在代表元上）
	// 无需更新 sz[x]，因为我们不用 sz[x] 而是用 sz[find(x)] 获取集合大小，但 find(x) == y，我们不会再访问 sz[x]
	u.cc-- // 成功合并，连通块个数减一
	return true
}

// 返回 x 所在集合的大小
func (u unionFind) size(x int) int {
	return u.sz[u.find(x)] // 集合大小保存在代表元上
}

// ------- 带权并查集 ------- //
// 初始化：uf := newUnionFindWithIndex(n)

type unionFindWithIndex struct {
	fa  []int // 代表元
	dis []int // dis[x] 表示 x 到（x 所在集合的）代表元的距离
}

func newUnionFindWithIndex(n int) unionFindWithIndex {
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己，自己到自己的距离是 0
	fa := make([]int, n)
	dis := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFindWithIndex{fa, dis}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩
func (u unionFindWithIndex) find(x int) int {
	if u.fa[x] != x {
		root := u.find(u.fa[x])
		u.dis[x] += u.dis[u.fa[x]] // 递归更新 x 到其代表元的距离
		u.fa[x] = root
	}
	return u.fa[x]
}

// 判断 x 和 y 是否在同一个集合（同普通并查集）
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
	u.dis[x] = value + u.dis[to] - u.dis[from] // 计算 x 到其代表元 y 的距离
	u.fa[x] = y
	return true
}
