GO=go
DOCKER=docker
DOCKER_TAG=craigcryptoapp:main
GOLINT=golint
GOSTATICCHECK=staticcheck


.PHONY: build
build: code-check
	$(GO) build -o crypto-app ./...

.PHONY: vet
vet:
	$(GO) vet ./...

.PHONY: lint
lint:
	$(GOLINT) ./...

.PHONY: static-check
static-check:
	$(GOSTATICCHECK) ./...

.PHONY: image
image:
	$(DOCKER) build --tag $(DOCKER_TAG) -f ./images/Dockerfile .

.PHONY: code-check
code-check: vet lint static-check

.PHONY: get-gotools
get-gotools:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	# Go Lint is deprecated
	go get -u golang.org/x/lint/golint

.PHONY: clean
clean:
	rm -f crypto-app
	docker rmi $(DOCKER_TAG)

.PHONY: test
test:
	go test -v ./...
