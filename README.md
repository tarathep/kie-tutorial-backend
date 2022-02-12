# tutorial-backend

 Tutorial backend rest api work with [frontend vue.js](https://github.com/tarathep/tutorial-frontend)


Framework By [Gin Web Framework](https://github.com/gin-gonic/gin)


## prerequisites

- Go 10.x have been Installed
- MongoDB Available
  ```bash
  docker run -it --rm --name mongodb -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=password mongo:latest
  ```

## go dependencies

set go module
```bash
go env -w GO111MODULE=auto
```

gin gonic

```bash
go get -u github.com/gin-gonic/gin
```

mongodb

```bash
go get -u go.mongodb.org/mongo-driver/mongo
```

testify

```bash
go get -u github.com/stretchr/testify
```

### Go Module

set module
```bash
go env -w GO111MODULE=on

go mod init

go mod tidy

go mod vendor

go mod verify
```


## Environment Variables

- ``MONGODB_CONNECTION_STRING``
  - default : mongodb://127.0.0.1:27017
  - example : mongodb://root:password@192.168.1.102:27017
- ``PORT``
  - default : 8089



## Run

```bash
go run main.go
```


## APIs

- ``GET : /api/tutorials``

  response body
  ```json
  [{"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}]
  ```
- ``GET : /api/tutorials/602aa1e04f3b51804eca6917``

  response body
  ```json
  {"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}
  ```
- ``POST : /api/tutorials``
  
  request body
  ```json
  {"title":"xx","description":"xx Description"}
  ```
  
  response body
  ```json
  {"code":"200","message":"Inserted a single document Success"}
  ```

- ``PUT : /api/tutorials``
  
  request body
  ```json
  {"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}
  ```
  
  response body
  ```json
  {"code":"200","message":"Updated  a single document Success"}
  ```

- ``DELETE : /api/tutorials``
   
  response body
  ```json
  {"code":"200","message":"All deleted"}
  ```
- ``DELETE : /api/tutorials/602aa1e04f3b51804eca6917``
    
  response body
  ```json
  {"code":"200","message":"Deleted id 602aa1e04f3b51804eca6917"}
  ```

## Unit Test

test

```bash
go test
```

test with create cover profile
```bash
go test -coverprofile coverage.out ./...
```

export report html
```bash
go tool cover -html coverage.out -o report.html
```

