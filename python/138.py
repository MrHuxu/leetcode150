from typing import Dict, Optional


class Node:
    def __init__(self, x: int, next, random):
        self.val = int(x)
        self.next: Node = next
        self.random: Node = random


"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""


class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        map_node_to_enw_node: Dict[Node, Node] = {}

        dummy_head = Node(0, None, None)

        iter = dummy_head
        while head:
            if head not in map_node_to_enw_node:
                map_node_to_enw_node[head] = Node(head.val, None, None)

            if head.next:
                if head.next not in map_node_to_enw_node:
                    map_node_to_enw_node[head.next] = Node(
                        head.next.val, None, None)
                map_node_to_enw_node[head].next = map_node_to_enw_node[head.next]

            if head.random:
                if head.random not in map_node_to_enw_node:
                    map_node_to_enw_node[head.random] = Node(
                        head.random.val, None, None)
                map_node_to_enw_node[head].random = map_node_to_enw_node[head.random]

            iter.next = map_node_to_enw_node[head]

            iter, head = iter.next, head.next

        return dummy_head.next
