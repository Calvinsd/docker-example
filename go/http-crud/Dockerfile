# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine

# Creating a working directory 
WORKDIR /app

# Downloading deps
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copying all go files
COPY *.go ./

# Building go
RUN go build main.go

# Run the binary
CMD ["./main"]



