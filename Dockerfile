FROM golang:1.18

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ports.json ./
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

# Run
CMD ["/app/main"]