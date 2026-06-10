# 三笠 の 算法模板


leetcode 入门一年，leetcode 搜索 [忧郁小眼神](https://leetcode.cn/u/youyu-eyes/) \
leetcode 周赛 2300 Guardian，详见周赛号 [三笠](https://leetcode.cn/u/mikasha/)

整理一下常用的算法竞赛模板及对应的例题\
codeforces 题单 与 Python/C++ 部分模板见 [codeforces-classification](https://github.com/Youyu-eyes/codeforces-classification)

> 大部分是基础算法\
  图论，高级数据结构，数论与计算几何（除凸包外）没有涉及\
  字符串算法完全没有涉及

## Go 语言基础
[Go 语言快速入门](https://gobyexample-cn.github.io/)\
[Go 语言常用函数](https://github.com/Youyu-eyes/algorithm_go/blob/master/Go_NOTE/note.md)\
[Go 语言数据结构](https://github.com/emirpasic/gods)

## 算法分类

- 基础算法
  - 滑动窗口
    - 定长滑窗
    - 不定长滑窗
      - [越长越合法](https://leetcode.cn/problems/number-of-substrings-containing-all-three-characters/)
      - [越短越合法](https://leetcode.cn/problems/subarray-product-less-than-k/description/)
      - [恰好型滑窗](https://leetcode.cn/problems/binary-subarrays-with-sum/solutions/3972630/xiao-bai-si-lu-qia-hao-xing-hua-dong-chu-pk3z/)

  - 二分
    - 二分答案（最小化最大/最大化最小）
  - 双指针
    - 分组循环
    - [双指针优化二分](https://leetcode.cn/problems/longest-balanced-substring-after-one-swap/solutions/3950956/xiao-bai-si-lu-qian-zhui-he-fen-lei-tao-2e77c/)

- [位运算](https://github.com/Youyu-eyes/algorithm_go/tree/master/bitwise)
  - [位运算与集合论](https://leetcode.cn/discuss/post/3571304/cong-ji-he-lun-dao-wei-yun-suan-chang-ji-enve/)
  - [LogTrick](https://github.com/Youyu-eyes/algorithm_go/blob/master/bitwise/log_trick.go)
  - [线性基](https://github.com/Youyu-eyes/algorithm_go/blob/master/bitwise/xorBasis.go)
  - 拆位法
  - 试填法
  - 结论
    - [区间按位或](https://github.com/Youyu-eyes/codeforces-classification/blob/main/bitwise_operation/conclusion/bitwise_OR_of_interval.md)
    - [英文字符与位运算](https://github.com/Youyu-eyes/codeforces-classification/blob/main/bitwise_operation/conclusion/string_bit.md)

- [数据结构](https://github.com/Youyu-eyes/algorithm_go/tree/master/Data_Structrue)
  - [双端队列](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/deque.go)
  - [单调栈](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/monotonic_stack.go)
  - [懒删除堆](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/lazy_heap.go)
  - [树状数组](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/fenwick_tree.go)
  - 线段树
    - [普通线段树](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/segment_tree.go)
    - [Lazy 线段树](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/segment_tree_lazy.go)
    - [李超线段树](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/segment_tree_lichao.go)
    - 线段树动态开点
    - [可持久化线段树（主席树）](https://leetcode.cn/problems/minimum-operations-to-equalize-subarrays/solutions/3845357/zhong-wei-shu-tan-xin-ke-chi-jiu-hua-xia-etpv/)
  - [ST 表](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/sparse_table.go)
    - 一维 ST 表
    - 二维 ST 表
    - ST 表下标版本
    - [fast ST 表](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/sparse_table_fast.go)
  - [字典树](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/Trie.go)
    - 普通字典树
    - 01 字典树
    - 可持久化字典树
  - 根号算法
    - [分块](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/sqrt_decomposition.go)
    - [莫队算法](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/Mo.go)
      - 普通莫队
      - 带修莫队
      - 树上莫队
      - [回滚莫队](https://leetcode.cn/problems/threshold-majority-queries/description/)

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
    - [滚动最值优化](https://github.com/Youyu-eyes/codeforces-classification/blob/main/DP/DP_Optimization/Rolling_Extremum_Optimization/1945D_Seraphim_the_Owl.py)
    - [前缀和优化](https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-ii/solutions/3972184/xiao-bai-si-lu-dai-biao-ti-qian-zhui-he-mmtnj/)
    - 数据结构优化
      - [ST 表优化 RMQ](https://github.com/Youyu-eyes/codeforces-classification/blob/main/DP/DP_Optimization/Data_Structrue_Optimization/372C_Watching_Fireworks_is_Fun.go)
      - 单调栈优化
      - [单调队列优化](https://leetcode.cn/problems/maximum-sum-of-m-non-overlapping-subarrays-i/)
      - 线段树/树状数组优化
        - [二维偏序](https://leetcode.cn/problems/best-team-with-no-conflicts/solutions/3970519/xiao-bai-si-lu-er-wei-pian-xu-pai-xu-xia-v3uc/)
        - [李超线段树优化](https://github.com/Youyu-eyes/algorithm_go/blob/master/Data_Structrue/segment_tree_lichao.go)
          - [例题](https://leetcode.cn/problems/climbing-stairs-ii/solutions/3968349/xiao-bai-si-lu-jian-dan-dp-li-chao-xian-rxrkg/)
    - [凸包优化](https://github.com/Youyu-eyes/algorithm_go/blob/master/DP/Convex_Hull_Trick.go)
      - [例题](https://leetcode.cn/problems/minimum-partition-score/description/)
    - 矩阵快速幂优化
    - [WQS 二分优化](https://github.com/Youyu-eyes/algorithm_go/blob/master/DP/WQS.go)
      - [例题](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/solutions/536396/yi-chong-ji-yu-wqs-er-fen-de-you-xiu-zuo-x36r/)
      - [笔记](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/solutions/3981282/xiao-bai-si-lu-bi-ji-xi-lie-qian-xi-wqse-nhw1/)
    - CDQ 分治优化

- [图论](https://github.com/Youyu-eyes/algorithm_go/tree/master/graph)
  - 最短路
    - 单源最短路
      - [0-1 BFS](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/bfs_01.go)
      - [Dijkstra](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/dijkstra.go)
      - [bellman-ford](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/bellman_ford.go)
      - [SPFA](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/SPFA.go)
        - 差分约束系统
    - 全源最短路
      - [floyd](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/floyd.go)
    - 同余最短路
    - 分层图最短路 
  - [并查集](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/union_find.go)
    - 普通并查集
    - 中介并查集
    - 带权并查集
    - 置换环问题
      - [置换环分解](https://github.com/Youyu-eyes/codeforces-classification/blob/main/graph_theory/permutation_cycle/cycle_decomposition.md)
  - [拓扑序](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/topo_sort.go)
  - [二分图染色](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/bipartite.go)
  - 欧拉路径/回路
    - Hierholzer
  - 连通分量
    - Tarjan
  - [最小生成树](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/minimum_spanning_tree.go)
  - 网络流
    - 最大最小费用流
    - 二分图最大匹配
    - 带权二分图最大完美匹配
  - LCA 最近公共祖先
    - [树上倍增](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/LCA_tree.go)
    - 树链剖分
      - 轻重链剖分
  - [DFS 时间戳](https://github.com/Youyu-eyes/algorithm_go/blob/master/graph/DFS_time_stamp.go)

- [数学](https://github.com/Youyu-eyes/algorithm_go/blob/master/math/math.go)
  - [模运算的世界](https://leetcode.cn/discuss/post/3584387/fen-xiang-gun-mo-yun-suan-de-shi-jie-dan-7xgu/)
  - 预处理
    - 素数（埃氏筛）
    - 质因数分解
    - 因子
    - 阶乘及逆元
  - 质数判别
  - 离散化
  - 二维凸包
  - 矩阵乘法
  - 快速幂
    - 矩阵快速幂
  - 组合数学
    - 卡特兰数

- 贪心
  - [相邻不同系列问题·结论](https://leetcode.cn/problems/minimum-amount-of-time-to-fill-cups/description/)
  - [中位数贪心](https://leetcode.cn/problems/minimum-amount-of-time-to-fill-cups/description/)
  - [交换论证法](https://leetcode.cn/problems/minimum-processing-time/description/)
  - 反悔贪心

- 其他
  - [io模板](https://github.com/Youyu-eyes/algorithm_go/blob/master/io/io.go)
  - [严格众数·扩展摩尔投票](https://leetcode.cn/problems/majority-element/solutions/3978521/xiao-bai-si-lu-mo-er-tou-piao-fu-misra-g-18fe/)