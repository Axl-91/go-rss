up:
	docker compose up -d
down:
	docker compose down
run:
	go build -o ./tmp/go-rss && ./tmp/go-rss
