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