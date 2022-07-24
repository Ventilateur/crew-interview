
IMAGE_NAME = crew-interview

docker-build:
	docker build -t $(IMAGE_NAME):$$(git rev-parse --short HEAD) .

up:
	docker compose up -d

down:
	docker compose down

seed:
	go run main.go seed \
		--mongo-uri mongodb://root:root@localhost:27017 \
		--database crew \
		--collection talents \
		--crew-uri https://hiring.crew.work/v1/talents
