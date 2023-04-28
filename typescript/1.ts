function twoSum(nums: number[], target: number): number[] {
    const mapNumToPos: Map<number, number> = new Map();
    for (let i = 0; i < nums.length; i++) {
        if (mapNumToPos.has(target - nums[i])) {
            return [mapNumToPos.get(target - nums[i])!, i];
        }
        mapNumToPos.set(nums[i], i);
    }
    return [];
};

test('1', () => {
    expect(twoSum([2, 7, 11, 15], 9)).toStrictEqual([0, 1]);
});