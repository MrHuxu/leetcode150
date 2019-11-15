GITBOOK_VERSION=3.2.3

install_dependencies:
	npm install gitbook-cli -g && gitbook fetch ${GITBOOK_VERSION}

build:
	node scripts/merge.js
	sed -i -e 's/solution/index/g' SUMMARY.md
	rm .gitignore

	gitbook build

	sed -i -e 's/gitbook\/images\/apple-touch-icon-precomposed-152.png/https:\/\/raw.githubusercontent.com\/MrHuxu\/img-repo\/master\/leetcode150\/favicon.png/g' _book/index.html
	sed -i -e 's/gitbook\/images\/favicon.ico/https:\/\/raw.githubusercontent.com\/MrHuxu\/img-repo\/master\/leetcode150\/favicon.png/g' _book/index.html
	sed -i -e 's/<\/head>/<link href="https:\/\/cdn.bootcss.com\/highlight.js\/9.15.10\/styles\/tomorrow.min.css" rel="stylesheet"><\/head>/g' _book/index.html
	sed -i -e 's/<\/body>/<script src="https:\/\/cdn.bootcss.com\/highlight.js\/9.15.10\/highlight.min.js"><\/script><script src="https:\/\/cdn.bootcss.com\/highlight.js\/9.15.10\/languages\/go.min.js"><\/script><script>hljs.initHighlightingOnLoad();<\/script><\/body>/g' _book/index.html
	sed -i -e 's/GitBook<\/title>/xhu<\/title>/g' _book/index.html

	sed -i -e 's/..\/..\/gitbook\/images\/apple-touch-icon-precomposed-152.png/https:\/\/raw.githubusercontent.com\/MrHuxu\/img-repo\/master\/leetcode150\/favicon.png/g' _book/problems/*/index.html
	sed -i -e 's/..\/..\/gitbook\/images\/favicon.ico/https:\/\/raw.githubusercontent.com\/MrHuxu\/img-repo\/master\/leetcode150\/favicon.png/g' _book/problems/*/index.html
	sed -i -e 's/<\/head>/<link href="https:\/\/cdn.bootcss.com\/highlight.js\/9.15.10\/styles\/tomorrow.min.css" rel="stylesheet"><\/head>/g' _book/problems/*/index.html
	sed -i -e 's/<\/body>/<script src="https:\/\/cdn.bootcss.com\/highlight.js\/9.15.10\/highlight.min.js"><\/script><script src="https:\/\/cdn.bootcss.com\/highlight.js\/9.15.10\/languages\/go.min.js"><\/script><script>hljs.initHighlightingOnLoad();<\/script><\/body>/g' _book/problems/*/index.html
	sed -i -e 's/GitBook<\/title>/xhu<\/title>/g' _book/problems/*/index.html

	git checkout .gitignore
	sed -i -e 's/index/solution/g' SUMMARY.md
	rm -rf */*/index.md SUMMARY.md-e

dev:
	gitbook serve

test:
	cd problems; \
	go test ./...

solve:
	mkdir problems/$(id); \
	node scripts/new.js $(id)