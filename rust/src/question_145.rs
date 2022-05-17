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
    pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        Self::helper(root.as_ref())
    }

    fn helper(root: Option<&Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        if let Some(node) = root {
            let borrowed_node = node.borrow();
            vec![
                Self::helper(borrowed_node.left.as_ref()),
                Self::helper(borrowed_node.right.as_ref()),
                vec![borrowed_node.val],
            ]
            .concat()
        } else {
            vec![]
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::postorder_traversal(TreeNode::new_by_vec(vec![Some(1), None, Some(2), Some(3)])),
        vec![3, 2, 1]
    );
}
