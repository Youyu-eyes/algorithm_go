package string

func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR - i + 1)
		}
		for i + z[i] < n && s[z[i]] == s[i + z[i]] {
			boxL, boxR = i, i + z[i]
			z[i]++
		}
	}
	z[0] = n
	return z
}