import { ListNode, buildList, compareList } from "./list";

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function swapPairs(head: ListNode | null): ListNode | null {
    if (!head || !head.next) return head;

    const dummyHead = new ListNode(undefined, head);
    let prev = dummyHead;

    while (prev.next && prev.next.next) {
        const first = prev.next;
        const second = prev.next.next;

        first.next = second.next;
        second.next = first;
        prev.next = second;

        prev = first;
    }

    return dummyHead.next;
}

test('24', () => {
    expect(compareList(swapPairs(buildList([1, 2, 3, 4])), buildList([2, 1, 4, 3]))).toBeTruthy();
});