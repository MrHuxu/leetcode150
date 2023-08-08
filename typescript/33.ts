function search(nums: number[], target: number): number {
    let left = 0;
    let right = nums.length - 1;
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (nums[mid] < nums[0])
            right = mid - 1;
        else
            left = mid + 1;
    }

    console.log(left, right)

    const rotate_idx = Math.floor((left + right) / 2)
    if (target >= nums[0]) {
        left = 0;
        right = rotate_idx;
    } else {
        left = rotate_idx + 1;
        right = nums.length - 1;
    }

    console.log(left, right)

    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (nums[mid] === target)
            return mid;
        else if (nums[mid] < target)
            left = mid + 1;
        else
            right = mid - 1;
    }

    return -1;
};

test('33', () => {
    expect(search([3, 1], 1)).toBe(1);
    expect(search([4, 5, 6, 7, 0, 1, 2], 0)).toBe(4)
    expect(search([4, 5, 6, 7, 0, 1, 2], 3)).toBe(-1)
    expect(search([1], 0)).toBe(-1)
});