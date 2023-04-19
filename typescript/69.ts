function mySqrt(x: number): number {
    if (x === 0 || x === 1) return x;

    let left = 1;
    let right = Math.ceil(x / 2);
    while (left <= right) {
        const mid = Math.ceil((left + right) / 2);
        const pow = mid * mid;
        if (pow === x) {
            return mid;
        } else if (pow > x) {
            right = mid - 1;
        } else {
            left = mid + 1;
        }
    }

    return right;
};

test('69', () => {
    expect(mySqrt(4)).toBe(2);
    expect(mySqrt(8)).toBe(2);
});