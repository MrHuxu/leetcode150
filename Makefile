test:
	go test -test.short -v -covermode=count ./problems/...

solve:
	mkdir problems/$(id)
	echo 'package leetcode150' > problems/$(id)/solution.go
	echo 'package leetcode150' > problems/$(id)/solution_test.go
	echo '# 题意\n\n\n# 解答' > problems/$(id)/explanation.md