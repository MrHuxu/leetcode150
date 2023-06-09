function convert(s: string, numRows: number): string {
    if (numRows === 1) return s;

    const m: Map<number, string> = new Map();
    let vertical = true;
    let line = 0;
    for (let char of s) {
        m.set(line, (m.get(line) || '') + char);
        line += vertical ? 1 : -1;

        if (line === numRows - 1) vertical = false;
        else if (line === 0) vertical = true;
    }

    let ret = '';
    for (let i = 0; i < numRows; i++) ret += m.get(i) || '';
    return ret;
};

test('6', () => {
    expect(convert("PAYPALISHIRING", 3)).toBe("PAHNAPLSIIGYIR");
});