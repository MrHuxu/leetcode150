function firstMissingPositive(nums: number[]): number {
    const n = nums.length;
    for (let i = 0; i < n; i++) {
        const num = nums[i];
        if (num > 0 && num <= n && nums[num - 1] !== num) {
            nums[i] = nums[num - 1];
            nums[num - 1] = num;
            i--;
        }
    }
    for (let i = 0; i < n; i++) {
        if (nums[i] !== i + 1) return i + 1;
    }
    return n + 1;
};

test('41', () => {
    expect(firstMissingPositive([1, 2, 0])).toBe(3)
    expect(firstMissingPositive([3, 4, -1, 1])).toBe(2)
    expect(firstMissingPositive([7, 8, 9, 11, 12])).toBe(1)
})