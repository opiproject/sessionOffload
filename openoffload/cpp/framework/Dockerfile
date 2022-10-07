FROM grpcbuild:v1beta1
USER root
ENV PATH="/home/grpc/local/bin:${PATH}"

RUN mkdir -p /home/grpc/local/tests && \
    mkdir -p /home/grpc/local/tests/src/ && \
    mkdir -p /home/grpc/local/tests/include/ && \
    mkdir -p /home/grpc/local/tests/ssl/

WORKDIR /home/grpc/local/tests

COPY src/*.c /home/grpc/local/tests/src/
COPY src/*.cc /home/grpc/local/tests/src/
COPY include/*.h /home/grpc/local/tests/include/
COPY ssl/* /home/grpc/local/tests/ssl/
COPY Makefile /home/grpc/local/tests/Makefile

RUN mkdir -p /home/protos

COPY openoffload.proto /home/protos/

RUN ls ../../../protos && \
    mkdir -p bin && \
    mkdir -p obj && \
    mkdir -p log && \
    make clean && \
    make HOME=/home/grpc all

CMD ["sh", "-c", "tail -f /dev/null"]
