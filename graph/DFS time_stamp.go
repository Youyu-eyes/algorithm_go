package graph

func timeStamp(n int, g [][]int) ([]int, []int) {
	in := make([]int, n)
	out := make([]int, n)
	clock := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		clock++
		in[x] = clock
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
			}
		}
		out[x] = clock
	}
	dfs(0, -1)

	return in, out
}

// 判断 x 是否为 y 的祖先
func isAncestor (in, out []int, x, y int) bool {
	return in[x] < in[y] && in[y] <= out[x]
}