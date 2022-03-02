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
    pub fn sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        Self::sort(head.as_ref())
    }

    fn sort(head: Option<&Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }

        let mut fast = &mut Some(Box::new(ListNode {
            val: 0,
            next: Some(Box::clone(head.as_ref().unwrap())),
        }));
        let mut slow_head = Some(Box::new(ListNode {
            val: 0,
            next: Some(Box::clone(head.as_ref().unwrap())),
        }));
        let mut slow = &mut slow_head;
        while let Some(fast_node) = fast.as_mut() {
            fast = &mut fast_node.next;
            if let Some(fast_node) = fast.as_mut() {
                fast = &mut fast_node.next;
            } else {
                break;
            }

            if fast.is_some() {
                slow = &mut slow.as_mut().unwrap().next;
            }
        }

        if slow.as_ref().unwrap().next.as_ref().eq(&head) {
            return Some(Box::clone(head.as_ref().unwrap()));
        }

        let slow_next = slow.as_mut().unwrap().next.take();
        let list_1 = Self::sort(slow_head.as_ref().unwrap().next.as_ref());
        let list_2 = Self::sort(slow_next.as_ref());

        Self::merge_2_lists(list_1, list_2)
    }

    fn merge_2_lists(
        list_1: Option<Box<ListNode>>,
        list_2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list_1, list_2) {
            (None, None) => None,
            (Some(node_1), None) => Some(node_1),
            (None, Some(node_2)) => Some(node_2),
            (Some(node_1), Some(node_2)) => {
                if node_1.val <= node_2.val {
                    Some(Box::new(ListNode {
                        val: node_1.val,
                        next: Self::merge_2_lists(node_1.next, Some(node_2)),
                    }))
                } else {
                    Some(Box::new(ListNode {
                        val: node_2.val,
                        next: Self::merge_2_lists(Some(node_1), node_2.next),
                    }))
                }
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::sort_list(ListNode::new_by_vec(vec![-1, 5, 3, 4, 0])),
        ListNode::new_by_vec(vec![-1, 0, 3, 4, 5])
    );
}
