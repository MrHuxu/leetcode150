use super::list::*;

struct Solution;

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
                    Some(Box::new(ListNode {
                        val: node1.val,
                        next: Solution::merge_two_lists(node1.next, Some(node2)),
                    }))
                } else {
                    Some(Box::new(ListNode {
                        val: node2.val,
                        next: Solution::merge_two_lists(Some(node1), node2.next),
                    }))
                }
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::merge_two_lists(
            ListNode::new_by_vec(vec![1, 2, 4]),
            ListNode::new_by_vec(vec![1, 3, 4]),
        ),
        ListNode::new_by_vec(vec![1, 1, 2, 3, 4, 4])
    );
}
