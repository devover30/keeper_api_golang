FROM golang:1.20.2-alpine3.17 as builder

RUN apk update && apk add --no-cache git

WORKDIR /usr/src/app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /usr/src/app/build

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /usr/src/app/main .
COPY --from=builder /usr/src/app/.env .

# Expose port 9500 to the outside world
EXPOSE 9500

#Command to run the executable
CMD ["./main"]
