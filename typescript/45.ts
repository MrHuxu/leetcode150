function jump(nums: number[]): number {
    const dp: number[] = Array(nums.length).fill(0);

    for (let i = 0; i < nums.length; i++) {
        for (let j = i + 1; j <= i + nums[i] && j < nums.length; j++) {
            dp[j] = dp[j] === 0 ? dp[i] + 1 : Math.min(dp[j], dp[i] + 1);
        }
    }

    return dp[dp.length - 1];
};

test('45', () => {
    expect(jump([2, 3, 1, 1, 4])).toBe(2);
    expect(jump([2, 3, 0, 1, 4])).toBe(2);
});
