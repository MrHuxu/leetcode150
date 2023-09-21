function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
    const m = nums1.length;
    const n = nums2.length;

    const left = Math.floor((m + n + 1) / 2);
    const right = Math.floor((m + n + 2) / 2);

    return (getKth(nums1, 0, m - 1, nums2, 0, n - 1, left) + getKth(nums1, 0, m - 1, nums2, 0, n - 1, right)) / 2;
};

const getKth = (nums1: number[], start1: number, end1: number, nums2: number[], start2: number, end2: number, k: number): number => {
    const len1 = end1 - start1 + 1;
    const len2 = end2 - start2 + 1;

    if (len1 > len2) return getKth(nums2, start2, end2, nums1, start1, end1, k);
    if (!len1) return nums2[start2 + k - 1];
    if (k === 1) return Math.min(nums1[start1], nums2[start2]);

    const i = Math.min(len1, k / 2) - 1;
    const j = Math.min(len2, k / 2) - 1;

    if (nums1[start1 + i] > nums2[start2 + j])
        return getKth(nums1, start1, end1, nums2, start2 + j + 1, end2, k - j - 1);
    return getKth(nums1, start1 + i + 1, end1, nums2, start2, end2, k - i - 1);
}

test('4', () => {
    expect(findMedianSortedArrays([1, 3], [2])).toBe(2);
    expect(findMedianSortedArrays([1, 2], [3, 4])).toBe(2.5);
});