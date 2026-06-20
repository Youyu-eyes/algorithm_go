package Data_Structure

type Matrix [][]int

func Build(matrix [][]int) Matrix {
    m, n := len(matrix), len(matrix[0])
    sum := make([][]int, m+1)
    sum[0] = make([]int, n+1)
    for i, row := range matrix {
        sum[i+1] = make([]int, n+1)
        for j, x := range row {
            sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
        }
    }
    return sum
}

// 返回左上角在 (r1, c1)，右下角在 (r2, c2) 的子矩阵元素和
func (s Matrix) SumRegion(r1, c1, r2, c2 int) int {
    return s[r2+1][c2+1] - s[r2+1][c1] - s[r1][c2+1] + s[r1][c1]
}