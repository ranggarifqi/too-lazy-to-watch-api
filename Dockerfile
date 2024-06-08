# Use an official Golang image as the base
FROM golang:alpine

# Set the working directory in the container to /app
WORKDIR /app

# Copy the current directory contents into the container at location /app
COPY . /app/

# Install any needed dependencies
RUN go mod download
RUN make build

# Expose port 3000 and run command when the container starts
EXPOSE 3000
CMD ["./app"]