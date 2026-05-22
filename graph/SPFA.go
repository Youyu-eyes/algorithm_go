package graph

// SPFA 计算单源最值路径与环路检测（Bellman-Ford 的队列优化版）
// 参数:
//   n     int     : 点的总数 (默认节点编号从 0 到 n-1，如果是 1 到 n，请传入 n+1)
//   edges [][]int : 边集，edges[i] = [u, v, w]
//   start int     : 起点
//   short bool    : true (求最短路并判断负权环)，false (求最长路并判断正权环)
// 返回值:
//   []int : 距离数组 dist
//   bool  : 是否存在对应的环 (short=true 时返回是否有负环，short=false 时返回是否有正环)
func SPFA(n int, edges [][]int, start int, short bool) ([]int, bool) {
	dist := make([]int, n)
	if short {
		for i := range dist {
			dist[i] = inf
		}
	} else {
		for i := range dist {
			dist[i] = -inf
		}
	}
	dist[start] = 0

	g := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], [2]int{v, w})
	}

	queue := make([]int, 0, n)
	inQueue := make([]bool, n)
	cnt := make([]int, n)

	queue = append(queue, start)
	inQueue[start] = true
	cnt[start] = 1

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		inQueue[u] = false

		for _, edge := range g[u] {
			v, w := edge[0], edge[1]
			relaxed := false

			if short {
				if dist[u] != inf && dist[u]+w < dist[v] {
					dist[v] = dist[u] + w
					relaxed = true
				}
			} else {
				if dist[u] != -inf && dist[u]+w > dist[v] {
					dist[v] = dist[u] + w
					relaxed = true
				}
			}

			if relaxed {
				if !inQueue[v] {
					queue = append(queue, v)
					inQueue[v] = true
					cnt[v]++
					if cnt[v] >= n {
						return dist, true
					}
				}
			}
		}
	}

	return dist, false
}