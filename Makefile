
# Enable Go modules:
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
export GOPRIVATE=gitlab.cee.redhat.com/service

# Disable CGO so that we always generate static binaries:
export CGO_ENABLED=0


all:
	go build .
