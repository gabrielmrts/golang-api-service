FROM golang:1.17 as development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 4000
CMD reflex -r '\.go$' go run main.go --start-service