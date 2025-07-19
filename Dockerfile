# --- Build Stage ---
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules and source code
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
# -ldflags="-w -s" strips debug information, reducing binary size
# CGO_ENABLED=0 creates a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /main ./cmd/web

# --- Production Stage ---
FROM gcr.io/distroless/static-debian11

# Set the working directory
WORKDIR /app

# Copy the static binary from the builder stage
COPY --from=builder /main /main

# Copy templates and static assets
COPY web ./web
COPY public ./public

# Expose the application port
EXPOSE 8080

# Set the entrypoint for the container
ENTRYPOINT ["/main"]