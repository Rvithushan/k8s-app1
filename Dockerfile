#Stage 1: Build the Go binary statically
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Download dependencies 
COPY go.mod ./
RUN go mod download

#Copy source code 
COPY main.go .

#Build static binary with disabled CGO
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -0 sre-app1 main.go

#Stage 2: Minimal Distroless/static-debinan12:nonroot

WORKDIR /

# Copy executable from builder
COPY --from=builder /app/sre-app1 /sre-app1

# Run as non-root user for security compliance
USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/sre-app1"]

