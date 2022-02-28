use super::list::*;

struct Solution;

impl Solution {
    pub fn reverse_k_group(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        if k <= 1 || head.is_none() {
            return head;
        }

        let mut dummy_head = ListNode::new(0);
        let mut prev_next = &mut dummy_head.next;
        let mut head = head;

        let mut tmp = vec![];
        while let Some(mut node) = head.take() {
            head = node.next.take();
            tmp.push(node);
            if tmp.len() == k as usize {
                while tmp.len() > 0 {
                    *prev_next = tmp.pop();
                    prev_next = &mut prev_next.as_mut().unwrap().next;
                }
            }
        }
        while tmp.len() > 0 {
            *prev_next = Some(tmp.remove(0));
            prev_next = &mut prev_next.as_mut().unwrap().next;
        }
        *prev_next = None;

        dummy_head.next
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::reverse_k_group(ListNode::new_by_vec(vec![1, 2, 3, 4, 5]), 2),
        ListNode::new_by_vec(vec![2, 1, 4, 3, 5])
    );
}
