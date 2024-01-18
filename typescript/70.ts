function climbStairs(n: number): number {
    if (n <= 2) return n;
    
    let dp1 = 1, dp2 = 2;
    for (let i = 3; i <= n; i++) {
        const tmp = dp2;
        dp2 = dp1 + dp2;
        dp1 = tmp;
    }
    return dp2
};

test('70', () => {
    expect(climbStairs(2)).toBe(2);
    expect(climbStairs(3)).toBe(3);
});