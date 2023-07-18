import unittest
from typing import Dict, Optional


class LRUNode:
    def __init__(self, key, value):
        self.key: int = key
        self.val: int = value
        self.prev: Optional[LRUNode] = None
        self.next: Optional[LRUNode] = None


class LRUCache:

    def __init__(self, capacity: int):
        self.cap: int = capacity
        self.size: int = 0
        self.data: Dict[int, Optional[LRUNode]] = {}
        self.head: Optional[LRUNode] = LRUNode(-1, -1)
        self.tail: Optional[LRUNode] = LRUNode(-1, -1)
        self.head.next = self.tail
        self.tail.prev = self.head

    def get(self, key: int) -> int:
        if key not in self.data:
            return -1

        node = self.data[key]
        if not node:
            return -1
        if node.prev:
            node.prev.next = node.next
        if node.next:
            node.next.prev = node.prev

        if not self.head:
            return -1
        tmp = self.head.next
        self.head.next = node
        node.prev = self.head
        node.next = tmp
        if tmp:
            tmp.prev = node

        return node.val

    def put(self, key: int, value: int) -> None:
        if key in self.data:
            node = self.data[key]
            if node:
                if node.prev:
                    node.prev.next = node.next
                if node.next:
                    node.next.prev = node.prev
                if not self.head or not self.head.next:
                    self.head = node
                    self.tail = node
                else:
                    tmp = self.head.next
                    self.head.next = node
                    node.prev = self.head
                    node.next = tmp
                    tmp.prev = node
                node.val = value
            return

        node = LRUNode(key, value)
        self.data[key] = node
        if not self.head:
            self.head = node
            self.tail = node
        else:
            tmp = self.head.next
            self.head.next = node
            node.prev = self.head
            node.next = tmp
            if tmp:
                tmp.prev = node

        if self.size == self.cap:
            if self.tail and self.tail.prev:
                del self.data[self.tail.prev.key]
                self.tail.prev = self.tail.prev.prev
            if self.tail and self.tail.prev:
                self.tail.prev.next = self.tail
        else:
            self.size += 1


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
class TestSolution(unittest.TestCase):
    def test(self):
        lru_cache = LRUCache(2)
        lru_cache.put(1, 1)
        lru_cache.put(2, 2)
        self.assertEqual(lru_cache.get(1), 1)
        lru_cache.put(3, 3)
        self.assertEqual(lru_cache.get(2), -1)
        lru_cache.put(4, 4)
        self.assertEqual(lru_cache.get(1), -1)
        self.assertEqual(lru_cache.get(3), 3)
        self.assertEqual(lru_cache.get(4), 4)
