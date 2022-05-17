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
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut ret = Vec::new();
        Self::helper(root.as_ref(), &mut ret, 0);
        ret
    }

    fn helper(root: Option<&Rc<RefCell<TreeNode>>>, ret: &mut Vec<Vec<i32>>, depth: usize) {
        if let Some(node) = root {
            if depth >= ret.len() {
                ret.push(Vec::new());
            }

            ret[depth].push(node.borrow().val);
            Self::helper(node.borrow().left.as_ref(), ret, depth + 1);
            Self::helper(node.borrow().right.as_ref(), ret, depth + 1);
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::level_order(TreeNode::new_by_vec(vec![
            Some(3),
            Some(9),
            Some(20),
            None,
            None,
            Some(15),
            Some(7)
        ])),
        vec![vec![3], vec![9, 20], vec![15, 7]]
    );
}
