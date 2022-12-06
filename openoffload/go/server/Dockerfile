# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19.4-bullseye

WORKDIR /app

# Dependencies for goebpf
RUN apt-get -y update && \
    DEBIAN_FRONTEND="noninteractive" apt-get --no-install-recommends -y install tzdata clang-9 libclang-9-dev libllvm9 libelf-dev git golang llvm-9-dev gcc flex bison cmake make python libpcap-dev ca-certificates && \
    apt-get -y clean all && \
    rm -rf /var/lib/apt/lists/*

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# build an app
COPY *.go ./
RUN go build -o /opi-sessionoffload-bridge

EXPOSE 50051
CMD [ "/opi-sessionoffload-bridge" ]
