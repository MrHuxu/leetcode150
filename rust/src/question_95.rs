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
    pub fn generate_trees(n: i32) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
        if n == 0 {
            return vec![];
        }

        Self::helper(1, n)
    }

    fn helper(start: i32, end: i32) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
        if start > end {
            return vec![None];
        }

        let mut ret = Vec::new();

        for i in start..=end {
            for left in Self::helper(start, i - 1).iter() {
                for right in Self::helper(i + 1, end).iter() {
                    ret.push(Some(Rc::new(RefCell::new(TreeNode {
                        val: i,
                        left: if left.is_some() {
                            Some(Rc::clone(left.as_ref().unwrap()))
                        } else {
                            None
                        },
                        right: if right.is_some() {
                            Some(Rc::clone(right.as_ref().unwrap()))
                        } else {
                            None
                        },
                    }))));
                }
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::generate_trees(1),
        vec![TreeNode::new_by_vec(vec![Some(1)])]
    );
}
