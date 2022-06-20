# Application
Simple TODO service.
## Project strutrure
Project structure is based on https://github.com/golang-standards/project-layout
- `/assets` - image files
- `/build` - Dockerfile
- `/cmd` - main.go
- `/configs` - configuration example
- `/docs` - Swagger documentation
- `/internal` - internal packages
  - `/config` - configuration package
  - `/db` - database package
  - `/server` - server package
## Features:
### Authentication
- Register
- Login
---
### Tasks
- Get task
- Get tasks
- Get completed tasks
- Get uncompleted tasks
- Create task
- Update task
- Delete task
- Complete task
- Uncomplete task
- Add label to task
- Remove label from task
---
### Comments
- Get comment
- Get comments
- Create comment
- Update comment
- Delete comment
---
### Labels
- Get label
- Get labels
- Create label
- Update label
- Delete label
- Get tasks with label

## Run
1. `make run`
2. Wait for the server to start (it may fail a copuple of times if the database is not ready)
3. Visit Swagger - http://127.0.0.1:8080/swagger/index.html
4. Register with `username` and `password`
5. Login with `username` and `password` and receive a `token`
6. User token as `Authorization` header in all requests - `Authorization: Bearer <token>`
## API
http://127.0.0.1:8080/swagger/index.html

![image routes](https://raw.githubusercontent.com/vasjaj/todo/main/assets/routes.png)
## Database
For database actions I use https://gorm.io package since scale of project is small.

![image database](https://raw.githubusercontent.com/vasjaj/todo/main/assets/db.png)
## Testing
- Created test for `/login` handler in `internal/server/user_test.go` as an example.
- Run `make test`.
## Docker image - https://hub.docker.com/r/vasjajj/todo
## Deployment
### 1. CICD
   1. Lint pipeline with `make lint`
   2. Test pipeline with `make test`
   3. Build pipeline with `make docker-upload`
---
### 2. After service is Dockerized - `vasjajj/todo` there are several scenarios:
- Run application in Docker swarm
  1. Connect to the server
  2. Run application in Docker swarm using `docker-compose` file
- Run application in Kubernetes cluster
  1. Create service manifest using `vasjajj/todo` image
  2. Deploy service to cluster
- Deploy application to Heroku/Elastic Beanstalk/LightSail
  1. May skip Dockerization
  2. Deploy application code or docker image `vasjajj/todo` to one of those platforms
---
### 3. Post-deployment
1. Make use of TLS certificates
2. Make use of load balancer
## Possible improvements
- [ ] Check database indexes
- [ ] Check time operations
- [ ] Add planning features
- [ ] Add notifications features
- [ ] Add statistics features











