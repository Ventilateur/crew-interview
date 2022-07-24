
IMAGE_NAME = crew-interview

docker-build:
	docker build -t $(IMAGE_NAME):$$(git rev-parse --short HEAD) .

up:
	docker compose up -d

down:
	docker compose down
