# syntax=docker/dockerfile:1

# Create stage for building the application
ARG GO_VERSION=1.22.5
FROM golang:${GO_VERSION}-alpine AS builder

# Set the workdir
WORKDIR /app

# Copy the Go source code
COPY . .

# Install system dependencies including 'make'
RUN apk update && apk add --no-cache gcc libc-dev make

# Build the Go binary
RUN make

# Create a new image for the runtime
FROM alpine:3.17

# Copy the built binary from the builder stage
COPY --from=builder /app/gupi /bin/gupi
