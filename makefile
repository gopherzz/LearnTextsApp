.PHONY: build
build:
	go build -o learntexts.exe ./pkg/app

.PHONY: test
test:
	go test ./pkg/textwork/worderaser