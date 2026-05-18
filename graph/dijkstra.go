package graph

import (
	"container/heap"
)

// 返回从起点 start 到每个点的最短路长度 dis，如果节点 x 不可达，则 dis[x] = inf
// 要求：没有负数边权
// 时间复杂度 O(n + mlogm)，注意堆中有 O(m) 个元素
func Dijkstra(n int, edges [][]int, start int) ([]int, [][]int) {
	// 注：如果节点编号从 1 开始（而不是从 0 开始），可以把 n 加一
	type edge struct{ to, wt int }
	g := make([][]edge, n) // 邻接表
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		// g[y] = append(g[y], edge{x, wt}) // 无向图加上这行
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	h := &hp{{0, start}}

	pre := make([][]int, n)

	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		disX, x := p.dis, p.x
		if disX > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newDisY := disX + e.wt
			if newDisY < dis[y] {
				dis[y] = newDisY
				heap.Push(h, pair{newDisY, y})

				// 找到更短的路径，重置前驱节点
				pre[y] = []int{x}
			} else if newDisY == dis[y] {
				// 一样短的路径
				pre[y] = append(pre[y], x)
			}
		}
	}

	return dis, pre
}

// 需要用堆的时候，需要定义的东西
type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
