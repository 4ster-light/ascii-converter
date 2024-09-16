FROM golang:1.23-alpine3.20

RUN apk --no-cache add \
  gcc \
  musl-dev \
  ca-certificates \
  wget \
  bash \
  xz-libs \
  && apk --no-cache upgrade

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=1 go build -o main ./cmd/server/main.go
EXPOSE 8080
CMD ["./main"]