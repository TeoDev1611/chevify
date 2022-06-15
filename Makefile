dprint-sh:
	curl -fsSL https://dprint.dev/install.sh | sh

dprint-ps1:
	iwr https://dprint.dev/install.ps1 -useb | iex

dependencies:
	go mod tidy

format:
	dprint fmt
	go fmt ./..

install:
	go install
