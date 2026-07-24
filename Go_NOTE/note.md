[Go 语言快速入门](https://gobyexample-cn.github.io/)
# io 效率测试
[Codeforces 13E-Holes](https://codeforces.com/problemset/problem/13/E)
[Codeforces 372C-Watching Fireworks is Fun](https://codeforces.com/contest/372/problem/C)

# 类型转换
| 函数 | 作用 |
| --- | --- |
| FormatInt(i int64, base int) string | 整数按进制转字符串 |
| FormatFloat(f float64, fmt byte, prec, bitSize int) string | 浮点数格式化 |
| strconv.Itoa(i int) string | int → 十进制字符串 |
| Atoi(s string) (int, error) | 十进制字符串 → int |
| ParseInt(s string, base int, bitSize int) (int64, error) | 解析整数（带进制、位宽） |
| ParseFloat(s string, bitSize int) (float64, error) | 解析浮点数 |
| ParseBool(s string) (bool, error) | 解析布尔值 |

# 排序
| 功能 | Go 语言实现方式 (使用 `slices` 包，需要导入 `"cmp"` 包，Go 1.21+) |
| :--- | :--- |
| **普通排序** | `slices.Sort(nums)` |
| **倒序排序** | `slices.Sort(nums); slices.Reverse(nums)` |
| **按二维切片第一项升序排序** | `slices.SortFunc(nums, func(a, b []int) int { return cmp.Compare(a[0], b[0]) })` |
| **按二维切片第一项降序排序** | `slices.SortFunc(nums, func(a, b []int) int { return cmp.Compare(b[0], a[0]) })` |
| **双关键字排序**<br>（如第一项升序，第二项升序） | `slices.SortFunc(queries, func(a, b []int) int { ` <br> `return cmp.Or(` <br> `cmp.Compare(a[0], b[0]),` <br> `cmp.Compare(a[1], b[1]),`<br>`)`<br>`})` |
| **元组（pair）双关键字排序**<br>（first 升序，second 降序）需要提前定义 pair 结构体 | `slices.SortFunc(pairs, func(a, b Pair) int {`<br> `if c := cmp.Compare(a.First, b.First); c != 0 { return c }`<br> `return cmp.Compare(b.Second, a.Second)`<br>`})` |
| **两个切片一一对应，按 nums1 排序，nums2 同步移动** 需要提前定义 pair 结构体 | 构造 `pairs` 切片后：<br>`slices.SortFunc(pairs, func(a, b Pair) int { return cmp.Compare(a.First, b.First) })`<br>然后拆解回两个切片。 |
| **使用静态映射表进行稳定排序**<br>（如按元音出现次数的负值排序） | 使用 `slices.SortStableFunc`：<br>`slices.SortStableFunc(vowels, func(a, b byte) int {`<br> `return cmp.Compare(-cnt[mp[a]], -cnt[mp[b]])`<br>`})`<br> |

# 二分
$s$ 中 $\ge v$ 的第一个数
`sort.SearchInts(s, v) -> int`

$s$ 中 $\le v$ 的第一个数
`slices.BinarySearch(s, v) -> (int, bool)`
手写二分注意写成 `left + (right - left) >> 1` 防止溢出

# 内部函数
没有 `accumulate`，`sum`，`abs`
对整个切片取 `max` 为 `slices.Max`
没有 `pop` 函数，需要自己实现

# 翻转
切片：`slices.Reverse(nums)`
字符串没有相关内置函数，手动实现一个
```go
func ReverseString(s string) string {
    var b strings.Builder
    b.Grow(len(s))
    for i := len(s) - 1; i >= 0; i-- {
        b.WriteByte(s[i])
    }
    return b.String()
}
```

# 运算优先级
在 Go 中，位运算符与算术运算符的优先级从高到低排列如下：

| 优先级 | 包含的运算符 | 说明 |
|:---:|:---|:---|
| 最高（乘除级） | `*` `/` `%` `<<` `>>` `&` `&^` | 移位、按位与、按位清除与乘除同级 |
| 中间（加减级） | `+` `-` `\|` `^` | 按位或、按位异或与加减同级 |
| 比较级 | `==` `!=` `<` `<=` `>` `>=` | |
| 逻辑与 | `&&` | |
| 最低 | `\|\|` | |

所有二元运算符都是**左结合**（同优先级时从左向右计算）。建议在混合使用算术与位运算时用括号明确运算顺序。

# 分支变量

go 的变量，如果是在分支中初始化的，如
``` go
if !vis[x] {
    y := g[x]
}
```
这里的 $y$ 只能在这个 `if` 分支中成立

# 字符串

`strings.Repeat(a, n)`：相当于 python 字符串乘法，生成长度为 $n$ 的 `a` 字符串，注意 `a` 是 **string** 而非 **byte**，需要 "" 而不是 ''；
`unicode.IsUpper(c)`：判断 `c` 是否为大写字母，注意 `c` 的数据类型是 `rune` 而非 `byte`

`string(slices)`：将切片转换成 `string`，注意切片类型一定是 `byte` 切片

# 位运算

`bits.Len(uint(x))`       位长度
`bits.OnesCount(uint(x))` 置位数

`&^` 位清空运算，`c &^32` 等价于 `c & ~32`，~表示补码

# 数据结构

## 堆
Go 语言堆的库函数需要首先实现 $5$ 个函数，具体见 [堆与队列库函数合集（Python/Go）](https://leetcode.cn/discuss/post/3708255/python-dui-he-dui-lie-de-ku-han-shu-he-j-dbgn/)

# 杂项 & 技巧

## leetcode 函数题
如果在循环内一定有输出，不一定需要最后 `return`，则需要在最后加上 `panic("impossible")`，详见 [lc3613](https://leetcode.cn/problems/minimize-maximum-component-cost/description/)

请不要用哈希表统计数字出现次数，请一定一定用数组！！！哈希表的常数达到无法想象