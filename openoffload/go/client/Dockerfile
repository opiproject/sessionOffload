# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM docker.io/library/golang:1.20.2-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# build an app
COPY *.go ./
RUN go build -o /opi-sessionoffload-client

EXPOSE 50051
CMD [ "/opi-sessionoffload-client" ]

