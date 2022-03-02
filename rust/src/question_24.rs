use super::list::*;

struct Solution;

// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
//
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//

impl Solution {
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }
        if head.as_ref().unwrap().next.is_none() {
            return head;
        }

        let mut curr = head.unwrap();
        let mut next = curr.next.unwrap();
        curr.next = Solution::swap_pairs(next.next);
        next.next = Some(curr);

        Some(next)
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::swap_pairs(ListNode::new_by_vec(vec![1, 2, 3, 4])),
        ListNode::new_by_vec(vec![2, 1, 4, 3])
    );
}
