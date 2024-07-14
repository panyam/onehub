# Build stage

FROM golang:latest AS BuildStage
WORKDIR /app

COPY locallinks ./locallinks
COPY Makefile go.mod go.sum ./
RUN mkdir cmd clients
COPY cmd ./cmd
COPY clients ./clients
RUN ls -al

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main cmd/dbsync/main.go

# Deploy Stage
FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage /app/main /app/main

CMD ["/app/main"]
