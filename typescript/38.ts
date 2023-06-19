function countAndSay(n: number): string {
    let ret = '1';
    for (let i = 1; i < n; i++) {
        let next = '';
        let ch = ret[0];
        let count = 1;
        for (let j = 1; j < ret.length; j++) {
            if (ret[j] === ch) {
                count++;
                continue;
            }

            next = `${next}${count}${ch}`;
            ch = ret[j];
            count = 1;
        }
        ret = `${next}${count}${ch}`;
    }
    return ret;
};

test('38', () => {
    expect(countAndSay(4)).toBe('1211');
});