package graph

// 返回有向无环图（DAG）的其中一个拓扑序
// 如果图中有环，返回 nil
// 节点编号从 0 到 n-1
func topoSort(n int, edges [][]int) []int {
	g := make([][]int, n)
	inDeg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		inDeg[y]++ // 统计 y 的先修课数量
	}

	q := make([]int, 0, n)
	topoOrder := q
	for i, d := range inDeg {
		if d == 0 { // 没有先修课，可以直接上
			q = append(q, i) // 加入学习队列
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			inDeg[y]--         // 修完 x 后， y 的先修课数量减一
			if inDeg[y] == 0 { // y 的先修课全部上完
				q = append(q, y) // 加入学习队列
			}
		}
	}

	if cap(q) > 0 { // 图中有环
		return nil
	}
	return topoOrder[:n]
}
