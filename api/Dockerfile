FROM golang:1.22-alpine

WORKDIR /app
COPY main.go .
RUN go mod init api && go mod tidy
RUN go get github.com/gin-gonic/gin
EXPOSE 8080
CMD ["go", "run", "main.go"]
