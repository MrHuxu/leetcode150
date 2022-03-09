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
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        Self::helper(-200, head)
    }

    fn helper(pre_val: i32, head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        match head {
            None => None,
            Some(node) => {
                if node.val == pre_val
                    || (node.next.is_some() && node.next.as_ref().unwrap().val == node.val)
                {
                    Self::helper(node.val, node.next)
                } else {
                    Some(Box::new(ListNode {
                        val: node.val,
                        next: Self::helper(node.val, node.next),
                    }))
                }
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::delete_duplicates(ListNode::new_by_vec(vec![1, 2, 3, 3, 4, 4, 5])),
        ListNode::new_by_vec(vec![1, 2, 5])
    );
    assert_eq!(
        Solution::delete_duplicates(ListNode::new_by_vec(vec![1, 1, 1, 2, 3])),
        ListNode::new_by_vec(vec![2, 3])
    );
}
