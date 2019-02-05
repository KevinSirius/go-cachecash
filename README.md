# go-cachecash

## Running the test server

There is a binary named `testserverd` that can be useful in situations where you do not need a full network.  It runs an
HTTP origin (which serves the test artifacts in `testdata/content`), a publisher pointed at that origin, and four
caches.  The publisher/cache relationships are automatically configured for you.

You can point the `cachecash-curl` program (or another client) at this publisher to see it fetch data.  The resulting
file (here, `output.bin`) should exactly match the original artifact (here, `testdata/content/file0.bin`).

The `-logLevel` option can be changed to control output verbosity for each program.

```
go build -o bin/testserverd ./cmd/testserverd && ./bin/testserverd -logLevel=info
go build -o bin/cachecash-curl ./cmd/cachecash-curl && ./bin/cachecash-curl -o output.bin -logLevel=debug cachecash://localhost:8080/file0.bin
diff output.bin testdata/content/file0.bin
```

## Setting up a development environment

This repository uses `git-lfs` for test data artifacts, among other things; you'll need to install it.

```
# Ubuntu
apt-get install git-lfs

# macOS
brew install git-lfs
```

You will need a working Go toolchain.  We tend to stay on the latest stable version.

You should check out this repository into your `GOPATH` (e.g. `~/src/go/github.com/cachecashproject/go-cachecash`).

You will also need some extra code generation tools.

```
go get -u github.com/rubenv/sql-migrate/...
go get -u github.com/volatiletech/sqlboiler/...
go get -u github.com/volatiletech/sqlboiler-sqlite3/...
go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/...
```

TODO: You will need `protoc` and `gogo-protobuf`.  We should use a Docker container for these tools; otherwise,
variations in tool versions will cause all sorts of headaches.
