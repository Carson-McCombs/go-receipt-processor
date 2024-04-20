# syntax=docker/dockerfile:1
# Hosts a local server on port 8080, written in Go 1.20, and packages into a container image.

FROM golang:1.20

WORKDIR /app

#Copies go.mod and go.sum, which holds references to all of the dependencies, into the application folder
COPY go.mod go.sum ./

#Downloads all the dependencies needed to run
RUN go mod download

#Copies the source code into the application folder
COPY . ./

#Builds the binaries
RUN  CGO_ENABLED=0 go build  -o /go-receipt-processor .

#Exposes what port the application communicates to Docker with
EXPOSE 8080

#Runs the executable
CMD ["/go-receipt-processor"]


