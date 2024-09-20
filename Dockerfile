# The builder stage image 
FROM golang:1.22-alpine AS builder 
 
# Set destination for COPY 
WORKDIR /app 
 
# Copy the mod and sum files 
COPY go.mod go.sum ./ 
 
# Download Go modules 
RUN go mod download 
 
# Copy the source code  
COPY ./ ./ 
 
# Build 
RUN CGO_ENABLED=0 GOOS=linux go build -o /main 
 
# The final stage image 
FROM alpine:latest 
 
# Copy the binary from the builder stage 
COPY --from=builder /main /main 
 
# Run the binary 
CMD ["/main"]
