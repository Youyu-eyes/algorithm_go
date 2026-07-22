package string

func smallestRepresentation(s string) string {
    n := len(s)
    s += s
    i := 0
    j := 1
    for j < n {
        k := 0
        for k < n && s[i + k] == s[j + k] {
            k++
        }
        if k >= n {
            break
        }

        if s[i + k] < s[j + k] { // 注：如果求字典序最大，改成 >
            j += k + 1
        } else {
            i, j = j, max(j, i + k) + 1
        }
    }
    return s[i : i+n]
}
