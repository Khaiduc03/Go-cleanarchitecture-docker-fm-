
FROM golang:1.21-alpine AS api

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine AS prod

WORKDIR /app

COPY --from=api /app/main .

COPY --from=api /app/.env .

COPY --from=api /app/firebase.json .
ENV CLOUDINARY_SKIP_TLS_VERIFY 1
EXPOSE 7000

CMD ["./main"]

