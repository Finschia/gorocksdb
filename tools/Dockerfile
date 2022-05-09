FROM golang:1.18-bullseye AS build

RUN apt-get update && apt-get install -y --no-install-recommends \
    make tar wget

FROM build AS install

ARG version="6.20.3"
ARG rocksdb="rocksdb"
ARG archive="v${version}.tar.gz"

WORKDIR .
RUN rm -rf ${rocksdb} ${rocksdb}-${archive}
RUN wget -q -O ${rocksdb}-${archive} https://github.com/facebook/rocksdb/archive/${archive}
RUN tar -zxvf ${rocksdb}-${archive}
RUN mv ${rocksdb}-${version} ${rocksdb}
RUN cd ${rocksdb} \
    && DEBUG_LEVEL=0 make -j4 shared_lib \
    && make install-shared \
    && ldconfig
RUN rm -rf ./rocksdb-*.tar.gz rocksdb

# Install golangci for CI
RUN go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.0

# Download dependency modules for CI
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
RUN rm -rf go.mod go.sum
