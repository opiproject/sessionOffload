FROM debian:bullseye-20221024

RUN apt-get -y update && \
    DEBIAN_FRONTEND="noninteractive" apt-get --no-install-recommends -y install tzdata clang-9 libclang-9-dev libllvm9 libelf-dev git golang llvm-9-dev gcc flex bison cmake python libpcap-dev && \
    apt-get -y clean all && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p /ot/src
WORKDIR /opt/src
