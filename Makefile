dprint-sh:
	curl -fsSL https://dprint.dev/install.sh | sh

dprint-ps1:
	iwr https://dprint.dev/install.ps1 -useb | iex

deps:
	go mod tidy

fmt:
	dprint fmt
	gofmt -l -s -w .

install:
	go install
