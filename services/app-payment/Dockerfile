FROM golang:1.21rc2-alpine AS builder
LABEL app-payment https://github.com/MikeMwita/fedha.git/services/app-payment

WORKDIR /app
COPY . .
RUN go build -o main main.go


FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .


FROM scratch
COPY /services/app-payment/ .

ENTRYPOINT [ "/app/main" ]