function spiralOrder(matrix: number[][]): number[] {
    const ret: number[] = [];

    const m = matrix.length;
    const n = matrix[0].length;
    for (let level = 0; m > 0 && ret.length < m * n; level++) {
        for (let i = level; i < n - level; i++) ret.push(matrix[level][i]);
        for (let i = level + 1; i < m - level - 1; i++) ret.push(matrix[i][n - level - 1]);
        for (let i = n - level - 1; i >= level && level !== m - level - 1; i--) ret.push(matrix[m - level - 1][i]);
        for (let i = m - level - 2; i > level && level !== n - level - 1; i--) ret.push(matrix[i][level]);
    }

    return ret;
};

test('54', () => {
    expect(spiralOrder([[1, 2, 3], [4, 5, 6], [7, 8, 9]])).toStrictEqual([1, 2, 3, 6, 9, 8, 7, 4, 5]);
    expect(spiralOrder([[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]])).toStrictEqual([1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]);
})