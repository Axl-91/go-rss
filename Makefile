up:
	docker compose up -d
down:
	docker compose down
build:
	go build -o ./tmp/go-rss
run: build
	./tmp/go-rss
