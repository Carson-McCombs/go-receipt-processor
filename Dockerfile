# syntax=docker/dockerfile:1
# Hosts a local server on port 8080, written in Go 1.22.2, and packages into a container image.

# Uses a multistage docker build to minimize image size

FROM golang:1.22.2-bullseye as base

#adds non-root user
RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  non-root-user

# Sets the active work directory
WORKDIR /app

# Copies go.mod and go.sum, which holds references to all of the dependencies, into the application folder
COPY go.mod go.sum ./

# Downloads all the dependencies needed to run
RUN go mod download

# Copies the source code into the application folder
COPY . ./

# Builds the binaries
RUN  CGO_ENABLED=0 GOOS=linux go build -o /go-receipt-processor .

# Runs tests
FROM base AS run-test
RUN go test -v ./...

FROM scratch

# Keeps specific information for security with minimum function
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

# Keeps our actual project
COPY --from=base /go-receipt-processor .

# Sets user / user-permissions
USER non-root-user:non-root-user

# Exposes internal port for communication
EXPOSE 8080

CMD ["./go-receipt-processor"]