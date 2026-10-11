package dp

import (
	"math/big"
)

// v0.y > 0
// v1.x 单调递增

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

// a.det(b) > 0 => a 到 b 逆时针
// a.det(b) < 0 => a 到 b 顺时针
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }

func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

// --- 上凸包 --- //
// uh := UpperHull{}

type UpperHull struct {
	hull []vec
}

func (uh *UpperHull) add(p vec) {
	for len(uh.hull) > 1 {
		back := uh.hull[len(uh.hull)-1]
		prev := uh.hull[len(uh.hull)-2]
		if back.sub(prev).detCmp(p.sub(back)) >= 0 {
			uh.hull = uh.hull[:len(uh.hull)-1]
		} else {
			break
		}
	}
	uh.hull = append(uh.hull, p)
}

// v0.x 单调递增
func (uh *UpperHull) queryMonotonic(p vec) int {
	for len(uh.hull) > 1 && p.dot(uh.hull[0]) <= p.dot(uh.hull[1]) {
		uh.hull = uh.hull[1:]
	}
	return p.dot(uh.hull[0])
}

// v0.x 单调递减
// func (uh *UpperHull) queryMonotonic(p vec) int {
//     for len(uh.hull) > 1 && p.dot(uh.hull[len(uh.hull)-1]) <= p.dot(uh.hull[len(uh.hull)-2]) {
//         uh.hull = uh.hull[:len(uh.hull)-1]
//     }
//     return p.dot(uh.hull[len(uh.hull)-1])
// }

func (uh *UpperHull) queryBinary(p vec) int {
	l, r := 0, len(uh.hull)-1
	for l < r {
		mid := (l + r) >> 1
		if p.dot(uh.hull[mid]) <= p.dot(uh.hull[mid+1]) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return p.dot(uh.hull[l])
}

func (uh *UpperHull) addFront(p vec) {
	for len(uh.hull) > 1 {
		front := uh.hull[0]
		next := uh.hull[1]

		if front.sub(p).detCmp(next.sub(front)) >= 0 {
			uh.hull = uh.hull[1:]
		} else {
			break
		}
	}
	uh.hull = append([]vec{p}, uh.hull...)
}

// uh1 max(x) < uh2 min(x)
func mergeUpper(uh1, uh2 UpperHull) UpperHull {
	if len(uh1.hull) <= len(uh2.hull) {
		for len(uh1.hull) > 0 {
			p := uh1.hull[len(uh1.hull)-1]
			uh1.hull = uh1.hull[:len(uh1.hull)-1]
			uh2.addFront(p)
		}
		return uh2
	} else {
		for len(uh2.hull) > 0 {
			p := uh2.hull[0]
			uh2.hull = uh2.hull[1:]
			uh1.add(p)
		}
		return uh1
	}
}

// --- 下凸包 --- //
// lh := LowerHull{}

type LowerHull struct {
	hull []vec
}

func (lh *LowerHull) add(p vec) {
	for len(lh.hull) > 1 {
		back := lh.hull[len(lh.hull)-1]
		prev := lh.hull[len(lh.hull)-2]
		if back.sub(prev).detCmp(p.sub(back)) <= 0 {
			lh.hull = lh.hull[:len(lh.hull)-1]
		} else {
			break
		}
	}
	lh.hull = append(lh.hull, p)
}

// v0.x 单调递减
func (lh *LowerHull) queryMonotonic(p vec) int {
	for len(lh.hull) > 1 && p.dot(lh.hull[0]) >= p.dot(lh.hull[1]) {
		lh.hull = lh.hull[1:]
	}
	return p.dot(lh.hull[0])
}

// v0.x 单调递增
// func (lh *LowerHull) queryMonotonic(p vec) int {
//     for len(lh.hull) > 1 && p.dot(lh.hull[len(lh.hull)-1]) >= p.dot(lh.hull[len(lh.hull)-2]) {
//         lh.hull = lh.hull[:len(lh.hull)-1]
//     }
//     return p.dot(lh.hull[len(lh.hull)-1])
// }

func (lh *LowerHull) queryBinary(p vec) int {
	l, r := 0, len(lh.hull)-1
	for l < r {
		mid := (l + r) >> 1
		if p.dot(lh.hull[mid]) >= p.dot(lh.hull[mid+1]) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return p.dot(lh.hull[l])
}

func (lh *LowerHull) addFront(p vec) {
	for len(lh.hull) > 1 {
		front := lh.hull[0]
		next := lh.hull[1]

		if front.sub(p).detCmp(next.sub(front)) <= 0 {
			lh.hull = lh.hull[1:]
		} else {
			break
		}
	}
	lh.hull = append([]vec{p}, lh.hull...)
}

// uh1 max(x) < uh2 min(x)
func mergeLower(lh1, lh2 LowerHull) LowerHull {
	if len(lh1.hull) <= len(lh2.hull) {
		for len(lh1.hull) > 0 {
			p := lh1.hull[len(lh1.hull)-1]
			lh1.hull = lh1.hull[:len(lh1.hull)-1]
			lh2.addFront(p)
		}
		return lh2
	} else {
		for len(lh2.hull) > 0 {
			p := lh2.hull[0]
			lh2.hull = lh2.hull[1:]
			lh1.add(p)
		}
		return lh1
	}
}
