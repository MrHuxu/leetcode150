function strStr(haystack: string, needle: string): number {
    for (let i = 0; i <= haystack.length - needle.length; i++) {
        let mismatch = false;
        for (let j = 0; j < needle.length; j++) {
            if (haystack[i + j] !== needle[j]) {
                mismatch = true;
                break;
            }
        }
        if (!mismatch) return i;
    }
    return -1;
};

test('28', () => {
    expect(strStr('sadbutsad', 'sad')).toBe(0);
    expect(strStr('leetcode', 'leeto')).toBe(-1);
});