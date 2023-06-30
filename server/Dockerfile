# official go image as a base
FROM golang:latest

# working directory
WORKDIR /app


# copy the go go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the source code
COPY . .

# build the application
RUN go build -o main

# expose the api port
EXPOSE 8080

# start the api server
CMD ["/app/main"]
