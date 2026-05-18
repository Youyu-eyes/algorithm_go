package graph

// 返回图的二染色
// 如果是二分图，返回每个节点的颜色，用 1 和 2 表示两种颜色
// 如果不是二分图，返回空列表
// 时间复杂度 O(n+m)，n 是点数，m 是边数
func colorBipartite(n int, edges [][]int) []int8 {
	// 建图（节点编号从 0 到 n-1）
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// colors[i] = 0 表示未访问节点 i
	// colors[i] = 1 表示节点 i 为红色
	// colors[i] = 2 表示节点 i 为蓝色
	colors := make([]int8, n)

	var dfs func(int, int8) bool
	dfs = func(x int, c int8) bool {
		colors[x] = c // 节点 x 染成颜色 c
		for _, y := range g[x] {
			// 邻居 y 的颜色与 x 的相同，说明不是二分图，返回 false
			// 或者继续递归，发现不是二分图，返回 false
			if colors[y] == c ||
				colors[y] == 0 && !dfs(y, 3-c) { // 1 和 2 交替染色
				return false
			}
		}
		return true
	}

	// 可能有多个连通块
	for i, c := range colors {
		if c == 0 && !dfs(i, 1) {
			// 从节点 i 开始递归，发现 i 所在连通块不是二分图
			return nil
		}
	}
	return colors
}
