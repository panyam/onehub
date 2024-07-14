# Build stage

FROM golang:latest AS BuildStage
WORKDIR /app

COPY locallinks ./locallinks
COPY Makefile go.mod go.sum ./
RUN mkdir cmd clients datastore services gen
COPY cmd ./cmd
COPY clients ./clients
COPY datastore ./datastore
COPY services ./services
COPY gen ./gen
RUN ls -al

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main cmd/backend/main.go

# Deploy Stage
FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage /app/main /app/main

CMD ["/app/main"]
