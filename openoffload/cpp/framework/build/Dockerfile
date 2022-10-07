FROM centos:8.3.2011
USER root

RUN mkdir -p /home/grpc/local

WORKDIR /home/grpc/local

ENV GRPC_INSTALL /home/grpc/local

RUN echo $GRPC_INSTALL && \
    yum -y update && \
    yum -y install wget && \
    yum -y install which && \
    yum -y groupinstall "Development tools" && \
    wget --no-check-certificate -q -O cmake-linux.sh https://github.com/Kitware/CMake/releases/download/v3.17.5/cmake-3.17.5-Linux-x86_64.sh && \
    ["sh", "cmake-linux.sh",  "--", "--skip-license", "--prefix=/home/grpc/local"] && \
    rm cmake-linux.sh

ENV PATH="$GRPC_INSTALL/bin:${PATH}"

RUN echo $PATH && \
    git config --global  http.sslVerify false && \
    git clone --recurse-submodules -b v1.30.0 https://github.com/grpc/grpc && \
    cd grpc && \
    mkdir -p /home/grpc/local/grpc/cmake/build

WORKDIR /home/grpc/local/grpc/cmake/build

RUN echo $PATH

#CMD ["sh", "-c", "tail -f /dev/null"]

RUN pwd && \
    cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=/home/grpc/local ../.. && \
    make -j 4 && \
    make install && \
    cd ../.. && \
    mkdir -p /home/grpc/tmp

WORKDIR /home/grpc/tmp

RUN wget http://hyperrealm.github.io/libconfig/dist/libconfig-1.7.2.tar.gz && \
    tar xvzf libconfig-1.7.2.tar.gz

WORKDIR /home/grpc/tmp/libconfig-1.7.2

RUN ./configure && \
    make && \
    RUN make install

WORKDIR /home/grpc/
#
CMD ["sh", "-c", "tail -f /dev/null"]
