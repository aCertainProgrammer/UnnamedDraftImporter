FROM golang:alpine AS builder

WORKDIR /app

COPY . .
RUN go build -o main .

FROM alpine

COPY --from=builder /app/main /

EXPOSE 80
CMD ["./main"]
