# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy files
COPY ./backend/ .
COPY ./frontend/dist/ ./dist/

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN GOPROXY="https://gocenter.io" go mod download

# Build the Go app
RUN GOPROXY="https://gocenter.io" go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]
