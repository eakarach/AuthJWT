#base image
FROM golang:1.25.3-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build Application
RUN go build -o main .

# Run and expose the server
EXPOSE 3000

# Run the web service on container startup.
CMD ["./main"]