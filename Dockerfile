FROM golang:1.13-alpine AS builder
WORKDIR /app

# prevent the reinstallation of vendors at every change in the source code
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy and build the app
COPY . .
RUN go build server.go
RUN go build cli.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/server .
COPY --from=builder /app/cli .

RUN chown -R nobody: /app

USER nobody

CMD [ "./server" ]