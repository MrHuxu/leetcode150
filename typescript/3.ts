function lengthOfLongestSubstring(s: string): number {
    let ret = 0;

    let front = 0;
    const cnt: Map<string, number> = new Map();
    for (let i = 0; i < s.length; i++) {
        const ch = s[i];
        if (!cnt.has(ch)) {
            ret = Math.max(ret, i - front + 1);
        } else {
            while (cnt.has(ch)) {
                cnt.delete(s[front]);
                front++;
            }
        }
        cnt.set(ch, 1);
    }

    return ret;
}

test('3', () => {
    expect(lengthOfLongestSubstring('bbbbbb')).toBe(1);
    expect(lengthOfLongestSubstring('tmmzuxt')).toBe(5);
});