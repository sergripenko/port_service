# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Serhii Ripenko <sergripenko@gmail.com>"

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY ports.json ./
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
#COPY --from=builder /app/.env .
COPY --from=builder /app/ports.json .
COPY --from=builder /app/go.mod /root/
COPY --from=builder /app/go.sum /root/

# Command to run the executable
CMD ["./main"]



#FROM golang:1.18
#
## Set destination for COPY
#WORKDIR /app
#
## Download Go modules
#COPY go.mod go.sum ./
#RUN go mod download
#
#COPY ports.json ./
#COPY . .
#
## Build
#RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go
#
## Run
#CMD ["/app/main"]
