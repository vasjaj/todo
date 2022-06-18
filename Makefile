docker-build:
	docker build -f build/Dockerfile . -t todo

docker-tag:
	docker tag todo vasjajj/todo:latest

docker-push:
	docker push vasjajj/todo:latest

docker-upload: docker-build docker-tag docker-push

docker-run-latest: docker-upload
	docker run --pull=always -p 8080:80 vasjajj/todo

lint:
	golangci-lint run

init-swagger:
	cd cmd
	swag init -g ./internal/server/server.go -g ./internal/server/task.go