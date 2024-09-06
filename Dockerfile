FROM golang:1.23-alpine3.20
RUN apk update && apk upgrade
RUN apk add --no-cache gcc musl-dev
RUN apk --no-cache add ca-certificates wget bash xz-libs
WORKDIR /app
COPY go.mod go.sum ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=1 go build -o main ./cmd/server/main.go
EXPOSE 8080
CMD ["./main"]