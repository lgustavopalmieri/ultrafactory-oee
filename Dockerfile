FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest
RUN adduser --uid 1000 --disabled-password appuser
WORKDIR /app
COPY --from=builder /app/main ./main
COPY --from=builder /app/.env ./
RUN chmod +x ./main
USER appuser
EXPOSE 4003
CMD ["./main"]
