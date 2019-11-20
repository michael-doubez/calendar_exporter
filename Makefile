GITHUB_ORG  = mdoubez
GITHUB_REPO = calendar_exporter
VERSION     = 0.0.1

# Go projet
GO = go

# Inject version information
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
REVISION ?= $(shell git rev-parse --short HEAD)
BUILDUSER ?= $(USER)
BUILDDATE ?= $(shell date +%FT%T%z)
LDFLAGS = -s -X github.com/prometheus/common/version.Version=$(VERSION) \
		     -X github.com/prometheus/common/version.Branch=$(BRANCH) \
		     -X github.com/prometheus/common/version.Revision=$(REVISION) \
		     -X github.com/prometheus/common/version.BuildUser=$(BUILDUSER) \
		     -X github.com/prometheus/common/version.BuildDate=$(BUILDDATE)

# ------------------------------------------------------------------------
# Main targets
# - all: check code and build it
# - build: build exporter for current platform
# - clean: remove build files
# - check: run all following checks
#   - fmt: run formating check
#   - vet: vetting code
.PHONY: all build clean check fmt vet
EXPORTER=calendar_exporter

all:: vet fmt build

build: $(EXPORTER)

clean:
	@rm -f $(EXPORTER)

check: fmt vet

fmt:
	@$(GO) fmt ./...

vet:
	@$(GO) vet ./...

# ------------------------------------------------------------------------
# Build exporter

# Main source files
SRCS = $(wildcard *.go calendar/*.go)

$(EXPORTER): $(SRCS)
	@go build -ldflags "$(LDFLAGS)" $(EXPORTER).go

