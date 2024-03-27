# Define the baseImage
FROM golang:alpine

# Add metadata
LABEL maintainer="Oumar LAM <oumarlamsn@gmail.com>" \
      description="This image will be used to run our socialface app using a Dockerfile" \
      url="http://localhost:8080"

# Define /app as the working directory
WORKDIR /app

# Copy the project from the host to the working directory
COPY . /app/ 

# Download dependencies
RUN go mod download

# Install SQLite
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

# Compile the application
RUN go build -o socialface .

# Define the port where the docker container will listen
EXPOSE 8080

# Run the socialface executable
CMD ["./socialface"]
