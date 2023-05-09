from typing import List, Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def buildList(vals: List[int]) -> Optional[ListNode]:
    match vals:
        case []:
            return None
        case [head, *tail]:
            head = ListNode(head)
            head.next = buildList(tail)
            return head


def compareList(l1: Optional[ListNode], l2: Optional[ListNode]) -> bool:
    while l1 and l2:
        if l1.val != l2.val:
            return False
        l1, l2 = l1.next, l2.next
    return l1 is None and l2 is None
