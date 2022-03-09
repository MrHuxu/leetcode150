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
//   }
// }
impl Solution {
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list1, list2) {
            (None, None) => None,
            (Some(node1), None) => Some(node1),
            (None, Some(node2)) => Some(node2),
            (Some(node1), Some(node2)) => {
                if node1.val < node2.val {
                    let mut node = Box::clone(&node1);
                    node.next = Self::merge_two_lists(node1.next, Some(node2));
                    Some(node)
                } else {
                    let mut node = Box::clone(&node2);
                    node.next = Self::merge_two_lists(Some(node1), node2.next);
                    Some(node)
                }
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::merge_two_lists(ListNode::new_by_vec(vec![1]), ListNode::new_by_vec(vec![2])),
        ListNode::new_by_vec(vec![1, 2])
    );
    assert_eq!(
        Solution::merge_two_lists(
            ListNode::new_by_vec(vec![1, 2, 4]),
            ListNode::new_by_vec(vec![1, 3, 4])
        ),
        ListNode::new_by_vec(vec![1, 1, 2, 3, 4, 4])
    );
}
