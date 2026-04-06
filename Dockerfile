# Stage 1: Test
FROM golang:tip-alpine3.23 AS tester
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY internal ./internal
COPY tests ./tests
RUN go test -v ./...
RUN go test ./... -coverpkg=./... -coverprofile=coverage.out
RUN go tool cover -html=coverage.out -o coverage.html

# Stage 2: Build
FROM golang:tip-alpine3.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY internal ./internal
COPY cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

# Stage 3: Runtime (small image)
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ENV ENV=production
ENV GIN_MODE=release
EXPOSE 8080
COPY --from=builder /app/app .
COPY ./data/cv.json ./data/cv.json
CMD ["./app"]