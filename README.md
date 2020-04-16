# Hello, World!


## Build

```sh
# build the hello cmd, and output to /tmp
# -a skip the build cache
go build -a -o=/tmp/hello ./cmd/hello

# run hello
cd /tmp
./hello

# target a specific OS and architecture
GOOS=linux GOARCH=amd64 go build -a -o=/tmp/linux_amd64/hello ./cmd/hello
```

## Test

```sh
# -v verbose
# -count=1 don't cache test results
# -failfast stop on the first test failure
# -race detect race conditions in concurrent code
# -vet=off disable vetting when running tests
# -coverprofile generate a coverage report
go test -v -count=1 -failfast -race -vet=off -coverprofile=/tmp/profile.out ./...

# show coverage report for each func in stdout
go tool cover -func=/tmp/profile.out
```

## Static Analysis

```sh
go vet ./...
```

## Linting

```sh
go get -u golang.org/x/lint/golint
golint ./...
```

## Formatting

```sh
# recursively format all files in and beneath the current directory
# -w rewrite files in place
# -s apply simplifications where possible
# -d show diffs
gofmt -w -s -d .
```

## Manage Dependencies
These commands should be run pre-commit.

```sh
# view final versions that will be used in a build, including all direct and indirect dependencies
go list -m all

# view available minor and patch upgrades for all direct and indirect dependencies
go list -m -u all

# update all direct and indirect dependencies to latest minor upgrades
# -t include test dependencies
go get -u -t ./...

# prune any unused dependencies from go.mod
go mod tidy

# check that the local dependencies are the same as those defined in go.mod and go.sum
go mod verify

```
