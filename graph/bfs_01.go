package graph

func bfs_01(n int, edges [][]int, start int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		// g[y] = append(g[y], edge{x, w}) // 无向图加上这句
	}

	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = inf
	}
	dis[start] = 0
	q := &deque{}
	q.pushBack(start)

	for !q.empty() {
		x := q.front()
		q.popFront()

		for _, e := range g[x] {
			y, w := e.to, e.wt
			if dis[x]+w < dis[y] {
				dis[y] = dis[x] + w
				if w == 0 {
					q.pushFront(y)
				} else {
					q.pushBack(y)
				}
			}
		}
	}
	return dis
}

// ------- 双端队列 ------- //
// 本题需要双端队列，具体模板见数据结构板块
type deque struct{ l, r []int }

func (q deque) empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q *deque) pushFront(v int) {
	q.l = append(q.l, v)
}

func (q *deque) pushBack(v int) {
	q.r = append(q.r, v)
}

func (q *deque) popFront() (v int) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func (q *deque) popBack() (v any) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q deque) front() int {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}
