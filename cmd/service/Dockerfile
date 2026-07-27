FROM golang:1.25.11-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

RUN apk --no-cache add bash make gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM scratch

WORKDIR /app
COPY --from=builder /build/temp/bin/service ./service

ENTRYPOINT ["/app/service"]
EXPOSE 50051
