SERVICE_NAME ?= $(error SERVICE_NAME is not set)
RELEASE_VERSION ?= $(error RELEASE_VERSION is not set)

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: help
## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code

.PHONY: generate init commit tag push

all: generate init commit tag push

generate:
	buf generate --path proto/$(SERVICE_NAME)

init:
	cd golang/$(SERVICE_NAME) && \
	go mod init github.com/materials-resources/microservices-proto/golang/$(SERVICE_NAME) || true && \
	go mod tidy

commit:
	git add golang/$(SERVICE_NAME) && \
    git commit -m "Update $(SERVICE_NAME) proto" || true

tag:
	git tag -f golang/$(SERVICE_NAME)/$(RELEASE_VERSION) \
    		$(shell git subtree split --prefix golang/$(SERVICE_NAME) HEAD) \
    		-m "golang/$(SERVICE_NAME)/$(RELEASE_VERSION)"

push:
	git push origin golang/$(SERVICE_NAME)/$(RELEASE_VERSION)

