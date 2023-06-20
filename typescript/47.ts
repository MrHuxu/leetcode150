function permuteUnique(nums: number[]): number[][] {
    nums.sort((a, b) => a - b);
    return dfs(nums, new Set());
};

const dfs = (nums: number[], used: Set<number>): number[][] => {
    const ret: number[][] = [];

    for (let i = 0; i < nums.length; i++) {
        if (used.has(i) || (i > 0 && nums[i] === nums[i - 1] && !used.has(i - 1))) continue;

        used.add(i);
        for (let next of dfs(nums, used)) ret.push([nums[i], ...next]);
        used.delete(i);
    }

    if (!ret.length) ret.push([]);

    return ret;
};

test('47', () => {
    expect(permuteUnique([1, 1, 2])).toEqual([
        [1, 1, 2],
        [1, 2, 1],
        [2, 1, 1],
    ]);
    expect(permuteUnique([1, 2, 3])).toEqual([
        [1, 2, 3],
        [1, 3, 2],
        [2, 1, 3],
        [2, 3, 1],
        [3, 1, 2],
        [3, 2, 1],
    ]);
});