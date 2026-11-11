# multiset 多重集合（有序列表）

| 功能 | 方法 |
|------|------|
| 创建（默认升序） | `newTreap[K]()` |
| 创建（自定义比较） | `newTreapWith[K](comp func(a, b K) int)` |
| 插入元素（增加计数） | `put(key K)` |
| 删除元素（减少计数，计数归零时移除） | `delete(key K)` |
| 查找元素 | `find(key K) *node[K]` |
| 集合大小（含重复计数） | `size() int` |
| 判空 | `empty() bool` |
| 最小元素 | `min() *node[K]` |
| 最大元素 | `max() *node[K]` |
| 前驱（小于 key 的最大元素） | `prev(key K) *node[K]` |
| 后继（大于 key 的最小元素） | `next(key K) *node[K]` |
| 第 k 小（0-based，考虑重复） | `kth(k int) *node[K]` |
| 第一个 ≥ key 的下标（考虑重复） | `lowerBoundIndex(key K) int` |
| 第一个 > key 的下标（考虑重复） | `upperBoundIndex(key K) int` |
| 节点键值 | `node.key` |
| 节点计数 | `node.keyCnt` |