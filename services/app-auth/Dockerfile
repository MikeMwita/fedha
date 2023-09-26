FROM golang:1.21rc2-alpine AS builder
WORKDIR /app
COPY app-auth .
RUN go build -o main main.go


# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .



COPY /services/app-auth .

ENTRYPOINT [ "/app/main" ]