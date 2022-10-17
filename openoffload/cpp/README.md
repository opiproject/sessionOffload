# Installing gRPC build environment on CENTOS

The following scripts are taken from Reference 1. The scripts have been built into a Dockerfile that is the same directory as this README.
The framework directory has the latest build instructions.

## Running Docker

Creating the build container

```bash
docker build --tag openoffload:0.1 .
```

Accessing the buidl container

```bash
docker run -i -t openoffload:0.1 /bin/sh
```

## Build Instructions

```bash
mkdir -p $HOME/local
export GRPC_INSTALL=$HOME/local
```

```bash
yum install wget
yum group install "Development Tools"
```

```bash
wget -q -O cmake-linux.sh https://github.com/Kitware/CMake/releases/download/v3.23.4/cmake-3.23.4-Linux-x86_64.sh
sh cmake-linux.sh -- --skip-license --prefix=$GRPC_INSTALL
rm cmake-linux.sh
export PATH=$GRPC_INSTALL/bin:$PATH
```

```bash
git clone --recurse-submodules -b v1.49.1 https://github.com/grpc/grpc
cd grpc
mkdir -p cmake/build
pushd cmake/build
cmake -DgRPC_INSTALL=ON \
      -DgRPC_BUILD_TESTS=OFF \
      -DCMAKE_INSTALL_PREFIX=$GRPC_INSTALL \
      -DCMAKE_POSITION_INDEPENDENT_CODE=TRUE \
      -DCMAKE_CXX_STANDARD=17 \
      -DABSL_PROPAGATE_CXX_STD=TRUE \
      ../..
make -j 4
make install
popd
```

## References

1. [Building gRPC C++](https://grpc.io/docs/languages/cpp/quickstart/)
