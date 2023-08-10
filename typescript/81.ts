function search2(nums: number[], target: number): boolean {
    if (target === nums[0]) return true;

    const rotateIdx = findRotateIdx(0, nums.length - 1, nums);
    if (rotateIdx === null) {
        return false;
    }

    let [left, right] = target > nums[0] ? [0, rotateIdx] : [rotateIdx + 1, nums.length - 1];
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (nums[mid] === target) return true;
        if (nums[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return false;
}

const findRotateIdx = (left: number, right: number, nums: number[]): number | null => {
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (mid === nums.length - 1 || nums[mid] > nums[mid + 1]) {
            return mid;
        }
        if (nums[mid] < nums[0]) {
            right = mid - 1;
            continue;
        }
        if (nums[mid] > nums[0]) {
            left = mid + 1;
            continue;
        }
        const idx = findRotateIdx(left, mid - 1, nums);
        if (idx !== null) return idx;
        return findRotateIdx(mid + 1, right, nums);
    }
    return null;
}

test('81', () => {
    expect(search2([1, 3, 5], 3)).toBeTruthy()
})