# Forked gorocksdb, a Go wrapper for RocksDB for LBM
- https://github.com/tecbot/gorocksdb

## Install

You'll need to build [RocksDB](https://github.com/facebook/rocksdb) on your machine.
- Support from v5.16 to [v6.20.3](https://github.com/facebook/rocksdb/tree/v6.20.3)

After that, you can install gorocksdb using the following command:

    CGO_CFLAGS="-I/path/to/rocksdb/include" \
    CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm" \
        $(shell awk '/PLATFORM_LDFLAGS/ {sub("PLATFORM_LDFLAGS=", ""); print}' \
        < /path/to/rocksdb/make_config.mk \
        go get github.com/line/gorocksdb

Please note that this package might upgrade the required RocksDB version at any moment.
Vendoring is thus highly recommended if you require high stability.

*The [embedded CockroachDB RocksDB](https://github.com/cockroachdb/c-rocksdb) is no longer supported in gorocksdb.*
