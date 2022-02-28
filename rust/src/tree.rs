use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }

    #[inline]
    pub fn new_by_vec(vals: Vec<Option<i32>>) -> Option<Rc<RefCell<Self>>> {
        if vals.is_empty() || vals[0].is_none() {
            return None;
        }

        let mut full_vals: Vec<Vec<&Option<i32>>> = Vec::new();
        let mut idx = 0;
        while idx < vals.len() {
            if full_vals.is_empty() {
                full_vals.push(vec![&vals[idx]]);
                idx += 1;
                continue;
            }

            let mut level = Vec::new();
            for node in full_vals.last().unwrap() {
                if node.is_some() {
                    level.push(if idx < vals.len() && vals[idx].is_some() {
                        &vals[idx]
                    } else {
                        &None
                    });
                    level.push(if idx + 1 < vals.len() && vals[idx + 1].is_some() {
                        &vals[idx + 1]
                    } else {
                        &None
                    });

                    idx += 2;
                } else {
                    level.push(&None);
                    level.push(&None);
                }
            }

            full_vals.push(level);
        }

        Self::new_by_vec_helper(0, &full_vals.concat())
    }

    #[inline]
    fn new_by_vec_helper(idx: usize, vals: &Vec<&Option<i32>>) -> Option<Rc<RefCell<Self>>> {
        if idx >= vals.len() {
            return None;
        }

        match vals[idx] {
            None => None,
            Some(val) => Some(Rc::new(RefCell::new(TreeNode {
                val: *val,
                left: Self::new_by_vec_helper(idx * 2 + 1, vals),
                right: Self::new_by_vec_helper((idx + 1) * 2, vals),
            }))),
        }
    }
}
