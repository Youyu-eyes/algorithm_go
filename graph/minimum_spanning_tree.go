package graph

import (
	"slices"
	"math"
)

// 计算图的最小生成树的边权之和
// 如果图不连通，返回 math.MaxInt
// 节点编号从 0 到 n-1
// 时间复杂度 O(n + mlogm)，其中 m 是 edges 的长度
func mstKruskal(n int, edges [][]int) int {
	slices.SortFunc(edges, func(a, b []int) int { return a[2] - b[2] })

	uf := newUnionFind(n)
	sumWt := 0
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		if uf.merge(x, y) {
			sumWt += wt
		}
	}

	if uf.cc > 1 {
		return math.MaxInt
	}
	return sumWt
}

// 完整并查集模板见图论