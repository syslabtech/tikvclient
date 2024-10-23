# Use a valid Go version and base image
FROM golang:1.23-bookworm AS build

# Set the working directory inside the container
WORKDIR /app


# Copy the rest of the application source code
COPY . .

# Copy go.mod and download dependencies
RUN go mod download && go mod verify

# Build the Go binary for a Linux target
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server -a -ldflags="-s -w" -installsuffix cgo

# Install UPX to compress the binary
RUN apt-get update && apt-get install -y upx

# Compress the binary using UPX
RUN upx --ultra-brute -qq server && upx -t server 

# Stage 2: Use a minimal image for the final production build
FROM scratch

# Copy the compressed binary from the build stage
COPY --from=build /app/server /server

# Run the server binary
ENTRYPOINT ["/server"]
