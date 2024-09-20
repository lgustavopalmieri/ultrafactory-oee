FROM golang:alpine AS builder
WORKDIR /app
RUN apk add --no-cache git make \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY go.mod .
COPY go.sum .
COPY ./internal/infra/database/postgres/sql/migrations /app/internal/infra/database/postgres/sql/migrations
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest
RUN adduser --uid 1000 --disabled-password appuser
WORKDIR /app
COPY --from=builder /app/main ./main
RUN chmod +x ./main
USER appuser
EXPOSE 4003

CMD ["./main"]
