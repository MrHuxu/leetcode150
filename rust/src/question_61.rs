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
    pub fn rotate_right(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        if head.as_ref().is_none() {
            return None;
        }

        let mut nodes = Vec::new();
        let mut node_cnt = 0;

        let mut tmp = head.as_ref();
        while let Some(node) = tmp {
            node_cnt += 1;
            nodes.push(node);
            tmp = node.next.as_ref();
        }

        if k % node_cnt as i32 == 0 {
            return head;
        }

        nodes.rotate_right(k as usize % node_cnt);
        Self::helper(nodes)
    }

    fn helper(nodes: Vec<&Box<ListNode>>) -> Option<Box<ListNode>> {
        match nodes[..] {
            [first, ..] => Some(Box::new(ListNode {
                val: first.val,
                next: Self::helper(nodes[1..].to_vec()),
            })),
            _ => None,
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::rotate_right(ListNode::new_by_vec(vec![1, 2, 3, 4, 5]), 2),
        ListNode::new_by_vec(vec![4, 5, 1, 2, 3])
    );
    assert_eq!(
        Solution::rotate_right(ListNode::new_by_vec(vec![0, 1, 2]), 4),
        ListNode::new_by_vec(vec![2, 0, 1])
    );
}
