use super::list::*;
use super::tree::*;

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
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn sorted_list_to_bst(head: Option<Box<ListNode>>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut head_ref = &head;
        let mut len = 0;
        while let Some(node) = head_ref {
            len += 1;
            head_ref = &node.next;
        }

        let mut head = &head;
        Self::helper(&mut head, len)
    }

    pub fn helper(head: &mut &Option<Box<ListNode>>, len: usize) -> Option<Rc<RefCell<TreeNode>>> {
        if len == 0 {
            return None;
        }

        let left = Self::helper(head, len / 2);
        let node = head.as_ref().unwrap();
        let mut tree_node = TreeNode::new(node.val);
        tree_node.left = left;
        *head = &node.next;
        tree_node.right = Self::helper(head, len - len / 2 - 1);

        Some(Rc::new(RefCell::new(tree_node)))
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::sorted_list_to_bst(ListNode::new_by_vec(vec![-10, -3, 0, 5, 9])),
        TreeNode::new_by_vec(vec![Some(0), Some(-3), Some(9), Some(-10), None, Some(5)])
    );
}
