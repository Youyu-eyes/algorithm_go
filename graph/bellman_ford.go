package graph

import "math"

// BellmanFord 计算单源最值路径与环路检测
// 参数:
// n     int     : 点的总数 (默认节点编号从 0 到 n-1，如果是 1 到 n，请传入 n+1)
// edges [][]int : 边集，edges[i] = [u, v, w]
// start int     : 起点
// short bool    : true (求最短路并判断负权环)，false (求最长路并判断正权环)
// 返回值:
// []int : 距离数组 dist
// bool  : 是否存在对应的环 (short=true 时返回是否有负环，short=false 时返回是否有正环)

func BellmanFord(n int, edges [][]int, start int, short bool) ([]int, bool) {
	const INF = math.MaxInt / 2

	// 初始化距离数组
	dist := make([]int, n)
	if short {
		for i := range dist {
			dist[i] = INF
		}
	} else {
		for i := range dist {
			dist[i] = -INF
		}
	}
	dist[start] = 0

	for i := 0; i < n-1; i++ {
		updated := false
		for _, edge := range edges {
			u, v, w := edge[0], edge[1], edge[2]
			if short {
				if dist[u] != INF && dist[u]+w < dist[v] {
					dist[v] = dist[u] + w
					updated = true
				}
			} else {
				if dist[u] != -INF && dist[u]+w > dist[v] {
					dist[v] = dist[u] + w
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}

	hasCycle := false
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if short {
			if dist[u] != INF && dist[u]+w < dist[v] {
				hasCycle = true
				break
			}
		} else {
			if dist[u] != -INF && dist[u]+w > dist[v] {
				hasCycle = true
				break
			}
		}
	}

	return dist, hasCycle
}