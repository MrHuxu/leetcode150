GITBOOK_VERSION=3.2.3

install_dependencies:
	npm install gitbook-cli -g && gitbook fetch ${GITBOOK_VERSION}

build:
	node scripts/merge.js
	sed -i -e 's/solution/index/g' SUMMARY.md
	rm .gitignore
	gitbook build

clear_build:
	git checkout .gitignore
	rm */*/index.md
	sed -i -e 's/index/solution/g' SUMMARY.md

dev:
	gitbook serve

test:
	cd problems; \
	go test ./...

solve:
	cd problems; \
	mkdir $(id); \
	echo 'package main\n\nimport . "github.com/MrHuxu/leetcode150/problems/utils"\n\n// code' > $(id)/main.go; \
	echo 'package main\n\nimport (\n  "testing"\n\n  . "github.com/MrHuxu/leetcode150/problems/utils"\n\n  "github.com/stretchr/testify/assert"\n)\n\nfunc Test_func(t *testing.T) {\n  assert := assert.New(t)\n}' > $(id)/main_test.go; \
	echo '## 题意\n\n\n## 解答' > $(id)/solution.md