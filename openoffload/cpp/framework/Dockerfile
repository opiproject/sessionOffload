FROM grpcbuild:v1beta1
USER root
ENV PATH="/home/grpc/local/bin:${PATH}"

RUN bash -c "mkdir -p /home/grpc/local/tests/{src,include,ssl}"

WORKDIR /home/grpc/local/tests

COPY src/*.c /home/grpc/local/tests/src/
COPY src/*.cc /home/grpc/local/tests/src/
COPY include/*.h /home/grpc/local/tests/include/
COPY ssl/* /home/grpc/local/tests/ssl/
COPY Makefile /home/grpc/local/tests/Makefile

RUN mkdir -p /home/protos

COPY openoffload.proto /home/grpc/local/tests

RUN bash -c "mkdir -p {bin,obj,log}" && \
    make clean && \
    make HOME=/home/grpc all

USER guest
CMD ["sh", "-c", "tail -f /dev/null"]
