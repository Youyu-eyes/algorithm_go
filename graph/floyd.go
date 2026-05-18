package graph

// 返回一个二维列表，其中 (i,j) 这一项表示从 i 到 j 的最短路长度
// 如果无法从 i 到 j，则最短路长度为 math.MaxInt / 2
// 允许负数边权
// 如果计算完毕后，存在 i，使得从 i 到 i 的最短路长度小于 0，说明图中有负环
// 节点编号从 0 到 n-1
// 时间复杂度 O(n^3 + m)，其中 m 是 edges 的长度
func Floyd(n int, edges [][]int) [][]int {
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
		f[i][i] = 0
	}

	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		f[x][y] = min(f[x][y], wt) // 如果有重边，取边权最小值
		f[y][x] = min(f[y][x], wt) // 无向图
	}

	for k := range n {
		for i := range n {
			if f[i][k] == inf { // 针对稀疏图的优化
				continue
			}
			for j := range n {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j])
			}
		}
	}
	return f
}
