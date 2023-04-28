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

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    return helper(l1, l2, 0);
};

const helper = (l1: ListNode | null, l2: ListNode | null, carry: number): ListNode | null => {
    if (!l1 && !l2 && !carry) return null;
    const val = carry + (l1 ? l1.val : 0) + (l2 ? l2.val : 0);
    return new ListNode(val % 10, helper(l1 && l1.next, l2 && l2.next, Math.floor(val / 10)));
}

test('2', () => {
    expect(compareList(addTwoNumbers(
        buildList([2, 4, 3]), buildList([5, 6, 4])
    ), buildList([7, 0, 8]))).toBeTruthy();
});