package Data_Structure

import (
	"math/bits"
)

const maxN = 150005
const maxLog = 18

// func init() {
//     fmt.Println(bits.Len(uint(maxN)))
// }

// ------- 最大值 ST 表 ------- //
// 初始化：buildSTMax(nums)
// 查询：  queryMax(l, r)

var stMax = make([]int, maxLog * maxN) 

func buildSTMax(nums []int) {
	n := len(nums)
	w := bits.Len(uint(n))

	for i := range n {
		stMax[i] = nums[i]
	}

	for k := 1; k < w; k++ {
		step := 1 << (k - 1)
		currOffset := k * maxN
		prevOffset := (k - 1) * maxN
		
		for i := 0; i + step + (1 << (k - 1)) <= n; i++ {
			v1 := stMax[prevOffset + i]
			v2 := stMax[prevOffset + i + step]
			stMax[currOffset + i] = max(v1, v2)
		}
	}
}

func queryMax(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	offset := k * maxN
	
	v1 := stMax[offset + l]
	v2 := stMax[offset + r - (1 << k)]
	return max(v1, v2)
}


// ------- 最小值 ST 表 ------- //
// 初始化：buildSTMin(nums)
// 查询：  queryMin(l, r)

var stMin = make([]int, maxLog * maxN)

func buildSTMin(nums []int) {
	n := len(nums)
	w := bits.Len(uint(n))

	for i := range n {
		stMin[i] = nums[i] 
	}

	for k := 1; k < w; k++ {
		step := 1 << (k - 1)
		currOffset := k * maxN
		prevOffset := (k - 1) * maxN
		
		for i := 0; i + step + (1 << (k - 1)) <= n; i++ {
			v1 := stMin[prevOffset + i]
			v2 := stMin[prevOffset + i + step]
			stMin[currOffset + i] = min(v1, v2)
		}
	}
}

func queryMin(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	offset := k * maxN
	
	v1 := stMin[offset + l]
	v2 := stMin[offset + r - (1 << k)]

	return min(v1, v2)
}


// ------- GCD ST 表 ------- //
// 初始化：buildSTGcd(nums)
// 查询：  queryGcd(l, r)

var stGcd = make([]int, maxLog*maxN)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func buildSTGcd(nums []int) {
	n := len(nums)
	w := bits.Len(uint(n))

	for i := range n {
		stGcd[i] = nums[i] 
	}

	for k := 1; k < w; k++ {
		step := 1 << (k - 1)
		currOffset := k * maxN
		prevOffset := (k - 1) * maxN
		
		for i := 0; i + step + (1 << (k - 1)) <= n; i++ {
			v1 := stGcd[prevOffset + i]
			v2 := stGcd[prevOffset + i + step]
			stGcd[currOffset + i] = gcd(v1, v2)
		}
	}
}

func queryGcd(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	offset := k * maxN
	
	v1 := stGcd[offset + l]
	v2 := stGcd[offset + r - (1 << k)]
	
	return gcd(v1, v2)
}


// ------- 最大值下标 ST 表 ------- //
// 初始化：buildSTMaxIdx(nums)
// 查询：  queryMaxIdx(l, r) 返回区间 [l, r) 最大值的下标（相等时取最左）

var stMaxIdx = make([]int, maxLog*maxN)
var arrMaxIdx = make([]int, maxN)

func buildSTMaxIdx(nums []int) {
	n := len(nums)
	w := bits.Len(uint(n))

	for i := range n {
		arrMaxIdx[i] = nums[i]
		stMaxIdx[i] = i
	}

	for k := 1; k < w; k++ {
		step := 1 << (k - 1)
		currOffset := k * maxN
		prevOffset := (k - 1) * maxN
		
		for i := 0; i+step+(1<<(k-1)) <= n; i++ {
			idx1 := stMaxIdx[prevOffset+i]
			idx2 := stMaxIdx[prevOffset+i+step]

			if arrMaxIdx[idx1] >= arrMaxIdx[idx2] {
				stMaxIdx[currOffset+i] = idx1
			} else {
				stMaxIdx[currOffset+i] = idx2
			}
		}
	}
}

func queryMaxIdx(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	offset := k * maxN
	
	idx1 := stMaxIdx[offset+l]
	idx2 := stMaxIdx[offset+r-(1<<k)]

	if arrMaxIdx[idx1] >= arrMaxIdx[idx2] {
		return idx1
	}
	return idx2
}

// ------- 最小值下标 ST 表 ------- //
// 初始化：buildSTMinIdx(nums)
// 查询：  queryMinIdx(l, r) 返回区间 [l, r) 最小值的下标（相等时取最左）

var stMinIdx = make([]int, maxLog*maxN)
var arrMinIdx = make([]int, maxN)

func buildSTMinIdx(nums []int) {
	n := len(nums)
	w := bits.Len(uint(n))

	for i := range n {
		arrMinIdx[i] = nums[i]
		stMinIdx[i] = i
	}

	for k := 1; k < w; k++ {
		step := 1 << (k - 1)
		currOffset := k * maxN
		prevOffset := (k - 1) * maxN
		
		for i := 0; i+step+(1<<(k-1)) <= n; i++ {
			idx1 := stMinIdx[prevOffset+i]
			idx2 := stMinIdx[prevOffset+i+step]

			if arrMinIdx[idx1] <= arrMinIdx[idx2] {
				stMinIdx[currOffset+i] = idx1
			} else {
				stMinIdx[currOffset+i] = idx2
			}
		}
	}
}

func queryMinIdx(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	offset := k * maxN
	
	idx1 := stMinIdx[offset+l]
	idx2 := stMinIdx[offset+r-(1<<k)]
	
	if arrMinIdx[idx1] <= arrMinIdx[idx2] {
		return idx1
	}
	return idx2
}


const maxR = 505
const maxC = 505
const maxLogR = 10
const maxLogC = 10

const pageSize = maxR * maxC
const kStride = maxLogC * pageSize

// ======= 最大值 2D ST 表 ======= //
// 初始化：buildSTMax2D(matrix)
// 查询：  queryMax2D(r1, r2, c1, c2) 范围 [r1, r2), [c1, c2)

var stMax2D = make([]int, maxLogR*maxLogC*pageSize)

func buildSTMax2D(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m := len(matrix)
	n := len(matrix[0])
	wR := bits.Len(uint(m))
	wC := bits.Len(uint(n))

	for i := range m {
		for j := range n {
			stMax2D[i*maxC+j] = matrix[i][j]
		}
	}

	for l := 1; l < wC; l++ {
		stepC := 1 << (l - 1)
		currLOffset := l * pageSize
		prevLOffset := (l - 1) * pageSize

		for i := range m {
			for j := 0; j+stepC+(1<<(l-1)) <= n; j++ {
				v1 := stMax2D[prevLOffset+i*maxC+j]
				v2 := stMax2D[prevLOffset+i*maxC+j+stepC]
				stMax2D[currLOffset+i*maxC+j] = max(v1, v2)
			}
		}
	}

	for k := 1; k < wR; k++ {
		stepR := 1 << (k - 1)
		currKOffset := k * kStride
		prevKOffset := (k - 1) * kStride

		for l := range wC {
			lOffset := l * pageSize
			currBase := currKOffset + lOffset
			prevBase := prevKOffset + lOffset

			for i := 0; i+stepR+(1<<(k-1)) <= m; i++ {
				for j := 0; j+(1<<l) <= n; j++ {
					v1 := stMax2D[prevBase+i*maxC+j]
					v2 := stMax2D[prevBase+(i+stepR)*maxC+j]
					stMax2D[currBase+i*maxC+j] = max(v1, v2)
				}
			}
		}
	}
}

func queryMax2D(r1, r2, c1, c2 int) int {
	k := bits.Len(uint(r2-r1)) - 1
	l := bits.Len(uint(c2-c1)) - 1

	offset := k*kStride + l*pageSize
	lenR := 1 << k
	lenC := 1 << l

	v1 := stMax2D[offset+r1*maxC+c1]
	v2 := stMax2D[offset+(r2-lenR)*maxC+c1]
	v3 := stMax2D[offset+r1*maxC+(c2-lenC)]
	v4 := stMax2D[offset+(r2-lenR)*maxC+(c2-lenC)]

	max1 := max(v1, v2)
	max2 := max(v3, v4)

	return max(max1, max2)
}


// ======= 最小值 2D ST 表 ======= //
// 初始化：buildSTMin2D(matrix)
// 查询：  queryMin2D(r1, r2, c1, c2)

var stMin2D = make([]int, maxLogR*maxLogC*pageSize)

func buildSTMin2D(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m := len(matrix)
	n := len(matrix[0])
	wR := bits.Len(uint(m))
	wC := bits.Len(uint(n))

	for i := range m {
		for j := range n {
			stMin2D[i*maxC+j] = matrix[i][j]
		}
	}

	for l := 1; l < wC; l++ {
		stepC := 1 << (l - 1)
		currLOffset := l * pageSize
		prevLOffset := (l - 1) * pageSize

		for i := range m {
			for j := 0; j+stepC+(1<<(l-1)) <= n; j++ {
				v1 := stMin2D[prevLOffset+i*maxC+j]
				v2 := stMin2D[prevLOffset+i*maxC+j+stepC]
				stMin2D[currLOffset+i*maxC+j] = min(v1, v2)
			}
		}
	}

	for k := 1; k < wR; k++ {
		stepR := 1 << (k - 1)
		currKOffset := k * kStride
		prevKOffset := (k - 1) * kStride

		for l := range wC {
			lOffset := l * pageSize
			currBase := currKOffset + lOffset
			prevBase := prevKOffset + lOffset

			for i := 0; i+stepR+(1<<(k-1)) <= m; i++ {
				for j := 0; j+(1<<l) <= n; j++ {
					v1 := stMin2D[prevBase+i*maxC+j]
					v2 := stMin2D[prevBase+(i+stepR)*maxC+j]
					stMin2D[currBase+i*maxC+j] = min(v1, v2)
				}
			}
		}
	}
}

func queryMin2D(r1, r2, c1, c2 int) int {
	k := bits.Len(uint(r2-r1)) - 1
	l := bits.Len(uint(c2-c1)) - 1

	offset := k*kStride + l*pageSize
	lenR := 1 << k
	lenC := 1 << l

	v1 := stMin2D[offset+r1*maxC+c1]
	v2 := stMin2D[offset+(r2-lenR)*maxC+c1]
	v3 := stMin2D[offset+r1*maxC+(c2-lenC)]
	v4 := stMin2D[offset+(r2-lenR)*maxC+(c2-lenC)]

	min1 := min(v1, v2)
	min2 := min(v3, v4)

	return min(min1, min2)
}


// ======= GCD 2D ST 表 ======= //
// 初始化：buildSTGcd2D(matrix)
// 查询：  queryGcd2D(r1, r2, c1, c2)

var stGcd2D = make([]int, maxLogR*maxLogC*pageSize)

// func gcd(a, b int) int { for b != 0 { a, b = b, a%b }; return a }

func buildSTGcd2D(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m := len(matrix)
	n := len(matrix[0])
	wR := bits.Len(uint(m))
	wC := bits.Len(uint(n))

	for i := range m {
		for j := range n {
			stGcd2D[i*maxC+j] = matrix[i][j]
		}
	}

	for l := 1; l < wC; l++ {
		stepC := 1 << (l - 1)
		currLOffset := l * pageSize
		prevLOffset := (l - 1) * pageSize

		for i := range m {
			for j := 0; j+stepC+(1<<(l-1)) <= n; j++ {
				v1 := stGcd2D[prevLOffset+i*maxC+j]
				v2 := stGcd2D[prevLOffset+i*maxC+j+stepC]
				stGcd2D[currLOffset+i*maxC+j] = gcd(v1, v2)
			}
		}
	}

	for k := 1; k < wR; k++ {
		stepR := 1 << (k - 1)
		currKOffset := k * kStride
		prevKOffset := (k - 1) * kStride

		for l := range wC {
			lOffset := l * pageSize
			currBase := currKOffset + lOffset
			prevBase := prevKOffset + lOffset

			for i := 0; i+stepR+(1<<(k-1)) <= m; i++ {
				for j := 0; j+(1<<l) <= n; j++ {
					v1 := stGcd2D[prevBase+i*maxC+j]
					v2 := stGcd2D[prevBase+(i+stepR)*maxC+j]
					stGcd2D[currBase+i*maxC+j] = gcd(v1, v2)
				}
			}
		}
	}
}

func queryGcd2D(r1, r2, c1, c2 int) int {
	k := bits.Len(uint(r2-r1)) - 1
	l := bits.Len(uint(c2-c1)) - 1

	offset := k*kStride + l*pageSize
	lenR := 1 << k
	lenC := 1 << l

	v1 := stGcd2D[offset+r1*maxC+c1]
	v2 := stGcd2D[offset+(r2-lenR)*maxC+c1]
	v3 := stGcd2D[offset+r1*maxC+(c2-lenC)]
	v4 := stGcd2D[offset+(r2-lenR)*maxC+(c2-lenC)]

	return gcd(gcd(v1, v2), gcd(v3, v4))
}