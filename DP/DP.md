# DP 优化小结

1. 前缀和优化 DP
可以优化形如
$$
nf[i] = \sum_{j = left}^{right}(f[j])
$$
定义 $f$ 数组前缀和为 $Sf[i + 1] = Sf[i] + f[i]$，上述式子可以写成
$$
nSf[i + 1] - nSf[i] = Sf[right] - Sf[left] \\
\Rightarrow nSf[i + 1] = nSf[i] + Sf[right + 1] - Sf[left]
$$
如果 $right < i$，可以优化为一维。
时间复杂度通常为 $\mathcal{O}(kn)$
缺点：如果 $k$ 很大，很难优化掉 $k$.

2. 矩阵快速幂优化 DP
可以优化形如
$$
f[i][k] = \sum f[i - 1][t]
$$
将 $f[i]$ 写成列向量，总列数为 $k$ 的取值，将 $f[i - 1]$ 也写成列向量，总列数相等。
时间复杂度：$\mathcal{O}(k^3×logn)$，主要针对 $n$ 很大的情况，$k$ 通常不大，如果过大可以采用 $BM$ 算法优化数列求和。