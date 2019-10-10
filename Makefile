install_dependencies:
	cd problems && go mod download
	cd website && go mod download && npm install

dev:
	cd website; \
	npm run dev

test:
	cd problems; \
	go test ./...

solve:
	cd problems; \
	mkdir $(id); \
	echo 'package leetcode150\n\n// code' > $(id)/solution.go; \
	echo 'package leetcode150\n\nimport (\n  "testing"\n\n  "github.com/stretchr/testify/assert"\n)\n\nfunc Test_func(t *testing.T) {\n  assert := assert.New(t)\n}' > $(id)/solution_test.go; \
	echo '# 题意\n\n\n# 解答' > $(id)/explanation.md