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
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        Self::add(l1, l2, 0)
    }

    fn add(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
        carry: i32,
    ) -> Option<Box<ListNode>> {
        match (l1, l2) {
            (None, None) => {
                if carry == 0 {
                    None
                } else {
                    Some(Box::new(ListNode { val: 1, next: None }))
                }
            }
            (Some(val1), None) => {
                let sum = val1.val + carry;
                Some(Box::new(ListNode {
                    val: sum % 10,
                    next: Self::add(val1.next, None, sum / 10),
                }))
            }
            (None, Some(val2)) => {
                let sum = val2.val + carry;
                Some(Box::new(ListNode {
                    val: sum % 10,
                    next: Self::add(None, val2.next, sum / 10),
                }))
            }
            (Some(val1), Some(val2)) => {
                let sum = val1.val + val2.val + carry;
                Some(Box::new(ListNode {
                    val: sum % 10,
                    next: Self::add(val1.next, val2.next, sum / 10),
                }))
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::add_two_numbers(
            ListNode::new_by_vec(vec![2, 4, 3]),
            ListNode::new_by_vec(vec![5, 6, 4])
        ),
        ListNode::new_by_vec(vec![7, 0, 8])
    )
}
