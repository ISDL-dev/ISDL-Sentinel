# ベースステージ
FROM golang:1.23.1-alpine AS base
WORKDIR /go/src/app
RUN apk update && apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download

# 開発環境ステージ
FROM base AS dev
RUN go install github.com/air-verse/air@v1.60.0
COPY . .
CMD ["air", "-c", ".air.toml"]

# ビルドステージ
FROM base AS build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# 本番環境ステージ
FROM base AS prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/app/main .
CMD ["./main"]

EXPOSE 3000