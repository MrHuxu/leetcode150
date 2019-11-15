const { existsSync, writeFileSync } = require('fs');

const data = require('./data.json');

const badges = {
  1 : 'https://img.shields.io/badge/-easy-green?style=flat-square',
  2 : 'https://img.shields.io/badge/-medium-yellow?style=flat-square',
  3 : 'https://img.shields.io/badge/-hard-red?style=flat-square',
};

for (let d of data) {
  const { id, title, slug, difficulty } = d;

  if (!existsSync(`./problems/${id}`)) {
    continue;
  }


  const content = `# ${id}. ${title} ![badge](${badges[difficulty]})

[题目链接](https://leetcode.com/problems/${slug})

## 题意

## 解答

## 代码

`;
  writeFileSync(`./problems/${id}/solution.md`, content);
}