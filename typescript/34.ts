function searchRange(nums: number[], target: number): number[] {
    const ret = [-1, -1];

    if (nums.length === 0) return ret;

    let left = 0;
    let right = nums.length - 1;
    while (left < right) {
        const mid = Math.floor((left + right) / 2);
        if (nums[mid] < target) left = mid + 1;
        else right = mid;
    }
    if (nums[Math.floor((left + right) / 2)] == target)
        ret[0] = Math.floor((left + right) / 2);

    left = 0;
    right = nums.length - 1;
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (nums[mid] > target) right = mid - 1;
        else left = mid + 1;
    }
    if (nums[Math.floor((left + right) / 2)] == target)
        ret[1] = Math.floor((left + right) / 2);

    return ret;
};

test('34', () => {
    expect(searchRange([5, 7, 7, 8, 8, 10], 8)).toStrictEqual([3, 4]);
    expect(searchRange([5, 7, 7, 8, 8, 10], 6)).toStrictEqual([-1, -1]);
    expect(searchRange([], 0)).toStrictEqual([-1, -1]);
    expect(searchRange([1], 0)).toStrictEqual([-1, -1]);
});