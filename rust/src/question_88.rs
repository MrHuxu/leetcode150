struct Solution;

impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut m = m as usize;
        let mut n = n as usize;
        loop {
            match (m, n) {
                (0, 0) => break,
                (0, _) => {
                    nums1[m + n - 1] = nums2[n - 1];
                    n -= 1;
                }
                (_, 0) => {
                    nums1[m + n - 1] = nums1[m - 1];
                    m -= 1;
                }
                (_, _) => {
                    nums1[m + n] = if nums1[m - 1] > nums2[n - 1] {
                        m -= 1;
                        nums1[m]
                    } else {
                        n -= 1;
                        nums2[n]
                    }
                }
            }
        }
    }
}

#[test]
fn test() {
    let mut nums1 = vec![1, 2, 3, 0, 0, 0];
    let mut nums2 = vec![2, 5, 6];
    Solution::merge(&mut nums1, 3, &mut nums2, 3);
    assert_eq!(nums1, vec![1, 2, 2, 3, 5, 6]);

    nums1 = vec![1];
    nums2 = vec![];
    Solution::merge(&mut nums1, 1, &mut nums2, 0);
    assert_eq!(nums1, vec![1]);
}
