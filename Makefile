install_dependencies:
	cd problems && go mod download
	cd website && go mod download && npm install

dev:
	cd website; \
	gin --port 8283 --appPort 15050 --all -i

test:
	cd problems; \
	go test ./...

fail:
	cd problems; \
	go test ./... | grep -E "FAIL" | grep -E "problems/[0-9]+"

FOLDER=$(id)_$(shell grep '"id": $(id)' -A 2 website/data.json | awk -F '"' '/slug/{print $$4}' | sed 's/-/_/g')

solve:
	cd problems; \
	mkdir $(FOLDER); \
	echo 'package main\n\nimport . "github.com/MrHuxu/leetcode150/problems/utils"\n\n// code' > $(FOLDER)/main.go; \
	echo 'package main\n\nimport (\n  "testing"\n\n  . "github.com/MrHuxu/leetcode150/problems/utils"\n\n  "github.com/stretchr/testify/assert"\n)\n\nfunc Test_func(t *testing.T) {\n  assert := assert.New(t)\n}' > $(FOLDER)/main_test.go; \
	echo '## 题意\n\n\n## 解答' > $(FOLDER)/solution.md