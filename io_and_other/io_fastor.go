package main

// import (
// 	"os"
// )

// const (
// 	MOD = 1_000_000_007
// 	inf = 0x3f3f3f3f
// )

// var (
// 	rd    func() int
// 	rs    func() []byte
// 	wInt  func(int)
// 	wByte func(byte)
// 	flush func()
// )

// // n := rd()
// // wInt(ans)
// // wByte('\n')

// func solve() {

// }

// func main() {
// 	initIO()
// 	defer flush()

// 	T := rd()
// 	for i := 0; i < T; i++ {
// 		solve()
// 	}
// }

// func initIO() {
// 	const eof = 0
// 	buf := make([]byte, 4096)
// 	_i, _n := 0, 0

// 	// 底层读单字节
// 	rc := func() byte {
// 		if _i == _n {
// 			_n, _ = os.Stdin.Read(buf)
// 			if _n == 0 { // EOF
// 				return eof
// 			}
// 			_i = 0
// 		}
// 		b := buf[_i]
// 		_i++
// 		return b
// 	}

// 	// 读整数
// 	rd = func() (x int) {
// 		neg := false
// 		b := rc()
// 		for ; b != eof && b <= ' '; b = rc() { // 跳过所有空白字符
// 		}
// 		if b == eof {
// 			return
// 		}
// 		if b == '-' {
// 			neg = true
// 			b = rc()
// 		}
// 		for ; b >= '0' && b <= '9'; b = rc() {
// 			x = x*10 + int(b&15)
// 		}
// 		if neg {
// 			return -x
// 		}
// 		return
// 	}

// 	// 读连续字符串 (优化版：支持任何非空白字符)
// 	rs = func() (s []byte) {
// 		b := rc()
// 		for ; b != eof && b <= ' '; b = rc() { // 跳过空白字符
// 		}
// 		if b == eof {
// 			return
// 		}
// 		for ; b > ' '; b = rc() { // 读取直到遇到空格或换行
// 			s = append(s, b)
// 		}
// 		return
// 	}

// 	// 输出引擎
// 	const outputN int = 1e6
// 	const intWidth = 20
// 	outS := make([]byte, 0, outputN*(intWidth+2))
// 	tmpS := [intWidth]byte{}

// 	wInt = func(x int) {
// 		if x == 0 {
// 			outS = append(outS, '0', ' ')
// 			return
// 		}
// 		if x < 0 {
// 			x = -x
// 			outS = append(outS, '-')
// 		}
// 		p := len(tmpS)
// 		for ; x > 0; x /= 10 {
// 			p--
// 			tmpS[p] = '0' | byte(x%10)
// 		}
// 		outS = append(outS, tmpS[p:]...)
// 		outS = append(outS, ' ')
// 	}

// 	wByte = func(b byte) {
// 		outS = append(outS, b)
// 	}

// 	flush = func() {
// 		os.Stdout.Write(outS)
// 		outS = outS[:0] // 刷新后清空数组长度，防止重复输出
// 	}
// }
// func abs[T int | int64 | float64](x T) T {
// 	if x >= 0 { return x }
// 	return -x
// }
// func gcd(a, b int) int {
// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	return a
// }