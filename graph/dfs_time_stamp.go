package graph

func dfnOrder(root int, g [][]int, a []int) {
	n := len(a)
	dfnToVal := make([]int, n)
	// 闭区间 [l,r]   0 <= l <= r <= n-1
	nodes := make([]struct{ l, r int }, n)
	dfn := 0 // Depth-First Numbering
	var buildDFN func(int, int)
	buildDFN = func(v, fa int) {
		nodes[v].l = dfn
		dfnToVal[dfn] = a[v] // 写这里是先序遍历，如果题目要求后序遍历，把这一行和下一行的 dfn++ 移到 for 循环后面
		dfn++
		for _, w := range g[v] {
			if w != fa {
				buildDFN(w, v)
			}
		}
		nodes[v].r = dfn - 1
	}
	buildDFN(root, -1)

	// 判断 x 是否为 y 的祖先
	isAncestor := func(f, v int) bool {
		return nodes[f].l < nodes[v].l && nodes[v].l <= nodes[f].r
	}

	// dfnToNodeID := make([]int, n)
	// for v, o := range nodes {
	// 	dfnToNodeID[o.l] = v
	// }

	_ = isAncestor
}

func timeStamp(n int, g [][]int) {
	in := make([]int, n)
	out := make([]int, n)
	clock := 0
	var sign func(int, int)
	sign = func(x, fa int) {
		clock++
		in[x] = clock
		for _, y := range g[x] {
			if y != fa {
				sign(y, x)
			}
		}
		out[x] = clock
	}
	sign(0, -1)

	// 判断 x 是否为 y 的祖先
	isAncestor := func(x, y int) bool {
		return in[x] < in[y] && in[y] <= out[x]
	}

	_ = isAncestor
}