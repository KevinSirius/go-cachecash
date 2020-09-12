# Go-Cachecash

[![Build Status](https://travis-ci.com/cachecashproject/go-cachecash.svg?token=utLK2DGqpJaDNkKeJ4fh&branch=master)](https://travis-ci.com/cachecashproject/go-cachecash)

[![Coverage Status](https://coveralls.io/repos/github/cachecashproject/go-cachecash/badge.svg?t=0cosgH)](https://coveralls.io/github/cachecashproject/go-cachecash)

Go-Cachecash is a project focusing on decentralized content distribution.
## Cloning the git repository

This repository uses `git-lfs` for large test data artifacts, among other things; you'll need to install it:

```bash
# Ubuntu
apt-get install git-lfs

# macOS
brew install git-lfs

# Archlinux
pacman -S git-lfs

# Windows
Download from https://git-lfs.github.com/
```

Next, clone the cachecash repo:


# Please Note now Go no longer needs to keep codes under GOPATH anymore. You can replace the directory with any folder you want.
```bash
git clone git@github.com:cachecashproject/go-cachecash.git "$(go env GOPATH)/src/github.com/cachecashproject/go-cachecash"
```

And initialize git-lfs in that repo:

```bash
cd "$(go env GOPATH)/src/github.com/cachecashproject/go-cachecash"
git lfs install
git lfs fetch
git lfs checkout
```

## Setting up database secret
Go to /deploy/secret and create new .secret file named after each database container based on template files provided.
For example, this is my ledger.secret:
```bash
POSTGRES_USER=cachecash
POSTGRES_PASSWORD=cachecashPW
POSTGRES_DB=ledger
LEDGER_DATABASE=host=ledger-db port=5432 user=cachecash password=cachecashPW dbname=ledger sslmode=disable
```

## Setting up a development environment

You will need a working Go toolchain.  We tend to stay on the latest stable version.

You will also need some extra code generation tools:

```bash
make dev-setup
```

To generate source from proto files:

```bash
make gen
```

To generate documentation from the proto files:

```bash
make gen-docs
```
To generate all the binaries:
```bash
make all
```

## Running a local test network

`Please notice you need to update your docker-compose to the latest version. Download it here: https://docs.docker.com/compose/compose-file/`

For the first time of running, you need to build all images:


```bash
docker-compose build
```

Next, bring up the network:

```bash
docker-compose up
```

## Debug Cachecash network
If you are having trouble launching the network after building, or if the daemons keep showing errors, you need to examine the network by pull up the network in the background and see the docker containers' status.
Use the following command to run the network in background.
```bash
docker-compose up -d
```
Then examine the container's status
```bash
docker ps
```
If any containers exit with code 1, please check the logs
```bash
docker logs "daemon name"
```
## Test Network Locally
It's going to take 1-2 minutes until everything is initialized, this includes the postgres initialization, the caches
announcing themselves to the bootstrap service, the publisher requesting a list of all caches from the bootstrap service
and finally the publisher and the caches negotiating an escrow, and a distributed tracing backend store. After the escrow
is setup you can download from the caches using `cachecash-curl` or [typescript-cachecash]
In this case we're going to build the `cachecash-curl` program.'cachecash-curl' is a tool that tests the ability to transfer data between publisher and client.The resulting file
(here, `output.bin`) should exactly match the original artifact (here, `testdata/content/file0.bin`).

The `-logLevel` option can be changed to control output verbosity for each program.

To build cachecash-curl, please run 
```bash
make ccc
```

To test
```bash
./bin/cachecash-curl -o output.bin -logLevel=debug -trace http://localhost:14268 cachecash://localhost:7070/file0.bin
diff output.bin testdata/content/file0.bin
```
[typescript-cachecash]: https://github.com/cachecashproject/typescript-cachecash


## Using the base image

To get the base image we use for building our software, type `make
pull-base-image` to receive it. It will be used for all operations around
building and running the software in our development environment.

To build the image yourself:

```shell
$ make base-image
```

If you have write access to our repositories on docker hub, you can also push
an updated image:

```shell
$ make push-base-image
```

This is recommended whenever your changes land in the `master` branch, but
**not before then**.

## Running tests

```bash
# Remove cached results.  (Should not normally be necessary, but can be useful while working on the test suite.)
go clean -testcache

# Run unit and integration tests that do not have external dependencies.
go test -v -race ./...
```

## Running tests that use postgres

To run tests that use postgres, you must launch it first; you can use this command to launch your postgres container:

```shell
docker run --name db-hostname -p 5432:5432 -d -e POSTGRES_DB=dbname postgres:11

# Run a test that has external dependencies, including those automatically
# generated by the ORM system.

PSQL_USER=postgres PSQL_SSLMODE=disable PSQL_HOST=localhost PSQL_DBNAME=dbname \
go test -v -race -tags="external_test sqlboiler_test" ./path/to/test/...
```

## Environment variables

CACHECASH_INSECURE can be set for any binary to disable TLS checking against the
bootstrap/publisher/observability endpoints. Protocol cryptography is still
secure.

## Fuzzing

There are fuzzing targets that can be used with [go-fuzz](https://github.com/dvyukov/go-fuzz).

```
# install dependencies
go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
# start fuzzing
make fuzz
```
