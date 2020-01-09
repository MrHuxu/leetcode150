## 题意

实现一个 LRU(Least Recently Used) 缓存, 这个缓存暴露两个函数:

- `Get(key)`: 从缓存中根据 key 拿到一个值, 如果不存在则返回 -1.
- `Put(key, value)`: 将一个 key, value 对加入到缓存中, 如果缓存大小超过容量限制, 则删除最久未被使用的 key.

两个函数的时间复杂度都需要是 O(1).

## 解答

首先使用 O(1) 来实现 Get 和 Put 很容易想到用 map, 但是对于缓存的容量控制光用 map 是不够的, 需要一些额外操作, 我们可以使用双端链表来实现, 其中链表的头节点为最近操作的节点, 尾节点为最久没操作的节点.

其中缓存本身的结构为:

    type LRUCache struct {
      cap  int                // 缓存容量
      size int                // 缓存大小
      data map[int]*LRUNode   // 数据 map
      head *LRUNode           // 双端链表头节点
      tail *LRUNode           // 双端链表尾节点
    }

而双端链表节点的结构为:

    type LRUNode struct {
      key  int                // 缓存的键
      val  int                // 缓存的值
      prev *LRUNode           // 链表中前一个相邻节点
      next *LRUNode           // 链表中后一个相邻节点
    }

首先是 Get 函数, 操作步骤如下:

1. 判断 key 是否在 data 里, 不在的话直接返回 -1;
2. key 在 data 里的话, 可以通过 key 拿到其对应的链表节点 node, 再通过修改指针的操作, 将 node 放到链表的开头, 再返回 node 的值.

这样就可以通过 O(1) 实现了 Get 操作, 并且 key 会在最近操作的位置.

然后是 Put 函数, 操作步骤如下:

1. key 在 data 里, 先获得起对应的节点 node
  - 更新 node 的值为 value;
  - 通过修改指针的操作, 将 node 放到链表的开头.
2. key 不在 data 里
  - 使用 key 和 value 建立一个新节点 newNode;
  - 将 newNode 直接放到链表的开头位置;
  - 如果当前缓存大小未达到容量, 直接对大小累加 1;
  - 如果当前缓存大小等于容量, 则删除最久未访问的 key, 即 tail 指针指向的链表尾节点, 并从 data 里删除尾节点的 key.

这样就可以通过 O(1) 实现了 Put 操作, 每一次通过 Put 所创建/更新的 key 都会在最近操作的位置.
  