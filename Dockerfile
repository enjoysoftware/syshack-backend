FROM golang:bookworm AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ランタイムイメージを使用する (より軽量)
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
# アプリケーションのポート
EXPOSE 8080 

CMD ["./main"]
