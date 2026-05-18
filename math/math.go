package math

import (
	"math"
	"sort"
)

const (
	MOD    = 1_000_000_007
	inf    = math.MaxInt
	ll_inf = math.MaxInt
	MX     = 1_000_001   // 根据题目调整大小
)

// ------- 全局预处理变量 ------- //
var (
	isPrimeArr []bool
	primes     []int
	lpf        []int
	divisors   [][]int
	F          []int
	INV_F      []int
)

// func init() 会在 main 函数之前自动执行，用于预处理各类数据
func init() {
	// 1. 初始化预处理数组大小
	isPrimeArr = make([]bool, MX)
	lpf = make([]int, MX)
	divisors = make([][]int, MX)
	F = make([]int, MX)
	INV_F = make([]int, MX)

	// 2. 埃氏筛预处理质数
	for i := range isPrimeArr {
		isPrimeArr[i] = true
	}
	isPrimeArr[0], isPrimeArr[1] = false, false
	for i := 2; i < MX; i++ {
		if isPrimeArr[i] {
			primes = append(primes, i)
			for j := i * i; j < MX; j += i {
				isPrimeArr[j] = false // j 是质数 i 的倍数
			}
		}
	}

	// 3. 质因数分解 (LPF) 预处理
	for i := 2; i < MX; i++ {
		if lpf[i] == 0 { // i 是质数
			for j := i; j < MX; j += i {
				if lpf[j] == 0 { // 首次访问 j
					lpf[j] = i
				}
			}
		}
	}

	// 4. 预处理因子
	for i := 1; i < MX; i++ {
		for j := i; j < MX; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}

	// 5. 预处理阶乘及逆元
	F[0] = 1
	for i := 1; i < MX; i++ {
		F[i] = F[i-1] * i % MOD
	}
	INV_F[MX-1] = qpow(F[MX-1], MOD-2, MOD)
	for i := MX - 1; i > 0; i-- {
		INV_F[i-1] = INV_F[i] * i % MOD
	}
}

// ------- 基础数论 ------- //

// 时间复杂度 O(sqrt(n))
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2 // 1 不是质数
}

type Pair struct {
	p, e int
}

// 质因数分解
// 时间复杂度 O(log x)
func primeFactorization(x int) []Pair {
	var res []Pair
	for x > 1 {
		p := lpf[x]
		e := 1
		for x /= p; x%p == 0; x /= p {
			e++
		}
		res = append(res, Pair{p, e})
	}
	return res
}

// ------- 离散化 ------- //

// 返回离散化后的数组以及去重后的元素个数 m
func discretize(arr []int) ([]int, int) {
	sortedUnique := make([]int, len(arr))
	copy(sortedUnique, arr)
	sort.Ints(sortedUnique)

	// unique 去重
	n := 0
	for i := 0; i < len(sortedUnique); i++ {
		if i == 0 || sortedUnique[i] != sortedUnique[i-1] {
			sortedUnique[n] = sortedUnique[i]
			n++
		}
	}
	sortedUnique = sortedUnique[:n]
	m := len(sortedUnique)

	discretized := make([]int, len(arr))
	for i, x := range arr {
		discretized[i] = sort.SearchInts(sortedUnique, x)
	}
	return discretized, m
}

// ------- 计算几何：凸包 ------- //

type Vec struct {
	x, y int
}

func (v Vec) sub(b Vec) Vec {
	return Vec{v.x - b.x, v.y - b.y}
}

// 注意：如果坐标范围超过 3 * 10^9，这里的乘法可能会超过 64位 int 上限，此时可以使用 math/bits.Mul64 处理
// a.det(b) > 0 => a 到 b 逆时针
// a.det(b) < 0 => a 到 b 顺时针
func (v Vec) det(b Vec) int {
	return v.x*b.y - v.y*b.x
}

func (v Vec) dot(b Vec) int {
	return v.x*b.x + v.y*b.y
}

// Andrew 算法，计算 points 的凸包（逆时针顺序）
// 时间复杂度 O(n log n)
func convexHull(points []Vec) []Vec {
	// 排序优先按 x，相同时按 y
	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})

	var q []Vec

	// 计算下凸包
	for _, p := range points {
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}

	// 计算上凸包
	lowerSize := len(q)
	for i := len(points) - 2; i >= 0; i-- {
		p := points[i]
		for len(q) > lowerSize && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}

	// 此时首尾是同一个点 points[0]，需要去掉
	if len(q) > 0 {
		q = q[:len(q)-1]
	}

	return q
}

// ------- 矩阵 ------- //

type matrix [][]int

// 返回矩阵 a 和矩阵 b 相乘的结果，若 mod > 0 则取模
func mul(a, b matrix, mod int) matrix {
	n, m := len(a), len(b[0])
	c := make(matrix, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, m)
		for k := 0; k < len(a[0]); k++ {
			if a[i][k] == 0 {
				continue
			}
			aik := a[i][k]
			for j := 0; j < m; j++ {
				c[i][j] += aik * b[k][j]
				if mod > 0 {
					c[i][j] %= mod
				}
			}
		}
	}
	return c
}

// 计算 A^n * f0，若 mod > 0 则所有乘法取模
func powMul(a matrix, n int, f0 matrix, mod int) matrix {
	res := make(matrix, len(f0))
	for i := range f0 {
		res[i] = make([]int, len(f0[i]))
		copy(res[i], f0[i])
	}

	for n > 0 {
		if n&1 == 1 {
			res = mul(a, res, mod) // 左乘 A
		}
		a = mul(a, a, mod) // A 自乘
		n >>= 1
	}
	return res
}

// 快速幂模板，支持负指数计算逆元
func qpow(x, n, mod int) int {
	ans := 1
	base := x
	exp := n

	if mod > 0 {
		base %= mod
	}

	if exp < 0 {
		if mod == 0 {
			return 0
		}
		// 模质数下求逆元（费马小定理）
		inv, b, p := 1, base, mod-2
		for p > 0 {
			if p&1 == 1 {
				inv = (inv * b) % mod
			}
			b = (b * b) % mod
			p >>= 1
		}
		base = inv
		exp = -exp
	}

	for exp > 0 {
		if exp&1 == 1 {
			ans *= base
			if mod > 0 {
				ans %= mod
			}
		}
		base *= base
		if mod > 0 {
			base %= mod
		}
		exp >>= 1
	}
	return ans
}

// ------- 组合数学 ------- //

// 从 n 个数中选 m 个数的方案数
func comb(n, m int) int {
	if m < 0 || m > n {
		return 0
	}
	return F[n] * INV_F[m] % MOD * INV_F[n-m] % MOD
}

// 返回第 n 个卡特兰数 (C_n) 模 MOD
func catalan(n int) int {
	if n == 0 {
		return 1
	}
	// C_n = C(2n, n) / (n+1)
	C_2n_n := comb(2*n, n)
	inv_n1 := qpow(n+1, -1, MOD) // 求 (n+1) 的逆元
	return C_2n_n * inv_n1 % MOD
}
