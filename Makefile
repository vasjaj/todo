docker-build:
	docker build -f build/Dockerfile . -t todo

docker-tag:
	docker tag todo vasjajj/todo:latest

docker-push:
	docker push vasjajj/todo:latest

docker-upload: docker-build docker-tag docker-push

docker-run-latest: docker-upload
	docker run --pull=always -p 8080:80 vasjajj/todo

make gen: swagger mocks

test:
	go test ./...

lint:
	golangci-lint run

swagger:
	swag init -g internal/server/server.go

mocks:
	mockgen -source internal/db/db.go -destination internal/db/mock/mock_db.go

run:
	docker-compose up
