[![Build Status](https://travis-ci.com/dapperlabs/bamboo-node.svg?token=MYJ5scBoBxhZRGvDecen&branch=master)](https://travis-ci.com/dapperlabs/bamboo-node)

## Test
```
$ ./test.sh
```
If iterating just on failed test, then we can do so without rebuilding the system:
```
$ docker-compose up --build --no-deps test
```
TODO: move to Makefile (remove also shell script)


## Build (updates go modules)
```
$ GO111MODULE=on go build -o donotcommit ./cmd/execute/
$ GO111MODULE=on go build -o donotcommit ./cmd/security/
$ GO111MODULE=on go build -o donotcommit ./cmd/testhelpers/
```
TODO: move to Makefile


## Generate dependency injection
### Prerequisite 
Install wire: `$ GO111MODULE=on go get -u github.com/google/wire/cmd/wire`
### Command
```
$ GO111MODULE=on wire ./internal/execute/
$ GO111MODULE=on wire ./internal/security/
$ GO111MODULE=on wire ./internal/access/
```
TODO: move to Makefile

## Generate protobufs 
### Prerequisite 
1. Install prototools https://github.com/uber/prototool#installation  
2. $ go get -u github.com/golang/protobuf/protoc-gen-go
### Command
```
$ prototool generate proto/
```
TODO: move to Makefile
