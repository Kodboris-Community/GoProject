FROM golang:1.21-alpine AS builder
FROM golang:1.21-alpine AS builder
WORKDIR /app

RUN export GO111MODULE=on
COPY go.mod  .
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM golang:1.19-alpine
WORKDIR /app
COPY --from=builder /app/main ./main
COPY --from=builder /app/app.env ./app.env
COPY --from=builder /app/db-wait.sh ./db-wait.sh
EXPOSE 3000
EXPOSE 3000
CMD ["./main"]