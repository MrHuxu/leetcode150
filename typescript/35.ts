function searchInsert(nums: number[], target: number): number {
    if (nums[0] > target) return 0;
    if (nums[nums.length - 1] < target) return nums.length;

    let left = 0;
    let right = nums.length - 1;
    while (left < right) {
        const mid = Math.floor((left + right) / 2);
        if (nums[mid] < target) left = mid + 1;
        else right = mid;
    }
    return Math.floor((left + right) / 2);
};

test('35', () => {
    expect(searchInsert([1, 3, 5, 6], 5)).toBe(2);
    expect(searchInsert([1, 3, 5, 6], 2)).toBe(1);
    expect(searchInsert([1, 3, 5, 6], 7)).toBe(4);
    expect(searchInsert([1, 3], 3)).toBe(1);
});
