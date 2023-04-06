copy-files:
	cp -rf ../documents data/
	mkdir -p data/solutions
	cp -rf ../go data/solutions/
	rm -f data/solutions/go/go.*
	cp -rf ../rust data/solutions/
	cp -rf ../java data/solutions/
	cp -rf ../typescript data/solutions/

vercel-dev: copy-files
	vercel dev

tailwind-watch:
	tailwindcss --watch -o assets/built.css

go-dev: copy-files
	PORT=3000 go run main.go

docker-build:
	docker build -t lc150 .

docker-run:
	docker run -p 15050:15050 lc150