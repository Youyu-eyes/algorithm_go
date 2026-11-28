package graph

import (
	"math/bits"
)

func lca(edges [][]int) {
    n := len(edges) + 1
    type edge struct{ to, wt int }
    g := make([][]edge, n)
    for _, e := range edges {
        // 如果题目的节点编号从 1 开始，改成 x=e[0]-1 和 y=e[1]-1
        x, y, wt := e[0], e[1], e[2]
        g[x] = append(g[x], edge{y, wt})
        g[y] = append(g[y], edge{x, wt})
    }

    const mx = 17 // bits.Len(uint(n))
    pa := make([][mx]int, n)
    dep := make([]int, n)
    dis := make([]int, n) // 如果是无权树（边权为 1），dis 可以去掉，用 dep 代替

    var dfs func(int, int)
    dfs = func(x, p int) {
        pa[x][0] = p
        for _, e := range g[x] {
            y := e.to
            if y == p {
                continue
            }
            dep[y] = dep[x] + 1
            dis[y] = dis[x] + e.wt
            dfs(y, x)
        }
    }
    dfs(0, -1)

    for i := range mx - 1 {
        for x := range pa {
            p := pa[x][i]
            if p != -1 {
                pa[x][i+1] = pa[p][i]
            } else {
                pa[x][i+1] = -1
            }
        }
    }

	// kth(x, k int) = uptoDep(x, dep[x]-k int)
    uptoDep := func(x, d int) int {
        for k := uint32(dep[x] - d); k > 0; k &= k - 1 {
            x = pa[x][bits.TrailingZeros32(k)]
        }
        return x
    }

    // 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
    getLCA := func(x, y int) int {
        if dep[x] > dep[y] {
            x, y = y, x
        }
        y = uptoDep(y, dep[x]) // 使 y 和 x 在同一深度
        if y == x {
            return x
        }
        for i := mx - 1; i >= 0; i-- {
            px, py := pa[x][i], pa[y][i]
            if px != py {
                x, y = px, py // 同时往上跳 2^i 步
            }
        }
        return pa[x][0]
    }

    // 返回 x 到 y 的距离（最短路长度）
    getDis := func(x, y int) int { return dis[x] + dis[y] - dis[getLCA(x, y)]*2 }

	_, _ = getDis, getLCA
}