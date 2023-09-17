# Use the official Go image as the base image
FROM golang:1.21.0-alpine

# membuat direktori folder
# RUN mkdir /app

# set working direktori i
WORKDIR /app

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Starting our application
CMD ["go", "run", "main.go"]

# Exposing server port
EXPOSE 80