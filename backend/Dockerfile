FROM golang:1.24.1

WORKDIR /app

COPY /app/go.mod /app/
COPY /app/go.sum /app/

RUN go mod download

COPY /app/ /app/

ENTRYPOINT [ "go", "run", "main.go" ]