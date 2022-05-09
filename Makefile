PACKAGES=$(shell go list ./...)
DOCKER_NAME=gorocksdb-testing:main-golang-upgrade
DOCKER_IMAGE=line/$(DOCKER_NAME)

build-docker:
	## If you met the compile error, it might be the shortage of memory on docker with your machine
	## Please check your Docker Desktop and its settings of resources
	## for mac: https://docs.docker.com/desktop/mac/#resources
	## for windows: https://docs.docker.com/desktop/windows/#resources
	@docker build --rm --progress plain --tag="$(DOCKER_IMAGE)" -f ./tools/Dockerfile .
.PHONY: build-docker

bash-docker:
	## Should: `make build-docker`
	## bash
	@docker run --rm -it -v $(CURDIR):/workspace -w /workspace $(DOCKER_IMAGE) bash
.PHONY: bash-docker

lint-docker:
	## Should: `make build-docker`
	## Run go get & lint & verify
	@docker run --rm -v $(CURDIR):/workspace -w /workspace $(DOCKER_IMAGE) bash -c "\
	echo '-> Run go lint' \
	&& golangci-lint run -v \
	&& echo '-> Run go verify' \
    && go mod verify\
    "
.PHONY: lint-docker

test-docker:
	## Should: `make build-docker`
	## Run go test
	@docker run --rm -v $(CURDIR):/workspace -w /workspace $(DOCKER_IMAGE) bash -c "\
	echo '-> Run go test' \
	&& go test $(PACKAGES) -race -v \
	"
.PHONY: test-docker
