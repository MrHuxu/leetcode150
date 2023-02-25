function maxProfit(prices: number[]): number {
    let res = 0;
    let minPrice = Number.MAX_SAFE_INTEGER;
    for (let price of prices) {
        minPrice = Math.min(minPrice, price);
        res = price >= minPrice ? Math.max(res, price - minPrice) : res;
    }
    return res;
};

test('121', () => {
    expect(maxProfit([3, 9, 20, null, null, 15, 7])).toBe(5);
    expect(maxProfit([1, null, 2])).toBe(0);
});