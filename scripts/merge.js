const { existsSync, readFileSync, writeFileSync } = require('fs');

for (let i = 1; i <= 150; i++) {
  if (!existsSync(`./problems/${i}`)) {
    continue;
  }

  const solution = readFileSync(`./problems/${i}/solution.md`).toString();

  const code = readFileSync(`./problems/${i}/main.go`).toString();
  const codeWithIndention = code.split('// code')[1].split('\n').map(line => '    ' + line).join('\n');

  writeFileSync(`./problems/${i}/index.md`, solution + codeWithIndention);
}