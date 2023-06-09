/**
 Do not return anything, modify matrix in-place instead.
 */
function rotate(matrix: number[][]): void {
    const n = matrix.length;
    if (n <= 1) return;

    for (let i = 0; i < (n + 1) / 2; i++) {
        for (let j = i; j < (n - 1 - i); j++) {
            const tmp = matrix[i][j];
            matrix[i][j] = matrix[n - 1 - j][i]
            matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j]
            matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i]
            matrix[j][n - 1 - i] = tmp
        }
    }
};

test('48', () => {
    const matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]];
    rotate(matrix);
    expect(matrix).toEqual([[7, 4, 1], [8, 5, 2], [9, 6, 3]]);
})