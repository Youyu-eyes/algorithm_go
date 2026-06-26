# set 有序集合

| 功能 | 方法 |
|------|------|
| 创建（默认升序） | `newTreap[K]()` |
| 创建（自定义比较） | `newTreapWith[K](comp func(a, b K) int)` |
| 插入元素 | `put(key K)` |
| 删除元素 | `delete(key K)` |
| 查找元素 | `find(key K) *node[K]` |
| 集合大小 | `size() int` |
| 判空 | `empty() bool` |
| 最小元素 | `min() *node[K]` |
| 最大元素 | `max() *node[K]` |
| 前驱（小于 key 的最大元素） | `prev(key K) *node[K]` |
| 后继（大于 key 的最小元素） | `next(key K) *node[K]` |
| 第 k 小（0-based） | `kth(k int) *node[K]` |
| 第一个 ≥ key 的下标 | `lowerBoundIndex(key K) int` |
| 第一个 > key 的下标 | `upperBoundIndex(key K) int` |