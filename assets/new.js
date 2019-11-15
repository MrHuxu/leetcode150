const { writeFileSync } = require('fs');
const data = require('./data.json');

const id = process.argv[2];
const { title, slug, difficulty } = data.find(d => id == d.id);
console.log(id, title, slug, difficulty);

const badges = {
  1 : 'https://img.shields.io/badge/-easy-green?style=flat-square',
  2 : 'https://img.shields.io/badge/-medium-yellow?style=flat-square',
  3 : 'https://img.shields.io/badge/-hard-red?style=flat-square',
};

const solution  = `# ${id}. ${title} ![badge](${badges[difficulty]})

[题目链接](https://leetcode.com/problems/${slug})

## 题意

## 解答

## 代码
`;
writeFileSync(`./problems/${id}/solution.md`, solution);

const main = `package main

import . "github.com/MrHuxu/leetcode150/problems/utils"

// code
`;
writeFileSync(`./problems/${id}/main.go`, main);

const mainTest = `package main

import (
  "testing"

  . "github.com/MrHuxu/leetcode150/problems/utils"

  "github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
    assert := assert.New(t)
}
`;
writeFileSync(`./problems/${id}/main_test.go`, mainTest);
