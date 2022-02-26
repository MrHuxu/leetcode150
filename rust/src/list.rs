// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }

    #[inline]
    pub fn new_by_vec(vals: Vec<i32>) -> Option<Box<Self>> {
        match vals[..] {
            [first, ..] => Some(Box::new(ListNode {
                val: first,
                next: Self::new_by_vec(vals[1..].to_vec()),
            })),
            _ => None,
        }
    }
}
