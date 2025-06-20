# FROM golang:latest

# WORKDIR /app

# # Copy go mod files
# COPY go.mod ./
# RUN go mod download

# # Copy the source code
# COPY main.go ./
# COPY handler ./handler
# COPY server ./server

# # Expose the port your app listens on
# EXPOSE 8080

# # Run the binary
# CMD ["go","run","main.go"]

# syntax=docker/dockerfile:1

########################################################
# 1) Build Stage: Compile a statically-linked Go binary
########################################################
FROM golang:latest AS builder
WORKDIR /app

# Enable CGO=0 for a fully static binary (Distroless has no C libs) :contentReference[oaicite:0]{index=0}
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy dependency files and download modules
COPY go.mod ./
RUN go mod download

# Copy all source and build
COPY main.go ./
COPY handler ./handler
COPY server ./server
RUN go build -o main ./main.go

########################################################
# 2) Runtime Stage: Distroless static image
########################################################
FROM gcr.io/distroless/static

# Working directory (optional)
COPY --from=builder /app/main /main

USER nonroot:nonroot
ENTRYPOINT ["/main"]
