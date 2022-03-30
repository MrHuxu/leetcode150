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
    pub fn partition(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
        let mut head = head.as_ref();
        let mut smaller_nodes = Vec::new();
        let mut larger_nodes = Vec::new();
        while let Some(node) = head {
            head = node.next.as_ref();
            if node.val < x {
                smaller_nodes.push(node);
            } else {
                larger_nodes.push(node);
            }
        }

        Self::merge(smaller_nodes, larger_nodes)
    }

    fn merge(
        smaller_nodes: Vec<&Box<ListNode>>,
        larger_nodes: Vec<&Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        if !smaller_nodes.is_empty() {
            let node = smaller_nodes[0];
            return Some(Box::new(ListNode {
                val: node.val,
                next: Self::merge(smaller_nodes[1..].to_vec(), larger_nodes),
            }));
        }

        if !larger_nodes.is_empty() {
            let node = larger_nodes[0];
            return Some(Box::new(ListNode {
                val: node.val,
                next: Self::merge(smaller_nodes, larger_nodes[1..].to_vec()),
            }));
        }

        None
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::partition(ListNode::new_by_vec(vec![1, 4, 3, 2, 5, 2]), 3),
        ListNode::new_by_vec(vec![1, 2, 2, 4, 3, 5])
    );
}
