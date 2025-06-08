FROM golang:1.24.1

WORKDIR /app

COPY backend/app/go.mod /app/
COPY backend/app/go.sum /app/

RUN go mod download

COPY backend/app/ /app/

ENTRYPOINT [ "go", "run", "main.go" ]