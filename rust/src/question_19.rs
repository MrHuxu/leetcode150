use super::list::*;

struct Solution;

impl Solution {
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }

        let mut len = 0;
        let mut tmp = head.as_ref();
        while tmp.is_some() {
            len += 1;
            tmp = tmp.unwrap().next.as_ref();
        }

        let mut ret = Some(Box::new(ListNode { val: 0, next: head }));
        let mut modify_ret = ret.as_mut();
        for _i in 0..len - n {
            modify_ret = modify_ret.unwrap().next.as_mut();
        }
        let mut node = modify_ret.unwrap();
        let next = node.next.as_ref();
        node.next = match next.unwrap().next.as_ref() {
            None => None,
            Some(val) => Some(Box::clone(val)),
        };

        ret.unwrap().next
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::remove_nth_from_end(
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 2,
                    next: Some(Box::new(ListNode {
                        val: 3,
                        next: Some(Box::new(ListNode {
                            val: 4,
                            next: Some(Box::new(ListNode { val: 5, next: None }))
                        })),
                    })),
                })),
            })),
            2
        ),
        Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode {
                    val: 3,
                    next: Some(Box::new(ListNode { val: 5, next: None })),
                })),
            })),
        })),
    );
    assert_eq!(
        Solution::remove_nth_from_end(
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 2,
                    next: Some(Box::new(ListNode {
                        val: 3,
                        next: Some(Box::new(ListNode {
                            val: 4,
                            next: Some(Box::new(ListNode { val: 5, next: None }))
                        })),
                    })),
                })),
            })),
            1
        ),
        Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode {
                    val: 3,
                    next: Some(Box::new(ListNode { val: 4, next: None })),
                })),
            })),
        })),
    );
    assert_eq!(
        Solution::remove_nth_from_end(Some(Box::new(ListNode { val: 1, next: None })), 1),
        None,
    );
    assert_eq!(
        Solution::remove_nth_from_end(
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode { val: 2, next: None }))
            })),
            1
        ),
        Some(Box::new(ListNode { val: 1, next: None }))
    );
}
