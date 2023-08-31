FROM basegoimage

WORKDIR /app

COPY .air.* ./
COPY go.mod go.sum ./
RUN go mod download

# Command to run the executable
CMD ["air", "-c", ".air.toml"]