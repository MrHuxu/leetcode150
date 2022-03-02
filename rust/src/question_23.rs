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
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        if lists.len() == 0 {
            return None;
        }

        let mut lists = lists;
        while lists.len() > 1 {
            let mut next_lists: Vec<Option<Box<ListNode>>> = Vec::new();
            for i in (0..lists.len()).step_by(2) {
                if i < lists.len() - 1 {
                    next_lists.push(Solution::merge_2_lists(
                        lists[i].take(),
                        lists[i + 1].take(),
                    ));
                } else {
                    next_lists.push(lists[i].take());
                }
            }
            lists = next_lists;
        }

        lists[0].take()
    }

    fn merge_2_lists(
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
                        next: Self::merge_2_lists(node1.next, Some(node2)),
                    }))
                } else {
                    Some(Box::new(ListNode {
                        val: node2.val,
                        next: Self::merge_2_lists(Some(node1), node2.next),
                    }))
                }
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::merge_k_lists(vec![
            ListNode::new_by_vec(vec![1, 4, 5]),
            ListNode::new_by_vec(vec![1, 3, 4]),
            ListNode::new_by_vec(vec![2, 6]),
        ]),
        ListNode::new_by_vec(vec![1, 1, 2, 3, 4, 4, 5, 6])
    );
}
