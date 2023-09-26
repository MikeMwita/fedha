FROM golang:1.21rc2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go


FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .



COPY /services/app-notification/ .

ENTRYPOINT [ "/app/main" ]