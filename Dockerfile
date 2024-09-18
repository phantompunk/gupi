# syntax=docker/dockerfile:1

# Create stage for building the application
ARG VERSION="0.0.3"
ARG GO_VERSION=1.22.5
ARG ALPINE_VERSION=3.18
FROM golang:${GO_VERSION}-alpine AS builder

# Set the workdir
WORKDIR /app

# Install app dependencies
RUN --mount=type=bind,src=go.mod,target=go.mod \
    --mount=type=bind,src=go.sum,target=go.sum \
    --mount=type=cache,target=/go/pkg/mod \
    go mod download -x

# Copy the Go source code
COPY . /app

# Build the Go binary
RUN go build -ldflags "-X 'github.com/phantompunk/gupi/cmd.version=${VERSION}'" -o /app/gupi .

# Create a new image for the runtime
FROM alpine:${ALPINE_VERSION} AS final

# Copy the built binary from the builder stage
COPY --from=builder /app/gupi /bin/gupi
