FROM golang:latest

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download

# Command to run the executable
CMD ["air", "-c", ".air.toml"]
