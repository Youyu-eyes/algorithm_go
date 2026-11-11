package string

func kmp(s string, t string) (pos []int) {
    n, m := len(s), len(t)
    pi := make([]int, m)
    for i, j := 1, 0; i < m; i++ {
        for j > 0 && t[i] != t[j] {
            j = pi[j-1]
        }
        if t[i] == t[j] {
            j++
        }
        pi[i] = j
    }

    j := 0
    for i := range n {
        for j > 0 && t[j] != s[i] {
            j = pi[j-1]
        }
        if t[j] == s[i] {
            j++
        }
        if j == m {
            pos = append(pos, i - m + 1)
			j = pi[j-1]
        }
    }
	return
}