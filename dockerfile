FROM golang:1.17-alpine
WORKDIR /go/src/github.com/tarathep/tutorial-backend/
RUN go get -d -v github.com/gin-gonic/gin
RUN go get -d -v github.com/stretchr/testify
RUN go get -d -v go.mongodb.org/mongo-driver
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/tarathep/tutorial-backend/app ./
CMD ["./app"]