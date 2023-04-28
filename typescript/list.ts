export class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val === undefined ? 0 : val);
        this.next = (next === undefined ? null : next);
    }
}

export const buildList = (vals: number[]): ListNode | null => vals.length ? new ListNode(
    vals[0],
    buildList(vals.slice(1))
) : null;

export const compareList = (l1: ListNode | null, l2: ListNode | null): boolean => {
    if (!l1 && !l2) return true;
    if (!l1 || !l2) return false;
    return l1.val === l2.val && compareList(l1.next, l2.next);
};