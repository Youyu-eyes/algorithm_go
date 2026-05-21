# Go 语言算法模板

整理一下常用的算法竞赛模板及对应的例题

## 算法分类

- [动态规划](https://github.com/Youyu-eyes/algorithm_go/tree/master/DP)
  - 背包问题
    - [0 - 1背包](https://leetcode.cn/problems/target-sum/description/)
    - [完全背包](https://leetcode.cn/problems/coin-change/description/)
    - [分组背包](https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/description/)
    - [树形背包](https://leetcode.cn/problems/maximum-profit-from-trading-stocks-with-discounts/description/)
  - 经典线性 DP
    - [LIS 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/description/?envType=problem-list-v2&envId=HUyjbQzI)
    - [LCS 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/description/?envType=problem-list-v2&envId=HUyjbQzI)
  - 区间 DP
    - [中心扩展法](https://leetcode.cn/problems/longest-palindromic-substring/solutions/2958179/mo-ban-on-manacher-suan-fa-pythonjavacgo-t6cx/)
  - 划分型 DP
    - 求最多划分数
    - 约束划分个数
  - 状态机 DP
    - [买卖股票问题 · 系列](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/)
  - [数位 DP](https://github.com/Youyu-eyes/algorithm_go/blob/master/DP/digitDP.go)
    - 区间合法数字个数
    - 区间合法数字总价值
  - 树形 DP
    - [换根 DP](https://leetcode.cn/problems/maximum-subgraph-score-in-a-tree/solutions/3850874/huan-gen-dppythonjavacgo-by-endlesscheng-y5tw/)
  - 状压 DP
    - SOS DP 高维前缀和
  - 前后缀分解
  - [DP 输出具体方案](https://codeforces.com/contest/474/problem/E)
  - DP 优化
    - 滚动最值优化
    - [前缀和优化](https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-ii/solutions/3972184/xiao-bai-si-lu-dai-biao-ti-qian-zhui-he-mmtnj/)
    - 数据结构优化
      - ST 表优化 RMQ
      - 单调栈优化
      - 单调队列优化
      - 线段树/树状数组优化
        - [二维偏序](https://codeforces.com/problemset/problem/2167/G)
        - [李超线段树优化](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/segment_tree_lichao.go)
          - [例题](https://leetcode.cn/problems/climbing-stairs-ii/solutions/3968349/xiao-bai-si-lu-jian-dan-dp-li-chao-xian-rxrkg/)
    - 凸包优化
      - [例题](https://leetcode.cn/problems/minimum-partition-score/description/)
      - [凸包模板](https://github.com/Youyu-eyes/algorithm_go/blob/master/DP/Convex_Hull_Trick.go)
    - 矩阵快速幂优化
    - [WQS 二分优化](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/solutions/536396/yi-chong-ji-yu-wqs-er-fen-de-you-xiu-zuo-x36r/)
    - CDQ 分治优化