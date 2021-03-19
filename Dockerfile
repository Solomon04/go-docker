FROM golang:1.13-alpine AS builder
WORKDIR /app

# prevent the reinstallation of vendors at every change in the source code
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy and build the app
COPY . .
RUN go build -o go-docker .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/go-docker .

RUN chown -R nobody: /app

USER nobody

CMD [ "./go-docker" ]