# 红黑树
```go
import (
	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
)
```

### 外部接口汇总

| 类别 | 接口签名 | 功能描述 |
|------|----------|----------|
| **构造函数** | `New[K cmp.Ordered, V any]() *Tree[K, V]` | 创建红黑树，使用内置比较器 `cmp.Compare` （有序集合的话，`V` 存储空结构体 `struct{}`） |
| | `NewWith[K comparable, V any](comparator utils.Comparator[K]) *Tree[K, V]` | 创建红黑树，使用自定义比较器 |
| **Tree 基本操作** | `Put(key K, value V)` | 插入或更新键值对 |
| | `Get(key K) (value V, found bool)` | 根据键查找值，返回值和存在标志 |
| | `GetNode(key K) *Node[K, V]` | 根据键查找节点，返回节点指针（未找到返回 nil） |
| | `Remove(key K)` | 删除指定键的节点 |
| **Tree 状态查询** | `Empty() bool` | 判断树是否为空 |
| | `Size() int` | 返回树中节点总数 |
| | `Keys() []K` | 返回中序遍历的所有键 |
| | `Values() []V` | 返回中序遍历的所有值 |
| **Tree 极值查询** | `Left() *Node[K, V]` | 返回最小键节点（最左） |
| | `Right() *Node[K, V]` | 返回最大键节点（最右） |
| **Tree 查找邻近节点** | `Floor(key K) (floor *Node[K, V], found bool)` | 返回 ≤ key 的最大节点 |
| | `Ceiling(key K) (ceiling *Node[K, V], found bool)` | 返回 ≥ key 的最小节点 |
| **Tree 其他** | `Clear()` | 清空树 |
| | `String() string` | 返回树的可视化字符串 |
| **Node 方法** | `Size() int` | 计算以该节点为根的子树的节点数（动态递归） |
| | `String() string` | 返回节点键的字符串表示 |
| **导出字段** | `Tree.Root *Node[K, V]` | 树的根节点 |
| | `Tree.Comparator utils.Comparator[K]` | 键比较器 |
| | `Node.Key K` | 节点键 |
| | `Node.Value V` | 节点值 |
| | `Node.Left *Node[K, V]` | 左子节点 |
| | `Node.Right *Node[K, V]` | 右子节点 |
| | `Node.Parent *Node[K, V]` | 父节点 |