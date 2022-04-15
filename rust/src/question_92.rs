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
    pub fn reverse_between(
        head: Option<Box<ListNode>>,
        left: i32,
        right: i32,
    ) -> Option<Box<ListNode>> {
        let mut front = vec![0; left as usize - 1];
        let mut mid = vec![0; right as usize - left as usize + 1];
        let mut tail = Vec::new();

        Self::traverse(
            head.as_ref(),
            1,
            left,
            right,
            &mut front,
            &mut mid,
            &mut tail,
        );

        Self::build(front, mid, tail)
    }

    fn traverse(
        head: Option<&Box<ListNode>>,
        idx: i32,
        left: i32,
        right: i32,
        front: &mut Vec<i32>,
        mid: &mut Vec<i32>,
        tail: &mut Vec<i32>,
    ) {
        if head.is_none() {
            return;
        }

        let node = head.unwrap();
        if idx < left {
            front[idx as usize - 1] = node.val;
        } else if idx <= right {
            mid[(right - left - (idx - left)) as usize] = node.val;
        } else {
            tail.push(node.val);
        }

        Self::traverse(node.next.as_ref(), idx + 1, left, right, front, mid, tail)
    }

    fn build(front: Vec<i32>, mid: Vec<i32>, tail: Vec<i32>) -> Option<Box<ListNode>> {
        if !front.is_empty() {
            return Some(Box::new(ListNode {
                val: front[0],
                next: Self::build(front[1..].to_vec(), mid, tail),
            }));
        }

        if !mid.is_empty() {
            return Some(Box::new(ListNode {
                val: mid[0],
                next: Self::build(front, mid[1..].to_vec(), tail),
            }));
        }

        if !tail.is_empty() {
            return Some(Box::new(ListNode {
                val: tail[0],
                next: Self::build(front, mid, tail[1..].to_vec()),
            }));
        }

        None
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::reverse_between(ListNode::new_by_vec(vec![1, 2, 3, 4, 5]), 2, 4),
        ListNode::new_by_vec(vec![1, 4, 3, 2, 5])
    );

    assert_eq!(
        Solution::reverse_between(ListNode::new_by_vec(vec![5]), 1, 1),
        ListNode::new_by_vec(vec![5])
    );
}
