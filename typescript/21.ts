import { ListNode, compareList, buildList } from "./list";

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

function mergeTwoLists(list1: ListNode | null, list2: ListNode | null): ListNode | null {
    if (!list1 && !list2) return null;
    if (!list1) return list2;
    if (!list2) return list1;
    return new ListNode(
        list1.val < list2.val ? list1.val : list2.val,
        mergeTwoLists(
            list1.val < list2.val ? list1.next : list1,
            list1.val < list2.val ? list2 : list2.next
        )
    );
};

test('21', () => {
    expect(compareList(mergeTwoLists(
        buildList([1, 2, 4]), buildList([1, 3, 4])
    ), buildList([1, 1, 2, 3, 4, 4]))).toBeTruthy();
});