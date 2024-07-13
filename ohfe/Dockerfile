# Build stage

FROM golang:latest AS BuildStage
WORKDIR /app
# COPY go.mod go.sum ./
COPY . .

# RUN go build -o ./main cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main cmd/main.go

# Deploy Stage
FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage /app/main /app/main

CMD ["/app/main"]
