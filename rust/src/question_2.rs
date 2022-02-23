use super::list::*;

struct Solution;

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        Solution::add(l1, l2, 0)
    }

    fn add(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
        carry: i32,
    ) -> Option<Box<ListNode>> {
        match (l1, l2) {
            (None, None) => {
                if carry == 0 {
                    None
                } else {
                    Some(Box::new(ListNode { val: 1, next: None }))
                }
            }
            (Some(val1), None) => {
                let sum = val1.val + carry;
                Some(Box::new(ListNode {
                    val: sum % 10,
                    next: Solution::add(val1.next, None, sum / 10),
                }))
            }
            (None, Some(val2)) => {
                let sum = val2.val + carry;
                Some(Box::new(ListNode {
                    val: sum % 10,
                    next: Solution::add(None, val2.next, sum / 10),
                }))
            }
            (Some(val1), Some(val2)) => {
                let sum = val1.val + val2.val + carry;
                Some(Box::new(ListNode {
                    val: sum % 10,
                    next: Solution::add(val1.next, val2.next, sum / 10),
                }))
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::add_two_numbers(
            Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode {
                    val: 4,
                    next: Some(Box::new(ListNode { val: 3, next: None })),
                })),
            })),
            Some(Box::new(ListNode {
                val: 5,
                next: Some(Box::new(ListNode {
                    val: 6,
                    next: Some(Box::new(ListNode { val: 4, next: None })),
                })),
            }))
        ),
        Some(Box::new(ListNode {
            val: 7,
            next: Some(Box::new(ListNode {
                val: 0,
                next: Some(Box::new(ListNode { val: 8, next: None })),
            })),
        })),
    );
}
