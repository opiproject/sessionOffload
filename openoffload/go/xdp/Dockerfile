# syntax=docker/dockerfile:1

# First, we install everything to build the XDP code
FROM debian:bullseye-20221024 as build

RUN apt-get -y update && \
    DEBIAN_FRONTEND="noninteractive" apt-get --no-install-recommends -y install tzdata clang-9 libclang-9-dev libllvm9 libelf-dev git golang llvm-9-dev gcc flex bison cmake make python libpcap-dev ca-certificates && \
    apt-get -y clean all && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p /opt/xdp
WORKDIR /opt/xdp
COPY xdp/ .
RUN make

# Next, we build the golang application
FROM golang:1.20.1-bullseye as server-builder

WORKDIR /app

# Download necessary Go modules
COPY server/go.mod ./
COPY server/go.sum ./
RUN go mod download

# build an app
COPY server/*.go ./
RUN go build -o /opi-sessionoffload-bridge

# Finally, we install it all together into a final container
FROM debian:bullseye-20221024-slim

RUN apt-get -y update && \
    DEBIAN_FRONTEND="noninteractive" apt-get --no-install-recommends -y install iproute2 ethtool && \
    apt-get -y clean all && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir /opt/xdp
WORKDIR /opt/xdp
COPY --from=build /opt/xdp/scripts/start.sh /opt/xdp/start.sh
COPY --from=build /opt/xdp/xdp.elf /opt/xdp
COPY --from=build /opt/xdp/loader/xdploader /opt/xdp

COPY --from=server-builder /opi-sessionoffload-bridge /opi-sessionoffload-bridge
EXPOSE 50051
