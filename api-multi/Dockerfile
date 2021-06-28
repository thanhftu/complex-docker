FROM golang:1.16.3-alpine3.13 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main main.go



# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .

# CMD ["./main"]
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
# COPY app.env .
# COPY start.sh .
# COPY wait-for .
EXPOSE 8081
# COPY db/migration ./migration

# EXPOSE 8000
CMD [ "/main" ]
# ENTRYPOINT [ "/start.sh" ]