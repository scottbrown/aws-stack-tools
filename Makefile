.DEFAULT_GOAL: build

pkg := github.com/scottbrown/aws-stack-tools

build:
	go build -o .build/aws-stack-instances $(pkg)/cli/aws-stack-instances
	go build -o .build/aws-stacks $(pkg)/cli/aws-stacks

test:
	go test ./...

fmt:
	go fmt ./...
