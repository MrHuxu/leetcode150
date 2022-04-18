use super::tree::*;

struct Solution;

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
    pub fn build_tree(inorder: Vec<i32>, postorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        if inorder.is_empty() {
            return None;
        }

        let node_val = postorder.last().unwrap();
        let mut idx = 0;
        for (i, val) in inorder.iter().enumerate() {
            if val == node_val {
                idx = i;
                break;
            }
        }

        Some(Rc::new(RefCell::new(TreeNode {
            val: *node_val,
            left: Self::build_tree(inorder[..idx].to_vec(), postorder[..idx].to_vec()),
            right: Self::build_tree(
                inorder[idx + 1..].to_vec(),
                postorder[idx..postorder.len() - 1].to_vec(),
            ),
        })))
    }
}

#[test]
fn name() {
    assert_eq!(
        Solution::build_tree(vec![9, 3, 15, 20, 7], vec![9, 15, 7, 20, 3]),
        TreeNode::new_by_vec(vec![
            Some(3),
            Some(9),
            Some(20),
            None,
            None,
            Some(15),
            Some(7)
        ])
    );
}
