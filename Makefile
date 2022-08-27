docker-build:
	docker build -f build/Dockerfile . -t seamless

docker-tag:
	docker tag todo vasjajj/seamless:latest

docker-push:
	docker push vasjajj/seamless:latest

docker-upload: docker-build docker-tag docker-push

docker-run-latest: docker-upload
	docker run --pull=always -p 8080:80 vasjajj/seamless

test:
	go test ./...

lint:
	golangci-lint run


mocks:
	mockgen -source internal/db/db.go -destination internal/db/mock/mock_db.go

run:
	docker-compose up
