# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the entire application to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Start the application
CMD ["/app/main"]