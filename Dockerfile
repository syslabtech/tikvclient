# Use a valid Go version and base image
FROM golang:1.23-bookworm AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and download dependencies



# Copy the rest of the application source code
COPY . .
RUN go mod download && go mod verify
# Build the Go binary for a Linux target
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server -a -ldflags="-s -w" -installsuffix cgo

# Install xz-utils for extracting .tar.xz and curl to download the UPX binary
RUN apt-get update && apt-get install -y xz-utils curl

# Download UPX manually and install it
RUN curl -L https://github.com/upx/upx/releases/download/v4.0.2/upx-4.0.2-amd64_linux.tar.xz -o upx.tar.xz \
    && tar -xf upx.tar.xz \
    && cp upx-4.0.2-amd64_linux/upx /usr/local/bin/ \
    && rm -rf upx.tar.xz upx-4.0.2-amd64_linux

# Compress the binary using UPX
RUN upx --ultra-brute -qq server && upx -t server 

# Stage 2: Use a minimal image for the final production build
FROM debian:bookworm-slim

# Set the working directory inside the container
WORKDIR /app

# # Copy the compressed binary from the build stage
COPY --from=build /app/server /app/server
COPY --from=build /app/static/* /app/static
COPY --from=build /app/templates/* /app/templates

# Run the server binary
ENTRYPOINT ["/app/server"]
