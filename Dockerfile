FROM golang:1.21rc2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go


# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY start.sh .
COPY wait-for.sh /app/


COPY /services/app-db/db/migration ./db/migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]