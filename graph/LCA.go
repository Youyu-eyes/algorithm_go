package graph

import "math/bits"

// ------- 带权 LCA 模板 ------- //
type LcaBinaryLifting struct {
	n     int     // 节点个数
	m     int     // 二进制提升层数
	pa    [][]int // pa[k][v] 表示 v 的第 2^k 级祖先，-1 表示不存在
	depth []int   // 节点深度（根深度为 0）
	dis   []int   // 从根到节点的距离（路径权值和）
}

func newLcaBinaryLifting(edges [][]int) *LcaBinaryLifting {
	n := len(edges) + 1
	m := bits.Len(uint(n)) // 最多需要的二进制位数
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}

	pa := make([][]int, m)
	for i := range pa {
		pa[i] = make([]int, n)
		for j := range pa[i] {
			pa[i][j] = -1
		}
	}
	depth := make([]int, n)
	dis := make([]int, n)

	var dfs func(x, p int)
	dfs = func(x, p int) {
		pa[0][x] = p
		for _, e := range g[x] {
			y, w := e.to, e.wt
			if y == p {
				continue
			}
			depth[y] = depth[x] + 1
			dis[y] = dis[x] + w
			dfs(y, x)
		}
	}
	dfs(0, -1)

	for i := 0; i < m-1; i++ {
		for v := 0; v < n; v++ {
			if p := pa[i][v]; p != -1 {
				pa[i+1][v] = pa[i][p]
			}
		}
	}

	return &LcaBinaryLifting{
		n:     n,
		m:     m,
		pa:    pa,
		depth: depth,
		dis:   dis,
	}
}

// 返回 node 的第 k 个祖先节点，若不存在返回 -1
func (l *LcaBinaryLifting) getKthAncestor(node int, k int) int {
	for i := 0; i < l.m; i++ {
		if k>>i&1 == 1 {
			node = l.pa[i][node]
			if node == -1 {
				return -1
			}
		}
	}
	return node
}

// 返回 x 和 y 的最近公共祖先
func (l *LcaBinaryLifting) getLca(x, y int) int {
	if l.depth[x] > l.depth[y] {
		x, y = y, x
	}
	// 将 y 提升到与 x 同深度
	y = l.getKthAncestor(y, l.depth[y]-l.depth[x])
	if y == x {
		return x
	}
	for i := l.m - 1; i >= 0; i-- {
		if l.pa[i][x] != l.pa[i][y] {
			x = l.pa[i][x]
			y = l.pa[i][y]
		}
	}
	return l.pa[0][x]
}

// 返回 x 到 y 的距离（最短路径长度，仅适用于带权树）
func (l *LcaBinaryLifting) getDis(x, y int) int {
	lca := l.getLca(x, y)
	return l.dis[x] + l.dis[y] - 2*l.dis[lca]
}

// ------- 不带权 LCA 模板 ------- //
type LcaBinaryLiftingUnweighted struct {
	n     int
	m     int
	pa    [][]int
	depth []int
}

func newLcaBinaryLiftingUnweighted(edges [][]int) *LcaBinaryLiftingUnweighted {
	n := len(edges) + 1
	m := bits.Len(uint(n))
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	pa := make([][]int, m)
	for i := range pa {
		pa[i] = make([]int, n)
		for j := range pa[i] {
			pa[i][j] = -1
		}
	}
	depth := make([]int, n)

	var dfs func(x, p int)
	dfs = func(x, p int) {
		pa[0][x] = p
		for _, y := range g[x] {
			if y == p {
				continue
			}
			depth[y] = depth[x] + 1
			dfs(y, x)
		}
	}
	dfs(0, -1)

	for i := 0; i < m-1; i++ {
		for v := 0; v < n; v++ {
			if p := pa[i][v]; p != -1 {
				pa[i+1][v] = pa[i][p]
			}
		}
	}

	return &LcaBinaryLiftingUnweighted{
		n:     n,
		m:     m,
		pa:    pa,
		depth: depth,
	}
}

// 返回 node 的第 k 个祖先节点，若不存在返回 -1
func (l *LcaBinaryLiftingUnweighted) getKthAncestor(node int, k int) int {
	for i := 0; i < l.m; i++ {
		if k>>i&1 == 1 {
			node = l.pa[i][node]
			if node == -1 {
				return -1
			}
		}
	}
	return node
}

// 返回 x 和 y 的最近公共祖先
func (l *LcaBinaryLiftingUnweighted) getLca(x, y int) int {
	if l.depth[x] > l.depth[y] {
		x, y = y, x
	}
	y = l.getKthAncestor(y, l.depth[y]-l.depth[x])
	if y == x {
		return x
	}
	for i := l.m - 1; i >= 0; i-- {
		if l.pa[i][x] != l.pa[i][y] {
			x = l.pa[i][x]
			y = l.pa[i][y]
		}
	}
	return l.pa[0][x]
}

// 返回 x 到 y 的距离（无权重时即为树上边数，用深度差计算）
func (l *LcaBinaryLiftingUnweighted) getDis(x, y int) int {
	lca := l.getLca(x, y)
	return l.depth[x] + l.depth[y] - 2*l.depth[lca]
}
