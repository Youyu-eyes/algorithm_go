package graph

import (
	"container/heap"
)

// zkw = dinic + dijkstra + 势能函数
// 贪心地想，每次在最短路上增广

func zkw(n int, edges [][]int, s, t int) (maxFlow int, minCost int) {
	type edge struct{ to, wt, ct, rev int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt, ct := e[0], e[1], e[2], e[3]
		g[x] = append(g[x], edge{y, wt, ct, len(g[y])})
		g[y] = append(g[y], edge{x, 0, -ct, len(g[x]) - 1})
	}

	dis := make([]int, n)
	hi := make([]int, n)

	dijkstra := func() bool {
		for i := range dis {
			dis[i] = inf
		}
		dis[s] = 0
		h := &hp{{0, s}}

		for h.Len() > 0 {
			p := heap.Pop(h).(pair)
			disX, x := p.dis, p.x
			if disX > dis[x] {
				continue
			}
			for _, e := range g[x] {
				y := e.to
				if e.wt > 0 {
					newDisY := disX + e.ct + hi[x] - hi[y]
					if newDisY < dis[y] {
						dis[y] = newDisY
						heap.Push(h, pair{newDisY, y})
					}
				}
			}
		}
		return dis[t] < inf
	}

	var cur = make([]int, n)
	vis := make([]bool, n)

    var dfs func(int, int) int
	dfs = func(u, flow int) (pushed int) {
		if u == t || flow == 0 {
			return flow
		}
		vis[u] = true
		
		for ; cur[u] < len(g[u]); cur[u]++ {
			e := &g[u][cur[u]]
			v, wt, ct, rev := e.to, e.wt, e.ct, e.rev
			
			if !vis[v] && wt > 0 && ct + hi[u] - hi[v] == 0 {
				push := dfs(v, min(flow - pushed, wt))
				if push > 0 {
					e.wt -= push
					g[v][rev].wt += push
					pushed += push
					if pushed == flow {
						break
					}
				}
			}
		}
		vis[u] = false
		return
	}

	for dijkstra() {
		for i := range n {
			if dis[i] < inf {
				hi[i] += dis[i]
			}
			cur[i] = 0
		}

		for {
			push := dfs(s, inf)
			if push == 0 {
				break
			}
			maxFlow += push
			minCost += push * hi[t]
		}
	}
	return
}

// 堆用到的函数在 Dijkstra.go 中