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
    pub fn flatten(root: &mut Option<Rc<RefCell<TreeNode>>>) {
        if let Some(node) = root {
            Self::helper(Rc::clone(&node));
        }
    }

    fn helper(root: Rc<RefCell<TreeNode>>) -> Rc<RefCell<TreeNode>> {
        {
            let mut node = root.borrow_mut();
            if node.left.is_some() {
                let tail = Self::helper(Rc::clone(node.left.as_ref().unwrap()));
                tail.borrow_mut().right = node.right.clone();
                node.right = node.left.clone();
                node.left = None;
            }

            if node.right.is_some() {
                return Self::helper(Rc::clone(node.right.as_ref().unwrap()));
            }
        }

        root
    }
}

#[test]
fn test() {
    let mut root = TreeNode::new_by_vec(vec![
        Some(1),
        Some(2),
        Some(5),
        Some(3),
        Some(4),
        None,
        Some(6),
    ]);
    Solution::flatten(&mut root);
    assert_eq!(
        root,
        TreeNode::new_by_vec(vec![
            Some(1),
            None,
            Some(2),
            None,
            Some(3),
            None,
            Some(4),
            None,
            Some(5),
            None,
            Some(6)
        ])
    );
}
